package transaction

type CreateTxnResponse struct {
	ID             string `json:"id"`
	Status         string `json:"status"`
	SystemMessages struct {
		Type    string `json:"type"`
		Message string `json:"message"`
	} `json:"systemMessages"`
}

// TransactionParams represents the parameters for a transaction.
// for further reference: https://developers.fireblocks.com/reference/post_transactions
type CreateTransactionParams struct {
	// Operation specifies the type of transaction to be performed.
	// Valid values include: "TRANSFER", "MINT" or "BURN" etc.
	Operation string `json:"operation"`

	// Source describes the origin of the funds for the transaction.
	Source struct {
		// Type indicates the type of source.
		Type     string `json:"type"`
		ID       string `json:"id"`
		WalletID string `json:"walletId"`
		Name     string `json:"name"`
	} `json:"source"`

	// Destination specifies the intended recipient of the funds.
	Destination struct {
		Type           string `json:"type"`
		OneTimeAddress struct {
			// one-time address to receive funds.
			Address string `json:"address"`
		} `json:"oneTimeAddress"`
	} `json:"destination"`

	// Optional message attached to the transaction. Not sent to the blockchain.
	// Just used to describe the transaction at your Fireblocks workspace.
	Note string `json:"note"`

	// Optional but highly recommended parameter. Fireblocks will reject future transactions with same ID.
	// You should set this to a unique ID representing the transaction,
	// to avoid submitting the same transaction twice. This helps with
	// cases where submitting the transaction responds with an error code
	// due to Internet interruptions, but the transaction was actually sent and processed.
	// the unique identifier of the transaction outside of Fireblocks
	ExternalTxID string `json:"externalTxId"`

	// AssetID identifies the asset being transacted (e.g., currency).
	AssetID string `json:"assetId"`

	// TreatAsGrossAmount specifies whether the amount should be treated as gross (inclusive of fees).
	TreatAsGrossAmount bool `json:"treatAsGrossAmount"`

	// Amount is the value of the transaction in the specified asset.
	Amount string `json:"amount"`

	// Defines the blockchain fee level which will be paid for the
	// transaction (only for Ethereum and UTXO-based blockchains).
	// Set to MEDIUM by default. Valid values are ("LOW", "MEDIUM", "HIGH").
	FeeLevel string `json:"feeLevel"` // *

	// Fee is the estimated transaction fee, in the asset's smallest unit (Satoshi, Latoshi, etc)
	Fee string `json:"fee"` // *

	// FailOnLowFee specifies whether the transaction should fail
	// if the estimated fee is too low. to avoid getting stuck with no confirmation
	FailOnLowFee bool `json:"failOnLowFee"`

	// NetworkFee is the calculated network fee for the transaction.
	NetworkFee string `json:"networkFee"`

	// PriorityFee is the additional fee paid to prioritize transaction processing.
	PriorityFee string `json:"priorityFee"`

	// MaxFee is the maximum fee the user is willing to pay for the transaction.
	MaxFee string `json:"maxFee"`

	// GasLimit is the maximum amount of gas that can be used for the transaction.
	GasLimit string `json:"gasLimit"`

	// ReplaceTxByHash specifies the hash of a previous transaction to be replaced.
	ReplaceTxByHash string `json:"replaceTxByHash"`

	// CustomerRefID is an optional reference ID provided by the customer.
	CustomerRefID string `json:"customerRefId"`

	// GasPrice is the price per unit of gas for the transaction.
	GasPrice string `json:"gasPrice"`
}
