package vault

type CreateVaultParams struct {
	// Account Name
	Name string `json:"name"`
	// Optional - if true, the created account and all related transactions will not be shown on Fireblocks console
	// 	Set to true by default, hides this vault account from appearing in the Fireblocks Console.
	// This is the best practice when creating intermediate deposit vault accounts for your users
	// as it helps reduce visual clutter and improves UI loading time.
	// The best practice is configuring this setting so that only your omnibus account
	// and another operational vault account (or multiple) are visible in the Fireblocks Console.
	HiddenOnUI bool `json:"hiddenOnUI"`
	// Optional - Sets the autoFuel property of the vault account
	AutoFuel bool `json:"autoFuel"`
	// Optional - Sets a customer reference ID
	CustomerRefID string `json:"customerRefId"`
}

type VaultAccountResponse struct {
	ID            string  `json:"id"`
	Name          string  `json:"name"`
	Assets        []Asset `json:"assets"`
	HiddenOnUI    bool    `json:"hiddenOnUI"`
	CustomerRefID string  `json:"customerRefId"`
	AutoFuel      bool    `json:"autoFuel"`
}

type Asset struct {
	ID                   string `json:"id"`
	Total                string `json:"total"`
	Available            string `json:"available"`
	Pending              string `json:"pending"`
	Frozen               string `json:"frozen"`
	LockedAmount         string `json:"lockedAmount"`
	Staked               string `json:"staked"`
	TotalStakedCPU       int    `json:"totalStakedCPU"`
	TotalStakedNetwork   string `json:"totalStakedNetwork"`
	SelfStakedCPU        string `json:"selfStakedCPU"`
	SelfStakedNetwork    string `json:"selfStakedNetwork"`
	PendingRefundCPU     string `json:"pendingRefundCPU"`
	PendingRefundNetwork string `json:"pendingRefundNetwork"`
	BlockHeight          string `json:"blockHeight"`
	BlockHash            string `json:"blockHash"`
	RewardsInfo          struct {
		PendingRewards string `json:"pendingRewards"`
	} `json:"rewardsInfo"`
}

type CreateAssetResponse struct {
	ID                string `json:"id"`
	Address           string `json:"address"`
	LegacyAddress     string `json:"legacyAddress"`
	EnterpriseAddress string `json:"enterpriseAddress"`
	Tag               string `json:"tag"`
	EosAccountName    string `json:"eosAccountName"`
	Status            string `json:"status"`
	ActivationTxID    string `json:"activationTxId"`
}

type CreateAddressResponse struct {
	Address           string `json:"address"`
	LegacyAddress     string `json:"legacyAddress"`
	EnterpriseAddress string `json:"enterpriseAddress"`
	Tag               string `json:"tag"`
	Bip44AddressIndex int    `json:"bip44AddressIndex"`
}



type VaultAssetBalanceResponse struct {
	ID           string          `json:"id"`
	Total        decimal.Decimal `json:"total"`
	Balance      decimal.Decimal `json:"balance"`
	LockedAmount decimal.Decimal `json:"lockedAmount"`
	Available    decimal.Decimal `json:"available"`
	Pending      decimal.Decimal `json:"pending"`
	Frozen       decimal.Decimal `json:"frozen"`
	Staked       decimal.Decimal `json:"staked"`
	BlockHeight  string          `json:"blockHeight"`
}

type ListVaultsParams struct {
	NamePrefix         string  `json:"namePrefix"`
	NameSuffix         string  `json:"nameSuffix"`
	MinAmountThreshold float64 `json:"minAmountThreshold"`
	AssetId            string  `json:"assetId"`
	OrderBy            OrderBy `json:"orderBy"`
	Limit              int     `json:"limit"` // min of 1 &  max of 500
	Before             string  `json:"before"`
	After              string  `json:"after"`
}

type OrderBy string

const (
	OrderByAsc  OrderBy = "ASC"
	OrderByDesc OrderBy = "DESC"
)

type ListVaultsResponseAccountAssets struct {
	ID           string `json:"id"`
	Total        string `json:"total"`
	Available    string `json:"available"`
	Pending      string `json:"pending"`
	Frozen       string `json:"frozen"`
	LockedAmount string `json:"lockedAmount"`
	Staked       string `json:"staked"`
	// TotalStakedCPU       int    `json:"totalStakedCPU"`
	// TotalStakedNetwork   string `json:"totalStakedNetwork"`
	// SelfStakedCPU        string `json:"selfStakedCPU"`
	// SelfStakedNetwork    string `json:"selfStakedNetwork"`
	PendingRefundCPU     string `json:"pendingRefundCPU"`
	PendingRefundNetwork string `json:"pendingRefundNetwork"`
	BlockHeight          string `json:"blockHeight"`
	BlockHash            string `json:"blockHash"`
	RewardsInfo          struct {
		PendingRewards string `json:"pendingRewards"`
	} `json:"rewardsInfo"`
}

type ListVaultsResponseAccount struct {
	ID            string                            `json:"id"`
	Name          string                            `json:"name"`
	Assets        []ListVaultsResponseAccountAssets `json:"assets"`
	HiddenOnUI    bool                              `json:"hiddenOnUI"`
	CustomerRefID string                            `json:"customerRefId"`
	AutoFuel      bool                              `json:"autoFuel"`
}

type ListVaultsResponse struct {
	Accounts []ListVaultsResponseAccount `json:"accounts"`
	Paging   struct {
		Before string `json:"before"`
		After  string `json:"after"`
	} `json:"paging"`
	PreviousURL string `json:"previousUrl"`
	NextURL     string `json:"nextUrl"`
}
