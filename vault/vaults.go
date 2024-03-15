package vault

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/otyang/fireblocks/client"
)

type VaultService struct {
	client *client.Client
}

func New(client *client.Client) *VaultService {
	return &VaultService{
		client: client,
	}
}

// Creates a new vault account with the requested name.
// See: https://developers.fireblocks.com/reference/post_vault-accounts
func (v *VaultService) CreateVaultAccount(ctx context.Context, params CreateVaultParams) (*VaultAccountResponse, error) {
	var (
		path       = "/v1/vault/accounts"
		apiSuccess VaultAccountResponse
	)

	if strings.TrimSpace(params.Name) == "" {
		return nil, errors.New("vault name is required")
	}

	p := CreateVaultParams{
		Name:          "ZAB_USER_" + params.Name,
		HiddenOnUI:    true,
		AutoFuel:      true,
		CustomerRefID: params.CustomerRefID,
	}

	_, err := v.client.MakeRequest(ctx, "post", path, p, &apiSuccess)
	if err != nil {
		return nil, err
	}

	return &apiSuccess, nil
}

// Creates a wallet for a specific asset in a vault account.
// See: https://developers.fireblocks.com/reference/post_vault-accounts-vaultaccountid-assetid
func (v *VaultService) CreateAssetWallet(ctx context.Context, vaultAccountID, assetID string) (*CreateAssetResponse, error) {
	var (
		path       = fmt.Sprintf("/v1/vault/accounts/%s/%s", vaultAccountID, assetID)
		apiSuccess CreateAssetResponse
	)

	_, err := v.client.MakeRequest(ctx, "post", path, nil, &apiSuccess)
	if err != nil {
		return nil, err
	}

	return &apiSuccess, nil
}

// CreateAssetAddress Creates a wallet for a specific asset in a vault account.
// See: https://developers.fireblocks.com/reference/post_vault-accounts-vaultaccountid-assetid
func (v *VaultService) CreateAssetAddress(ctx context.Context, vaultAccountID, assetID string) (*CreateAddressResponse, error) {
	var (
		path       = fmt.Sprintf("/v1/vault/accounts/%s/%s/addresses", vaultAccountID, assetID)
		apiSuccess CreateAddressResponse
	)

	_, err := v.client.MakeRequest(ctx, "post", path, nil, &apiSuccess)
	if err != nil {
		return nil, err
	}

	return &apiSuccess, nil
}

// Returns the requested vault account.
// See: https://developers.fireblocks.com/reference/get_vault-accounts-vaultaccountid
func (v *VaultService) FindVaultAccountByID(ctx context.Context, vaultAccountID string) (*VaultAccountResponse, error) {
	var (
		path       = fmt.Sprintf("/v1/vault/accounts/%s", vaultAccountID)
		apiSuccess VaultAccountResponse
	)

	_, err := v.client.MakeRequest(ctx, "get", path, nil, &apiSuccess)
	if err != nil {
		return nil, err
	}

	return &apiSuccess, nil
}

// https://developers.fireblocks.com/reference/get_vault-accounts-paged
func (v *VaultService) ListVaults(ctx context.Context, params ListVaultsParams) (*ListVaultsResponse, error) {
	var (
		path       = "/v1/vault/accounts_paged"
		apiSuccess ListVaultsResponse
	)

	var qp []string
	if params.Limit < 1 {
		params.Limit = 1
	}

	if params.Limit > 500 {
		params.Limit = 500
	}
	qp = append(qp, fmt.Sprintf("limit=%v", params.Limit))

	if params.NamePrefix != "" {
		qp = append(qp, fmt.Sprintf("namePrefix=%s", params.NamePrefix))
	}
	if params.NameSuffix != "" {
		qp = append(qp, fmt.Sprintf("nameSuffix=%s", params.NameSuffix))
	}
	if params.MinAmountThreshold != 0 {
		qp = append(qp, fmt.Sprintf("minAmountThreshold=%v", params.MinAmountThreshold))
	}
	if params.AssetId != "" {
		qp = append(qp, fmt.Sprintf("assetId=%s", params.AssetId))
	}
	if params.OrderBy != "" {
		qp = append(qp, fmt.Sprintf("orderBy=%s", params.OrderBy))
	}
	if params.Before != "" {
		qp = append(qp, params.Before)
	}
	if params.After != "" {
		qp = append(qp, params.After)
	}

	if len(qp) > 0 {
		path = fmt.Sprintf("%s?%s", path, strings.Join(qp, "&"))
	}

	_, err := v.client.MakeRequest(ctx, "GET", path, nil, &apiSuccess)
	if err != nil {
		return nil, err
	}

	return &apiSuccess, nil
}
