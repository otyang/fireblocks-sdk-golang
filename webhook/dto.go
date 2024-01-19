package webhook

type ResendFailedWebhookResponse struct {
	MessagesCount int `json:"messagesCount"`
}

type WebhookPayload struct {
	Type      string `json:"type"`
	TenantID  string `json:"tenantId"`
	Timestamp int64  `json:"timestamp"`
	Data      struct {
		ID          string `json:"id"`
		CreatedAt   int64  `json:"createdAt"`
		LastUpdated int64  `json:"lastUpdated"`
		AssetID     string `json:"assetId"`
		Source      struct {
			ID      string `json:"id"`
			Type    string `json:"type"`
			Name    string `json:"name"`
			SubType string `json:"subType"`
		} `json:"source"`
		Destination struct {
			ID      string `json:"id"`
			Type    string `json:"type"`
			Name    string `json:"name"`
			SubType string `json:"subType"`
		} `json:"destination"`
		Amount                        int     `json:"amount"`
		NetworkFee                    float64 `json:"networkFee"`
		NetAmount                     int     `json:"netAmount"`
		SourceAddress                 string  `json:"sourceAddress"`
		DestinationAddress            string  `json:"destinationAddress"`
		DestinationAddressDescription string  `json:"destinationAddressDescription"`
		DestinationTag                string  `json:"destinationTag"`
		Status                        string  `json:"status"`
		TxHash                        string  `json:"txHash"`
		SubStatus                     string  `json:"subStatus"`
		SignedBy                      []any   `json:"signedBy"`
		CreatedBy                     string  `json:"createdBy"`
		RejectedBy                    string  `json:"rejectedBy"`
		AmountUSD                     int     `json:"amountUSD"`
		AddressType                   string  `json:"addressType"`
		Note                          string  `json:"note"`
		ExchangeTxID                  string  `json:"exchangeTxId"`
		RequestedAmount               int     `json:"requestedAmount"`
		FeeCurrency                   string  `json:"feeCurrency"`
		Operation                     string  `json:"operation"`
		CustomerRefID                 string  `json:"customerRefId"`
		NumOfConfirmations            int     `json:"numOfConfirmations"`
		AmountInfo                    struct {
			Amount          string `json:"amount"`
			RequestedAmount string `json:"requestedAmount"`
			NetAmount       string `json:"netAmount"`
			AmountUSD       any    `json:"amountUSD"`
		} `json:"amountInfo"`
		FeeInfo struct {
			NetworkFee string `json:"networkFee"`
		} `json:"feeInfo"`
		Destinations []any  `json:"destinations"`
		ExternalTxID string `json:"externalTxId"`
		BlockInfo    struct {
			BlockHeight string `json:"blockHeight"`
			BlockHash   string `json:"blockHash"`
		} `json:"blockInfo"`
		SignedMessages     []any `json:"signedMessages"`
		AmlScreeningResult struct {
			ScreeningStatus string `json:"screeningStatus"`
			BypassReason    string `json:"bypassReason"`
			Timestamp       int64  `json:"timestamp"`
		} `json:"amlScreeningResult"`
		AssetType string `json:"assetType"`
	} `json:"data"`
}
