package main

type ClassifiedTransaction struct {
	Category string `json:"category"`
}

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

type StarlingSpaces struct {
	SavingGoals    []StarlingSavingSpace   `json:"savingsGoals"`
	SpendingSpaces []StarlingSpendingSpace `json:"spendingSpaces"`
}

type StarlingSavingSpace struct {
	SavingsGoalUid  string                  `json:"savingsGoalUid"`
	Name            string                  `json:"name"`
	Target          SignedCurrencyAndAmount `json:"target"`
	TotalSaved      SignedCurrencyAndAmount `json:"totalSaved"`
	SavedPercentage float32                 `json:"savedPercentage"`
	SortOrder       int                     `json:"sortOrder"`
	State           string                  `json:"state"`
}

type StarlingSpendingSpace struct {
	SpaceUid           string                  `json:"spaceUid"`
	Name               string                  `json:"name"`
	SortOrder          int                     `json:"sortOrder"`
	State              string                  `json:"state"`
	SpendingSpaceType  string                  `json:"spendingSpaceType"`
	CardAssociationUid string                  `json:"cardAssociationUid"`
	Balance            SignedCurrencyAndAmount `json:"balance"`
}

type StarlingUser struct {
	AccountHolderName string `json:"accountHolderName"`
}

type SignedCurrencyAndAmount struct {
	Currency   string `json:"currency"`
	MinorUnits int    `json:"minorUnits"`
}

type StarlingBalanceAndSpacesResp struct {
	Balance SignedCurrencyAndAmount `json:"balance"`
	Spaces  StarlingSpaces          `json:"spaces"`
}

type AccountAndCategoryUid struct {
	AccountUid  string `json:"accountUid"`
	CategoryUid string `json:"categoryUid"`
}

type Transactions struct {
	FeedItems []Transaction `json:"feedItems"`
}

type TransactionResp struct {
	Transactions []Transaction `json:"transactions"`
}

type Transaction struct {
	FeedItemUid                        string                  `json:"feedItemUid"`
	CategoryUid                        string                  `json:"categoryUid"`
	Amount                             SignedCurrencyAndAmount `json:"amount"`
	SourceAmount                       SignedCurrencyAndAmount `json:"sourceAmount"`
	Direction                          string                  `json:"direction"`
	UpdatedAt                          string                  `json:"updatedAt"`
	TransactionTime                    string                  `json:"transactionTime"`
	SettlementTime                     string                  `json:"settlementTime"`
	RetryAllocationUntilTime           string                  `json:"retryAllocationUntilTime"`
	Source                             string                  `json:"source"`
	SourceSubType                      string                  `json:"sourceSubType"`
	Status                             string                  `json:"status"`
	TransactingApplicationUserUid      string                  `json:"transactingApplicationUserUid"`
	CounterPartyType                   string                  `json:"counterPartyType"`
	CounterPartyUid                    string                  `json:"counterPartyUid"`
	CounterPartyName                   string                  `json:"counterPartyName"`
	CounterPartySubEntityUid           string                  `json:"counterPartySubEntityUid"`
	CounterPartySubEntityName          string                  `json:"counterPartySubEntityName"`
	CounterPartySubEntityIdentifier    string                  `json:"counterPartySubEntityIdentifier"`
	CounterPartySubEntitySubIdentifier string                  `json:"counterPartySubEntitySubIdentifier"`
	ExchangeRate                       int                     `json:"exchangeRate"`
	TotalFees                          int                     `json:"totalFees"`
	TotalFeeAmount                     SignedCurrencyAndAmount `json:"totalFeeAmount"`
	Reference                          string                  `json:"reference"`
	Country                            string                  `json:"country"`
	SpendingCategory                   string                  `json:"spendingCategory"`
	UserNote                           string                  `json:"userNote"`
	RoundUp                            RoundUp                 `json:"roundUp"`
	HasAttachment                      bool                    `json:"hasAttachment"`
	HasReceipt                         bool                    `json:"hasReceipt"`
	BatchPaymentDetails                *BatchPaymentDetails    `json:"batchPaymentDetails"`
}

type RoundUp struct {
	GoalCategoryUid string                  `json:"goalCategoryUid"`
	Amount          SignedCurrencyAndAmount `json:"amount"`
}

type BatchPaymentDetails struct {
	BatchPaymentUid  string `json:"batchPaymentUid"`
	BatchPaymentType string `json:"batchPaymentType"`
}

type StarlingAccountResp struct {
	AccountUid string `json:"accountUid"`
}
