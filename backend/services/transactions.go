package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"starling/types"
	"starling/utils"
	"time"

	"github.com/joho/godotenv"
)

func UpdateCategoryForTransactions(categoryUpdate types.CategoryUpdatePostBody, accountUid string, categoryUid string) (types.CategoryUpdatePostBody, error) {
	accessToken, exists := os.LookupEnv("ACCESS_TOKEN")

	if !exists {
		fmt.Println("ERROR: ACCESS_TOKEN not set")
	}

	postBody, marshalErr := json.Marshal(types.CategoryUpdateReq{SpendingCategory: categoryUpdate.Category, PermanentSpendingCategoryUpdate: false, PreviousSpendingCategoryReferencesUpdate: false})
	if marshalErr != nil {
		fmt.Println("ERROR: ", marshalErr)
		return types.CategoryUpdatePostBody{}, marshalErr
	}

	client := &http.Client{}

	req, err := http.NewRequest(
		"PUT",
		fmt.Sprintf(
			"https://api.starlingbank.com/api/v2/feed/account/%s/category/%s/%s/spending-category", accountUid, categoryUid, categoryUpdate.FeedItemUid),
		bytes.NewBuffer(postBody))
	if err != nil {
		fmt.Println("ERROR: ", err)
	}
	req.Header.Add("Content-Type", "application/json;charset=UTF-8")

	req.Header.Set("Authorization", "Bearer "+accessToken)

	res, clientErr := client.Do(req)

	if clientErr != nil {
		fmt.Println("ERROR: ", clientErr)
		return types.CategoryUpdatePostBody{}, &types.RequestError{StatusCode: res.StatusCode, Message: clientErr.Error()}
	} else if res.StatusCode != 200 {
		return types.CategoryUpdatePostBody{}, &types.RequestError{StatusCode: res.StatusCode, Message: "Failed to update category"}
	} else {
		defer res.Body.Close()

		_, err := io.ReadAll(res.Body)

		if err != nil {
			fmt.Println("readerr", err)
		}

		return categoryUpdate, nil //is this best practice? returning thing we passed in
	}
}

func GetTransactionsForTimePeriod(accountUid string, categoryUid string, firstDate string, secondDate string) types.Transactions {
	if utils.IsDemo() {
		return types.TransactionsTestData
	}

	accessToken, exists := os.LookupEnv("ACCESS_TOKEN")

	if !exists {
		fmt.Println("ERROR: ACCESS_TOKEN not set")
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.starlingbank.com/api/v2/feed/account/%s/category/%s/transactions-between", accountUid, categoryUid), nil)
	if err != nil {
		fmt.Println("ERROR: ", err)
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)

	q := req.URL.Query()
	q.Add("minTransactionTimestamp", firstDate)

	q.Add("maxTransactionTimestamp", secondDate)
	req.URL.RawQuery = q.Encode()

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("ERROR: ", err)
		return types.Transactions{}
	} else {
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)

		if err != nil {
			fmt.Println(err)
		}
		var res types.Transactions
		json.Unmarshal(body, &res)

		return res
	}

}

func RunningKnnOnTransactions() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	accountUid := GetStarlingAccountAndCategoryUid()

	now := time.Now()
	yearAgo := now.AddDate(-1, 0, 0)

	transacations := GetTransactionsForTimePeriod(
		accountUid.AccountUid,
		accountUid.CategoryUid,
		yearAgo.Format(time.RFC3339),
		now.Format(time.RFC3339))

	allTransactions := transacations.FeedItems

	allTransactionsBytes, transactionsErr := json.Marshal(allTransactions)
	utils.Check(transactionsErr)

	err := os.WriteFile("/tmp/transactions.json", allTransactionsBytes, 0644)
	utils.Check(err)
}
