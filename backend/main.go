package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	accountUid := getStarlingAccountAndCategoryUid()

	now := time.Now()
	yearAgo := now.AddDate(-1, 0, 0)

	transacations := getTransactionsForTimePeriod(accountUid.AccountUid, accountUid.CategoryUid, yearAgo.Format(time.RFC3339), now.Format(time.RFC3339))

	allTransactions = transacations.FeedItems

	allTransactionsBytes, transactionsErr := json.Marshal(allTransactions)
	check(transactionsErr)

	err := os.WriteFile("/tmp/transactions.json", allTransactionsBytes, 0644)
	check(err)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	mux := mux.NewRouter()
	// mux.HandleFunc("/api/user", starlingUser).Methods("GET")
	mux.HandleFunc("/api/accounts", starlingAccount).Methods("GET")
	mux.HandleFunc("/api/transactions", starlingTransactions).Methods("GET")
	mux.HandleFunc("/api/knn", classifyTransaction).Methods("POST", "OPTIONS")

	http.ListenAndServe(":8080", mux)
}
