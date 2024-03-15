package custom

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"

	"crypto/internal/fireblocks/client"
	"crypto/internal/fireblocks/transaction"
	"crypto/internal/fireblocks/vault"
)

type CustomFlowSVC struct {
	client *client.Client
	vault  *vault.VaultService
	txn    *transaction.TransactionService
}

func New(client *client.Client, vault *vault.VaultService) *CustomFlowSVC {
	return &CustomFlowSVC{
		client: client,
		vault:  vault,
		txn:    transaction.New(client),
	}
}

// when u put a vault id it doesnt create it
// when u dont it creates a new vault id
func (v *CustomFlowSVC) CreateAddress(
	ctx context.Context, userID string, vaultAccountID *string, assetID string,
) (*CreateAddressResponse, error) {
	var vaultResponse *vault.VaultAccountResponse
	{
		var err error
		if vaultAccountID == nil {
			if strings.TrimSpace(userID) == "" {
				return nil, errors.New("user id is mandatory")
			}

			vaultResponse, err = v.vault.CreateVaultAccount(ctx, vault.CreateVaultParams{Name: userID, CustomerRefID: userID})
			if err != nil {
				return nil, err
			}
		}

		if vaultAccountID != nil {
			vaultResponse, err = v.vault.FindVaultAccountByID(ctx, *vaultAccountID)
			if err != nil {
				return nil, err
			}
		}
	}

	{ // wallet handling
		var vaultAlreadyHasAssetWallet bool

		for _, asset := range vaultResponse.Assets {
			if strings.EqualFold(asset.ID, assetID) {
				vaultAlreadyHasAssetWallet = true
			}
		}

		if !vaultAlreadyHasAssetWallet {
			_, err := v.vault.CreateAssetWallet(ctx, vaultResponse.ID, assetID)
			if err != nil {
				return nil, err
			}
		}
	}

	rsp, err := v.vault.CreateAssetAddress(ctx, vaultResponse.ID, assetID)
	if err != nil {
		return nil, err
	}

	return &CreateAddressResponse{
		VaultID:           vaultResponse.ID,
		Address:           rsp.Address,
		LegacyAddress:     rsp.LegacyAddress,
		EnterpriseAddress: rsp.EnterpriseAddress,
		Tag:               rsp.Tag,
		Bip44AddressIndex: rsp.Bip44AddressIndex,
	}, err
}

// This is a streamlined minimalistic version of Create Transaction
// With main aim of sending crypto to extrernal address only.
// It makes use of only whats needed to withdraw/send crypto to a given address
// See: https://developers.fireblocks.com/reference/post_transactions
func (v *CustomFlowSVC) SendToExternalAddress(
	ctx context.Context, p SendToExternalAddress, opts ...client.Option,
) (*transaction.CreateTxnResponse, error) {
	param := transaction.CreateTransactionParams{
		Operation: "TRANSFER",
		Source: struct {
			Type     string `json:"type"`
			ID       string `json:"id"`
			WalletID string `json:"walletId,omitempty"`
			Name     string `json:"name"`
		}{
			Type: "VAULT_ACCOUNT",
			ID:   p.SourceVaultID,
		},
		Destination: transaction.CreateTransactionParamsDestination{
			Type: "ONE_TIME_ADDRESS",
			OneTimeAddress: struct {
				Address string `json:"address"`
			}{
				Address: p.DestinationAddress,
			},
		},
		Note:               p.Note,
		ExternalTxID:       p.ExternalTxID,
		AssetID:            p.AssetID,
		TreatAsGrossAmount: p.TreatAsGrossAmount,
		Amount:             p.Amount,
		// FeeLevel:           p.FeeLevel,
		// Fee:                p.Fee,
		FailOnLowFee:  p.FailOnLowFee,
		CustomerRefID: p.CustomerRefID,
	}

	return v.txn.CreateTransaction(ctx, param, opts...)
}

func (v *CustomFlowSVC) SweepAssets(ctx context.Context, assetId string, minSweepAmount float64, treasuryVaultAccountId string) error {

	// Fetch accounts with pagination handling
	accounts, err := v.fetchAccounts(ctx, assetId, minSweepAmount)
	if err != nil {
		return err
	}

	workerCount := 50

	// Batch processing with configurable worker count
	if err := v.processAccounts(ctx, assetId, treasuryVaultAccountId, accounts, workerCount); err != nil {
		return err
	}

	return nil
}

