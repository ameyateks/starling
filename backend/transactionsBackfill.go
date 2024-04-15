package main

import (
	"fmt"
	"log"
	"starling/dao"
	"starling/dao/entities"
	"starling/services"
	"starling/utils"
	"time"

	"github.com/joho/godotenv"
)

func runTransactionsBackfill() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	now := time.Now()

	accountUid, categoryUid, createdAt := services.GetStarlingAccountAndCategoryUid().AccountUid, services.GetStarlingAccountAndCategoryUid().CategoryUid, services.GetStarlingAccountAndCategoryUid().CreatedAt

	transactions, getTransactionsErr := services.GetTransactionsForTimePeriod(accountUid, categoryUid, createdAt, now.Format(time.RFC3339))
	if getTransactionsErr != nil {
		log.Fatalf("Error getting transactions: %v", getTransactionsErr)
	}

	transactionsArr := transactions.FeedItems

	var transactionsDao []entities.Transaction

	for _, transaction := range transactionsArr {
		transactionDao := entities.TransactionDomainToDao(transaction)
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
