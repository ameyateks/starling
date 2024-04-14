package types

type ClassifiedTransaction struct {
	Category string `json:"category"`
}

type SignedCurrencyAndAmount struct {
	Currency   string `json:"currency"`
	MinorUnits int    `json:"minorUnits"`
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

type Transactions struct {
	FeedItems []Transaction `json:"feedItems"`
}

type CategoryUpdatePostBody struct {
	FeedItemUid string `json:"feedItemUid"`
	Category    string `json:"category"` //TODO: map Starling's categories from docs to own type
}

type CategoryUpdateReq struct {
	SpendingCategory                         string `json:"spendingCategory"`
	PermanentSpendingCategoryUpdate          bool   `json:"permanentSpendingCategoryUpdate"`
	PreviousSpendingCategoryReferencesUpdate bool   `json:"previousSpendingCategoryReferencesUpdate"`
}