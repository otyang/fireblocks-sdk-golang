package transaction

import (
	"context"
	"fmt"

	"github.com/otyang/fireblocks/client"
)

type TransactionService struct {
	client *client.Client
}

func New(client *client.Client) *TransactionService {
	return &TransactionService{
		client: client,
	}
}

// Estimates the transaction fee for a transaction request.
// See: https://developers.fireblocks.com/reference/post_transactions-estimate-fee
func (v *TransactionService) EstimateFeeForSendingToExternalAddress(
	ctx context.Context, assetID string, amount string, treatAsGrossAmount bool,
) (*EstimateFeeResponse, error) {
	var (
		path       = "/v1/vault/accounts"
		apiSuccess *EstimateFeeResponse
	)

	params := EstimateFeeParams{
		Operation: "TRANSFER",
		Source: struct {
			Type string `json:"type"`
		}{
			Type: "VAULT_ACCOUNT",
		},
		Destination: struct {
			Type string `json:"type"`
		}{
			Type: "ONE_TIME_ADDRESS",
		},
		AssetID:            assetID,
		Amount:             amount,
		TreatAsGrossAmount: treatAsGrossAmount,
	}

	_, err := v.client.MakeRequest("post", path, params, apiSuccess)
	if err != nil {
		return nil, err
	}

	return apiSuccess, nil
}

// Returns a transaction by ID.
// See: https://developers.fireblocks.com/reference/get_transactions-txid
func (v *TransactionService) FindByFireblocksTransactionId(ctx context.Context, txID string) (*TransactionResponse, error) {
	var (
		path       = fmt.Sprintf("/v1/transactions/%s", txID)
		apiSuccess *TransactionResponse
	)

	_, err := v.client.MakeRequest("get", path, nil, apiSuccess)
	if err != nil {
		return nil, err
	}

	return apiSuccess, nil
}

// Returns transaction by external transaction ID.
// See: https://developers.fireblocks.com/reference/get_transactions-external-tx-id-externaltxid
func (v *TransactionService) FindByExternalTransactionId(ctx context.Context, txID string) (*TransactionResponse, error) {
	var (
		path       = fmt.Sprintf("/v1/transactions/external_tx_id/%s", txID)
		apiSuccess *TransactionResponse
	)

	_, err := v.client.MakeRequest("get", path, nil, apiSuccess)
	if err != nil {
		return nil, err
	}

	return apiSuccess, nil
}

// Creates a new transaction.
// See: https://developers.fireblocks.com/reference/post_transactions
func (v *TransactionService) CreateTransaction(ctx context.Context, p CreateTransactionParams) (*CreateTxnResponse, error) {
	var (
		path       = "/v1/transactions"
		apiSuccess *CreateTxnResponse
	)

	_, err := v.client.MakeRequest("post", path, p, apiSuccess)
	if err != nil {
		return nil, err
	}

	return apiSuccess, nil
}