// Fetches accounts with pagination handling
func (v *CustomFlowSVC) fetchAccounts(ctx context.Context, assetId string, minSweepAmount float64) ([]vault.ListVaultsResponseAccount, error) {
	var accounts []vault.ListVaultsResponseAccount
	params := vault.ListVaultsParams{
		NamePrefix:         "ZAB_USER_",
		MinAmountThreshold: minSweepAmount,
		AssetId:            assetId,
		OrderBy:            "DESC",
		Limit:              500,
	}

	for {
		vaultList, err := v.vault.ListVaults(ctx, params)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, vaultList.Accounts...)
		if vaultList.Paging.After == "" {
			break
		}
		params.After = vaultList.Paging.After
	}

	return accounts, nil
}

// Processes accounts in batches using worker goroutines
func (v *CustomFlowSVC) processAccounts(ctx context.Context, assetId string, treasuryVaultAccountId string, accounts []vault.ListVaultsResponseAccount, workerCount int) error {

	batches := split(accounts, workerCount)

	var errs []error
	for _, batch := range batches {
		if err := func(b Batch) []error {
			err := v.processBatch(ctx, assetId, treasuryVaultAccountId, b.Accounts)
			if err != nil {
				return err
			}
			return nil
		}(batch); err != nil {
			errs = append(errs, err...)
		}
	}

	return errors.Join(errs...)
}

var extractAssetFn = func(assetId string, acc vault.ListVaultsResponseAccount) (vault.ListVaultsResponseAccountAssets, bool) {
	for _, asset := range acc.Assets {
		if asset.ID == assetId {
			return asset, true
		}
	}
	return vault.ListVaultsResponseAccountAssets{}, false
}

// Processes a batch of accounts in parallel
func (v *CustomFlowSVC) processBatch(ctx context.Context, assetId string, treasuryVaultAccountId string, accounts []vault.ListVaultsResponseAccount) []error {
	wg := sync.WaitGroup{}
	result := make(chan error)

	for _, account := range accounts {
		wg.Add(1)
		go func(assetId string, acc vault.ListVaultsResponseAccount) {
			defer wg.Done()

			asset, found := extractAssetFn(assetId, acc)
			if !found {
				result <- fmt.Errorf("asset %s not found in account %s", assetId, acc.ID)
				return
			}

			_, err := v.txn.CreateTransaction(ctx, transaction.CreateTransactionParams{
				Operation: "TRANSFER",
				AssetID:   assetId,
				Source: transaction.CreateTransactionParamsSource{
					Type: "VAULT_ACCOUNT",
					ID:   acc.ID,
				},
				Destination: transaction.CreateTransactionParamsDestination{
					Type: "VAULT_ACCOUNT",
					ID:   treasuryVaultAccountId,
				},
				Amount: asset.Total,
				Note:   "Sweeping assets to treasury",
			})
			result <- err
		}(assetId, account)
	}

	go func() {
		wg.Wait()
		close(result)
	}()

	var errs []error
	for err := range result {
		if err == nil {
			continue
		}
		errs = append(errs, err)
	}

	return errs
}

type Batch struct {
	Accounts []vault.ListVaultsResponseAccount
	Len      int
}

func split(accounts []vault.ListVaultsResponseAccount, batchSize int) []Batch {

	if len(accounts) < 1 || batchSize < 1 {
		return []Batch{}
	}

	// Handle base case: less than or equal to batch size
	if len(accounts) <= batchSize {
		return []Batch{
			{
				Accounts: accounts,
				Len:      len(accounts),
			},
		}
	}

	var batches []Batch
	// Loop through accounts and create batches
	for i := 0; i < len(accounts); i += batchSize {
		// Calculate remaining elements and ensure we don't go out of bounds
		remaining := len(accounts) - i
		batchSize = min(batchSize, remaining)

		// Create a new batch with the current slice of accounts
		batch := Batch{
			Accounts: accounts[i : i+batchSize],
			Len:      batchSize,
		}
		batches = append(batches, batch)
	}

	return batches
}

// Helper function to get minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
