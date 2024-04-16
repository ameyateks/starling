package dao

import (
	"fmt"
	"starling/dao/entities"

	"github.com/jmoiron/sqlx"
)

func FetchAllTransactions(db *sqlx.DB) ([]entities.Transaction, error) {
	query := `
	SELECT
		feed_item_uid
		, category_uid
		, amount
		, direction
		, transaction_time
		, counter_party_name
		, counter_party_sub_entity_name
		, reference
		, spending_category
		, user_note
	FROM transactions
	`

	var transactions []entities.Transaction
	err := db.Select(&transactions, query)

	if err != nil {
		return nil, fmt.Errorf("failed to fetch all transactions with error: %v", err)
	}

	return transactions, nil
}

func FetchTransactionsBetween(db *sqlx.DB, from string, to string) ([]entities.Transaction, error) {
	query := `
	SELECT
		feed_item_uid
		, category_uid
		, amount
		, direction
		, transaction_time
		, counter_party_name
		, counter_party_sub_entity_name
		, reference
		, spending_category
		, user_note
	FROM transactions
	WHERE transaction_time BETWEEN $1 AND $2;
	`

	var transactions []entities.Transaction
	err := db.Select(&transactions, query, from, to)

	if err != nil {
		return nil, fmt.Errorf("failed to fetch transactions with error: %v", err)
	}

	return transactions, nil
}

func InsertTransactions(db *sqlx.DB, transactions entities.Transaction) (error) {

	query := `
	INSERT INTO transactions
	(
		 feed_item_uid
        , category_uid
        , amount
        , direction
        , transaction_time
        , counter_party_name
        , counter_party_sub_entity_name
        , reference
        , spending_category
        , user_note
	) VALUES (
		:feed_item_uid
		, :category_uid
		, :amount
		, :direction
		, :transaction_time
		, :counter_party_name
		, :counter_party_sub_entity_name
		, :reference
		, :spending_category
		, :user_note
	);
	`

	_, err := db.NamedExec(
		query,
		transactions,
	)

	fmt.Print("err", err)

	if err != nil {
		return fmt.Errorf("failed to insert transactions with error: %v", err)
	}

	return nil
}