package webhook

import (
	"context"
	"fmt"

	"github.com/otyang/fireblocks/client"
	"github.com/otyang/fireblocks/transaction"
)

type WebHookService struct {
	client         *client.Client
	transactionSvc *transaction.TransactionService
}

func New(client *client.Client, transactionService *transaction.TransactionService) *WebHookService {
	return &WebHookService{
		client:         client,
		transactionSvc: transactionService,
	}
}

// ResendFailedWebhook sends a request to Fireblocks to resend all failed webhooks
// See: https://developers.fireblocks.com/reference/post_webhooks-resend
func (w *WebHookService) ResendFailedWebhook(ctx context.Context) error {
	var (
		path       = "/v1/webhooks/resend"
		apiSuccess *ResendFailedWebhookResponse
	)

	_, err := w.client.MakeRequest(ctx, "post", path, nil, apiSuccess)
	return err
}

// ResendFailedWebhookByTransactionID sends a request to Fireblocks to resend a failed webhook by transaction ID
// See: https://developers.fireblocks.com/reference/post_webhooks-resend-txid
func (w *WebHookService) ResendFailedWebhookByTransactionID(
	ctx context.Context, txnID string,
) (*ResendFailedWebhookResponse, error) {
	var (
		path       = fmt.Sprintf("/v1/webhooks/resend/%s", txnID)
		apiSuccess ResendFailedWebhookResponse
	)

	_, err := w.client.MakeRequest(ctx, "post", path, nil, &apiSuccess)
	return nil, err
}

// VerifyWebhookTransaction verifies a webhook transaction by checking its existence on fireblocks
func (w *WebHookService) VerifyWebhookTransaction(ctx context.Context, webhookPayload WebhookPayload) error {
	_, err := w.transactionSvc.FindByFireblocksTransactionId(ctx, webhookPayload.Data.ID)
	return err
}
