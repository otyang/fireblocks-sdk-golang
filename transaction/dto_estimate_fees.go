package transaction

type EstimateFeeParams struct {
	Operation string `json:"operation"`
	Source    struct {
		Type string `json:"type"`
	} `json:"source"`
	Destination struct {
		Type string `json:"type"`
	} `json:"destination"`
	AssetID            string `json:"assetId"`
	Amount             string `json:"amount"`
	TreatAsGrossAmount bool   `json:"treatAsGrossAmount"`
}

type EstimateFeeResponse struct {
	Low struct {
		FeePerByte  string `json:"feePerByte"`
		GasPrice    string `json:"gasPrice"`
		GasLimit    string `json:"gasLimit"`
		NetworkFee  string `json:"networkFee"`
		BaseFee     string `json:"baseFee"`
		PriorityFee string `json:"priorityFee"`
	} `json:"low"`
	Medium struct {
		FeePerByte  string `json:"feePerByte"`
		GasPrice    string `json:"gasPrice"`
		GasLimit    string `json:"gasLimit"`
		NetworkFee  string `json:"networkFee"`
		BaseFee     string `json:"baseFee"`
		PriorityFee string `json:"priorityFee"`
	} `json:"medium"`
	High struct {
		FeePerByte  string `json:"feePerByte"`
		GasPrice    string `json:"gasPrice"`
		GasLimit    string `json:"gasLimit"`
		NetworkFee  string `json:"networkFee"`
		BaseFee     string `json:"baseFee"`
		PriorityFee string `json:"priorityFee"`
	} `json:"high"`
}
