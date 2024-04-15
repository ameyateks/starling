package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"starling/types"
	"starling/utils"
	"time"
)

func UpdateCategoryForTransactions(categoryUpdate types.CategoryUpdatePostBody, accountUid string, categoryUid string) (types.CategoryUpdatePostBody, error) {
	accessToken, accessTokenErr := utils.SourceAccessToken()
	if(accessTokenErr != nil) {
		return types.CategoryUpdatePostBody{}, accessTokenErr
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
		return types.CategoryUpdatePostBody{}, &types.RequestError{StatusCode: http.StatusInternalServerError, Message: err.Error()}
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

		return categoryUpdate, nil
	}
}

func GetTransactionsForTimePeriod(accountUid string, categoryUid string, firstDate string, secondDate string) (types.Transactions, error) {
	if utils.IsDemo() {
		return types.TransactionsTestData, nil
	}

	accessToken, accessTokenErr := utils.SourceAccessToken()
	if(accessTokenErr != nil) {
		return types.Transactions{}, accessTokenErr
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.starlingbank.com/api/v2/feed/account/%s/category/%s/transactions-between", accountUid, categoryUid), nil)
	if err != nil {
		return types.Transactions{}, &types.RequestError{StatusCode: http.StatusInternalServerError, Message: err.Error()}
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)

	q := req.URL.Query()
	q.Add("minTransactionTimestamp", firstDate)

	q.Add("maxTransactionTimestamp", secondDate)
	req.URL.RawQuery = q.Encode()

	res, clientErr := client.Do(req)
	if clientErr != nil {
		return types.Transactions{}, &types.RequestError{StatusCode: res.StatusCode, Message: clientErr.Error()}
	}  else if res.StatusCode != 200 {
		return types.Transactions{}, &types.RequestError{StatusCode: res.StatusCode, Message: "Failed to update category"}
	} else {
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)

		if err != nil {
			return types.Transactions{}, &types.RequestError{StatusCode: res.StatusCode, Message: err.Error()}
		}
		var res types.Transactions
		json.Unmarshal(body, &res)

		return res, nil
	}

}

func RunningKnnOnTransactions() {

	accountUid := GetStarlingAccountAndCategoryUid()

	now := time.Now()
	yearAgo := now.AddDate(-1, 0, 0)

	transacations, getErr := GetTransactionsForTimePeriod(
		accountUid.AccountUid,
		accountUid.CategoryUid,
		yearAgo.Format(time.RFC3339), 
		//TODO: rather than just use arbitary one year ago today date, call Starling's /api/v2/accounts
		// to get account information specifically when main account was created_at and use that as the
		// start date of all transactions, this ensures we retrieve all available transactions
		now.Format(time.RFC3339))

	if getErr != nil { fmt.Print("ERROR: ", getErr)}

	allTransactions := transacations.FeedItems

	allTransactionsBytes, transactionsErr := json.Marshal(allTransactions)
	utils.Check(transactionsErr)

	err := os.WriteFile("/tmp/transactions.json", allTransactionsBytes, 0644)
	utils.Check(err)
}
