package transaction

type TransactionResponse struct {
	ID           string `json:"id"`
	ExternalTxID string `json:"externalTxId"`
	Status       string `json:"status"`
	SubStatus    string `json:"subStatus"`
	TxHash       string `json:"txHash"`
	Operation    string `json:"operation"`
	Note         string `json:"note"`
	AssetID      string `json:"assetId"`
	Source       struct {
		Type     string `json:"type"`
		SubType  string `json:"subType"`
		ID       string `json:"id"`
		Name     string `json:"name"`
		WalletID string `json:"walletId"`
	} `json:"source"`
	SourceAddress string `json:"sourceAddress"`
	Tag           string `json:"tag"`
	Destination   struct {
		Type     string `json:"type"`
		SubType  string `json:"subType"`
		ID       string `json:"id"`
		Name     string `json:"name"`
		WalletID string `json:"walletId"`
	} `json:"destination"`
	Destinations []struct {
		Destination struct {
			Type     string `json:"type"`
			SubType  string `json:"subType"`
			ID       string `json:"id"`
			Name     string `json:"name"`
			WalletID string `json:"walletId"`
		} `json:"destination"`
		Amount             string `json:"amount"`
		AmountUSD          string `json:"amountUSD"`
		AmlScreeningResult struct {
			Provider string   `json:"provider"`
			Payload  struct{} `json:"payload"`
		} `json:"amlScreeningResult"`
		AuthorizationInfo struct {
			AllowOperatorAsAuthorizer bool   `json:"allowOperatorAsAuthorizer"`
			Logic                     string `json:"logic"`
			Groups                    []struct {
				Th    int `json:"th"`
				Users struct {
					AdditionalProp string `json:"additionalProp"`
				} `json:"users"`
			} `json:"groups"`
		} `json:"authorizationInfo"`
	} `json:"destinations"`
	DestinationAddress            string `json:"destinationAddress"`
	DestinationAddressDescription string `json:"destinationAddressDescription"`
	DestinationTag                string `json:"destinationTag"`
	ContractCallDecodedData       struct {
		ContractName  string     `json:"contractName"`
		FunctionCalls []struct{} `json:"functionCalls"`
	} `json:"contractCallDecodedData"`
	AmountInfo struct {
		Amount          string `json:"amount"`
		RequestedAmount string `json:"requestedAmount"`
		NetAmount       string `json:"netAmount"`
		AmountUSD       string `json:"amountUSD"`
	} `json:"amountInfo"`
	TreatAsGrossAmount bool `json:"treatAsGrossAmount"`
	FeeInfo            struct {
		NetworkFee string `json:"networkFee"`
		ServiceFee string `json:"serviceFee"`
		GasPrice   string `json:"gasPrice"`
	} `json:"feeInfo"`
	FeeCurrency    string `json:"feeCurrency"`
	NetworkRecords []struct {
		Source struct {
			Type     string `json:"type"`
			SubType  string `json:"subType"`
			ID       string `json:"id"`
			Name     string `json:"name"`
			WalletID string `json:"walletId"`
		} `json:"source"`
		Destination struct {
			Type     string `json:"type"`
			SubType  string `json:"subType"`
			ID       string `json:"id"`
			Name     string `json:"name"`
			WalletID string `json:"walletId"`
		} `json:"destination"`
		TxHash             string `json:"txHash"`
		NetworkFee         string `json:"networkFee"`
		AssetID            string `json:"assetId"`
		NetAmount          string `json:"netAmount"`
		IsDropped          bool   `json:"isDropped"`
		Type               string `json:"type"`
		DestinationAddress string `json:"destinationAddress"`
		SourceAddress      string `json:"sourceAddress"`
		AmountUSD          string `json:"amountUSD"`
		Index              int    `json:"index"`
		RewardInfo         struct {
			SrcRewards  string `json:"srcRewards"`
			DestRewards string `json:"destRewards"`
		} `json:"rewardInfo"`
	} `json:"networkRecords"`
	CreatedAt         int      `json:"createdAt"`
	LastUpdated       int      `json:"lastUpdated"`
	CreatedBy         string   `json:"createdBy"`
	SignedBy          []string `json:"signedBy"`
	RejectedBy        string   `json:"rejectedBy"`
	AuthorizationInfo struct {
		AllowOperatorAsAuthorizer bool   `json:"allowOperatorAsAuthorizer"`
		Logic                     string `json:"logic"`
		Groups                    []struct {
			Th    int `json:"th"`
			Users struct {
				AdditionalProp string `json:"additionalProp"`
			} `json:"users"`
		} `json:"groups"`
	} `json:"authorizationInfo"`
	ExchangeTxID       string `json:"exchangeTxId"`
	CustomerRefID      string `json:"customerRefId"`
	AmlScreeningResult struct {
		Provider string   `json:"provider"`
		Payload  struct{} `json:"payload"`
	} `json:"amlScreeningResult"`
	ExtraParameters struct{} `json:"extraParameters"`
	SignedMessages  struct {
		Content        string `json:"content"`
		Algorithm      string `json:"algorithm"`
		DerivationPath []int  `json:"derivationPath"`
		Signature      struct {
			FullSig string `json:"fullSig"`
			R       string `json:"r"`
			S       string `json:"s"`
			V       int    `json:"v"`
		} `json:"signature"`
		PublicKey string `json:"publicKey"`
	} `json:"signedMessages"`
	NumOfConfirmations int `json:"numOfConfirmations"`
	BlockInfo          struct {
		BlockHeight string `json:"blockHeight"`
		BlockHash   string `json:"blockHash"`
	} `json:"blockInfo"`
	Index      int `json:"index"`
	RewardInfo struct {
		SrcRewards  string `json:"srcRewards"`
		DestRewards string `json:"destRewards"`
	} `json:"rewardInfo"`
	SystemMessages struct {
		Type    string `json:"type"`
		Message string `json:"message"`
	} `json:"systemMessages"`
	AddressType string `json:"addressType"`
}
