package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func isDemo() bool {
	envValue, exists := os.LookupEnv("IS_DEMO")

	var isDemo bool

	if exists {
		isDemo = strings.ToLower(envValue) == "true" //TODO: use lib helper function to parse boolean better!
	} else {
		isDemo = false
	}

	return isDemo
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	accountUid := getStarlingAccountAndCategoryUid()

	now := time.Now()
	yearAgo := now.AddDate(-1, 0, 0)

	transacations := getTransactionsForTimePeriod(
		accountUid.AccountUid,
		accountUid.CategoryUid,
		yearAgo.Format(time.RFC3339),
		now.Format(time.RFC3339))

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
	mux.HandleFunc("/api/category", updateCategoryForTransactionsHandler).Methods("POST", "OPTIONS")

	http.ListenAndServe(":8080", mux)
}
