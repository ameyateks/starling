package main

import (
	"bytes"
	"encoding/json"

	"fmt"
	"io"
	"net/http"
	"os"
)

var allTransactions []Transaction

const starlingAPIBaseUrl = "https://api.starlingbank.com/api/v2/"

func getSpaces(accountId string) StarlingSpaces {
	if(isDemo()) {
		return starlingSpacesTestData
	}


	accessToken, exists := os.LookupEnv("ACCESS_TOKEN")

	if !exists {
		fmt.Println("ERROR: ACCESS_TOKEN not set")
	} else {
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", starlingAPIBaseUrl+"account/"+accountId+"/spaces", nil)
	if err != nil {
		fmt.Println("ERROR: ", err)
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)

	res, _ := client.Do(req)

	if err != nil {
		fmt.Println("ERROR: ", err)
		return StarlingSpaces{}
	} else {
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
		}
		var res StarlingSpaces
		json.Unmarshal(body, &res)
		fmt.Printf("%+v\n", res)

		return res
	}

}

func getAccountBalance(accountId string) StarlingBalance {
	if(isDemo()) {
		return starlingBalanceTestData
	}

	accessToken, exists := os.LookupEnv("ACCESS_TOKEN")

	if !exists {
		fmt.Println("ERROR: ACCESS_TOKEN not set")
	} else {
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", starlingAPIBaseUrl+"accounts/"+accountId+"/balance", nil)
	if err != nil {
		fmt.Println("ERROR: ", err)
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)

	res, _ := client.Do(req)

	if err != nil {
		fmt.Println("ERROR: ", err)
		return StarlingBalance{}
	} else {
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
		}
		var res StarlingBalance
		json.Unmarshal(body, &res)
		fmt.Printf("%+v\n", res)

		return res
	}

}

func getStarlingAccountAndCategoryUid() AccountAndCategoryUid {
	if isDemo() {
		return accountAndCatUidTestData
	}

	accessToken, exists := os.LookupEnv("ACCESS_TOKEN")

	if !exists {
		fmt.Println("ERROR: ACCESS_TOKEN not set")
	} 

	client := &http.Client{}
	req, err := http.NewRequest("GET", starlingAPIBaseUrl+"accounts", nil)
	if err != nil {
		fmt.Println("ERROR: ", err)
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)

	res, _ := client.Do(req)

	if err != nil {
		fmt.Println("ERROR: ", err)
		return AccountAndCategoryUid{}
	} else {
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
		}
		var res StarlingAccountInfo
		json.Unmarshal(body, &res)

		return AccountAndCategoryUid{AccountUid: res.Accounts[0].AccountUid, CategoryUid: res.Accounts[0].DefaultCategory}

	}

}

func updateCategoryForTransactions(categoryUpdate CategoryUpdatePostBody, accountUid string, categoryUid string) CategoryUpdatePostBody {
	accessToken, exists := os.LookupEnv("ACCESS_TOKEN")

	if !exists {
		fmt.Println("ERROR: ACCESS_TOKEN not set")
	} 

	postBody, marshalErr := json.Marshal(CategoryUpdateReq{SpendingCategory: categoryUpdate.Category, PermanentSpendingCategoryUpdate: false, PreviousSpendingCategoryReferencesUpdate: false})
	if marshalErr != nil {
		fmt.Println("ERROR: ", marshalErr)
		//TODO: return error status code
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
	req.Header.Add("Content-Type", "application/json;charset=UTF-8"); 
	//without this was getting 415 invalid content type but this error was not displaying in FE. Need a way to display this FE
	
	req.Header.Set("Authorization", "Bearer "+accessToken)

	res, respErr := client.Do(req) 

	fmt.Print(res)

	if respErr != nil {
		fmt.Println("ERROR: ", err)
		return CategoryUpdatePostBody{}
	} else {
		defer res.Body.Close()

		_, err := io.ReadAll(res.Body)

		if err != nil {
			fmt.Println(err)
		}
		
		return categoryUpdate //is this best practice? returning thing we passed in
	}


	
}

func getTransactionsForTimePeriod(accountUid string, categoryUid string, firstDate string, secondDate string) Transactions {
	if(isDemo()) {
		return transactionsTestData
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
		return Transactions{}
	} else {
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)

		if err != nil {
			fmt.Println(err)
		}
		var res Transactions
		json.Unmarshal(body, &res)

		return res
	}

}

// func updateSpendCategoryForTransaction(accountUid string, categoryUid string, feedItemUid string) Transactions {
// 	if(isDemo()) {
// 		return transactionsTestData
// 	}
	
// 	accessToken, exists := os.LookupEnv("ACCESS_TOKEN")

// 	if !exists {
// 		fmt.Println("ERROR: ACCESS_TOKEN not set")
// 	} else {
// 	}

// 	client := &http.Client{}
// 	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.starlingbank.com/api/v2/feed/account/%s/category/%s/transactions-between", accountUid, categoryUid), nil)
// 	if err != nil {
// 		fmt.Println("ERROR: ", err)
// 	}

// 	req.Header.Set("Authorization", "Bearer "+accessToken)

// 	q := req.URL.Query()
// 	q.Add("minTransactionTimestamp", firstDate)

// 	q.Add("maxTransactionTimestamp", secondDate)
// 	req.URL.RawQuery = q.Encode()

// 	res, err := client.Do(req)
// 	if err != nil {
// 		fmt.Println("ERROR: ", err)
// 		return Transactions{}
// 	} else {
// 		defer res.Body.Close()

// 		body, err := io.ReadAll(res.Body)

// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		var res Transactions
// 		json.Unmarshal(body, &res)

// 		return res
// 	}

// }