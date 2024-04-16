package services

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"starling/dao"
	"starling/dao/entities"
	"starling/starlingapi"
	"starling/utils"
	"time"

	"github.com/jmoiron/sqlx"
)

func GetTransactions(db *sqlx.DB) ([]entities.Transaction, error) {
	if utils.IsDemo() {

		var transactions []entities.Transaction

		for _, starlingTransaction := range starlingapi.TransacationsTestDataArr {
			transaction := entities.StarlingTransactionToDao(starlingTransaction)
			transactions = append(transactions, transaction)

		}

		return transactions, nil

	}
	now := time.Now()

	thirtyDaysAgo := now.AddDate(0, 0, -30)

	return dao.FetchTransactionsBetween(db, thirtyDaysAgo.Format(time.RFC3339), now.Format(time.RFC3339))
}

func ClassifyTransaction(transaction entities.Transaction) ([]byte, error) {
	transactionFromReq, err := json.Marshal(transaction)
	if err != nil {
		return nil, err
	}
	pythonCommand := exec.Command("python3", "./data/knn.py", string(transactionFromReq))
	return pythonCommand.CombinedOutput()

}

func UpdateTransactionCategory(feedItemUid string, category string) error {
	accountUid, categoryUid := starlingapi.GetStarlingAccountAndCategoryUid().AccountUid, starlingapi.GetStarlingAccountAndCategoryUid().CategoryUid

	return starlingapi.UpdateCategoryForTransactions(feedItemUid, category, accountUid, categoryUid)
}

func RunningKnnOnTransactions(db *sqlx.DB) error {

	var transactions []entities.Transaction
	var getTransactionsErr error

	if utils.IsDemo() {
		for _, starlingTransaction := range starlingapi.TransacationsTestDataArr {
			transaction := entities.StarlingTransactionToDao(starlingTransaction)
			transactions = append(transactions, transaction)
		}
	} else {
		transactions, getTransactionsErr = dao.FetchAllTransactions(db)
		if getTransactionsErr != nil {
			return fmt.Errorf("failed to fetch all transactions with error: %v", getTransactionsErr)
		}
	}
	
	if getTransactionsErr != nil {
		return fmt.Errorf("failed to fetch all transactions with error: %v", getTransactionsErr)
	}

	transactionsResp, marshallErr := json.Marshal(transactions)
	if marshallErr != nil {
		return fmt.Errorf("failed to marshal transactions with error: %v", marshallErr)
	}

	err := os.WriteFile("/tmp/transactions.json", transactionsResp, 0644)
	utils.Check(err)
	return nil
}
