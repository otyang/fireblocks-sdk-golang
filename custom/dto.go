package custom

type CreateAddressResponse struct {
	VaultID           string
	Address           string `json:"address"`
	LegacyAddress     string `json:"legacyAddress"`
	EnterpriseAddress string `json:"enterpriseAddress"`
	Tag               string `json:"tag"`
	Bip44AddressIndex int    `json:"bip44AddressIndex"`
}

type SendToExternalAddress struct {
	// Destination specifies a one-time address of the intended recipient of the funds.
	DestinationAddress string

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

	// Optional but highly recommended parameter. Fireblocks will reject future transactions with same ID.
	// You should set this to a unique ID representing the transaction,
	// to avoid submitting the same transaction twice. This helps with
	// cases where submitting the transaction responds with an error code
	// due to Internet interruptions, but the transaction was actually sent and processed.
	// the unique identifier of the transaction outside of Fireblocks
	ExternalTxID string `json:"externalTxId"`

	// Optional message attached to the transaction. Not sent to the blockchain.
	// Just used to describe the transaction at your Fireblocks workspace.
	Note string `json:"note"`
}
