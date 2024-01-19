package fireblocks

import (
	"github.com/otyang/fireblocks/client"
	"github.com/otyang/fireblocks/custom"
	"github.com/otyang/fireblocks/transaction"
	"github.com/otyang/fireblocks/vault"
	"github.com/otyang/fireblocks/webhook"
)

const (
// BaseURL   = "https://sandbox-api.fireblocks.io/v1"
)

type SDK struct {
	Vault       *vault.VaultService
	Webhook     *webhook.WebHookService
	Transaction *transaction.TransactionService
	Custom      *custom.CustomFlowSVC
}

func New(apiKey, privateKey, baseURL string) (*SDK, error) {
	client := client.New(apiKey, privateKey, baseURL, nil)

	vaultSvc := vault.New(client)
	txnSvc := transaction.New(client)

	return &SDK{
		Vault:       vaultSvc,
		Webhook:     webhook.New(client, txnSvc),
		Transaction: txnSvc,
		Custom:      custom.New(client, vaultSvc),
	}, nil
}
