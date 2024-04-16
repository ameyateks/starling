package entities

import "starling/starlingapi"


type Transaction struct {
	FeedItemUid               string `db:"feed_item_uid" json:"feedItemUid"`
	CategoryUid               string `db:"category_uid" json:"categoryUid"`
	Amount                    int    `db:"amount" json:"amount"`
	Direction                 string `db:"direction" json:"direction"`
	TransactionTime           string `db:"transaction_time" json:"transactionTime"`
	CounterPartyName          string `db:"counter_party_name" json:"counterPartyName"`
	CounterPartySubEntityName string `db:"counter_party_sub_entity_name" json:"counterPartySubEntityName"`
	Reference                 string `db:"reference" json:"reference"`
	SpendingCategory          string `db:"spending_category" json:"spendingCategory"`
	UserNote                  string `db:"user_note" json:"userNote"`
}

func StarlingTransactionToDao(transaction starlingapi.Transaction) Transaction {
	return Transaction{
		FeedItemUid:               transaction.FeedItemUid,
		CategoryUid:               transaction.CategoryUid,
		Amount:                    transaction.Amount.MinorUnits,
		Direction:                 transaction.Direction,
		TransactionTime:           transaction.TransactionTime,
		CounterPartyName:          transaction.CounterPartyName,
		CounterPartySubEntityName: transaction.CounterPartySubEntityName,
		Reference:                 transaction.Reference,
		SpendingCategory:          transaction.SpendingCategory,
		UserNote:                  transaction.UserNote,
	}
}
