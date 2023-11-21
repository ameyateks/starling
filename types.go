package main

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

type StarlingBalance struct {
	ClearedBalance        SignedCurrencyAndAmount `json:"clearedBalance"`
	EffectiveBalance      SignedCurrencyAndAmount `json:"effectiveBalance"`
	PendingTransactions   SignedCurrencyAndAmount `json:"pendingTransactions"`
	AcceptedOverdraft     SignedCurrencyAndAmount `json:"acceptedOverdraft"`
	Amount                SignedCurrencyAndAmount `json:"amount"`
	TotalClearedBalance   SignedCurrencyAndAmount `json:"totalClearedBalance"`
	TotalEffectiveBalance SignedCurrencyAndAmount `json:"totalEffectiveBalance"`
}

type SignedCurrencyAndAmount struct {
	Currency   string `json:"currency"`
	MinorUnits int    `json:"minorUnits"`
}

type StarlingBalanceResp struct {
	Balance StarlingBalance `json:"balance"`
}

type StarlingAccountResp struct {
	AccountUid string `json:"accountUid"`
}
