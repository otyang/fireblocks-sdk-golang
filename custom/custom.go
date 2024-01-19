package custom

import (
	"context"
	"errors"
	"strings"

	"github.com/otyang/fireblocks/client"
	"github.com/otyang/fireblocks/transaction"
	"github.com/otyang/fireblocks/vault"
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

			vaultResponse, err = v.vault.CreateVaultAccount(ctx, vault.CreateVaultParams{Name: userID})
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

		if vaultAlreadyHasAssetWallet == false {
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
	ctx context.Context, p SendToExternalAddress,
) (*transaction.CreateTxnResponse, error) {
	param := transaction.CreateTransactionParams{
		Operation: "TRANSFER",
		Source: struct {
			Type     string `json:"type"`
			ID       string `json:"id"`
			WalletID string `json:"walletId"`
			Name     string `json:"name"`
		}{
			Type: "VAULT_ACCOUNT",
		},
		Destination: struct {
			Type           string `json:"type"`
			OneTimeAddress struct {
				Address string `json:"address"`
			} `json:"oneTimeAddress"`
		}{
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
		FeeLevel:           p.FeeLevel,
		Fee:                p.Fee,
		FailOnLowFee:       p.FailOnLowFee,
	}

	return v.txn.CreateTransaction(ctx, param)
}
