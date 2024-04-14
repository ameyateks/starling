package types

type StarlingAccountInfo struct {
	Accounts []StarlingAccount `json:"accounts"`
}

type StarlingAccount struct {
	AccountUid      string `json:"accountUid"`
	AccountType     string `json:"accountType"`
	DefaultCategory string `json:"defaultCategory"`
	Currency        string `json:"currency"`
	CreatedAt       string `json:"createdAt"`
	Name            string `json:"name"`
}


type AccountAndCategoryUid struct {
	AccountUid  string `json:"accountUid"`
	CategoryUid string `json:"categoryUid"`
}

type StarlingAccountResp struct {
	AccountUid string `json:"accountUid"`
}

type StarlingBalance struct {
	ClearedBalance        SignedCurrencyAndAmount `json:"clearedBalance"`
	EffectiveBalance      SignedCurrencyAndAmount `json:"effectiveBalance"`
	PendingTransactions   SignedCurrencyAndAmount `json:"pendingTransactions"`
	AcceptedOverdraft     SignedCurrencyAndAmount `json:"acceptedOverdraft"`
	Amount                SignedCurrencyAndAmount `json:"amount"`
	TotalClearedBalance   SignedCurrencyAndAmount `json:"totalClearedBalance"`
	TotalEffectiveBalance SignedCurrencyAndAmount `json:"totalEffectiveBalance"`
}
