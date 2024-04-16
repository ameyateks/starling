package main

import (
	"fmt"
	"log"
	"starling/dao"
	"starling/dao/entities"
	"starling/starlingapi"
	"starling/utils"
	"time"

	"github.com/joho/godotenv"
)

// this function is a one time script to backfill transactions in the database, keeping here for clarity

func runTransactionsBackfill() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	now := time.Now()

	accountUid, categoryUid, createdAt := starlingapi.GetStarlingAccountAndCategoryUid().AccountUid, starlingapi.GetStarlingAccountAndCategoryUid().CategoryUid, starlingapi.GetStarlingAccountAndCategoryUid().CreatedAt

	transactions, getTransactionsErr := starlingapi.GetTransactionsForTimePeriod(accountUid, categoryUid, createdAt, now.Format(time.RFC3339))
	if getTransactionsErr != nil {
		log.Fatalf("Error getting transactions: %v", getTransactionsErr)
	}

	transactionsArr := transactions.FeedItems

	var transactionsDao []entities.Transaction

	for _, transaction := range transactionsArr {
		transactionDao := entities.StarlingTransactionToDao(transaction)
		transactionsDao = append(transactionsDao, transactionDao)
	}

	db := utils.ConnectToDB()

	tx := db.MustBegin()

	for _, transactionDao := range transactionsDao {
		err := dao.InsertTransactions(db, transactionDao)
		if err != nil {
			tx.Rollback()
			panic(err)
		}
	}

	commitErr := tx.Commit()
	fmt.Print("Backfill Completed!")
	if commitErr != nil {
		panic(commitErr)
	}

}
