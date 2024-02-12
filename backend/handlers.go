package main

import (
	"encoding/json"
	"os/exec"
	"time"

	"fmt"
	"io"
	"net/http"
	"os"
)

const starlingAPIBaseUrl = "https://api.starlingbank.com/api/v2/"

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
}

func starlingAccount(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	accountUid := getStarlingAccountAndCategoryUid()

	balance := getAccountBalance(accountUid.AccountUid)

	spaces := getSpaces(accountUid.AccountUid)

	balanceResp, err := json.Marshal(StarlingBalanceAndSpacesResp{Balance: balance.EffectiveBalance, Spaces: spaces})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error marshalling resp: ", err)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(balanceResp)
	}

}

func classifyTransaction(w http.ResponseWriter, r *http.Request) {
	testTransaction, oneErr := json.Marshal(classifiedTransactionTestData)
	transactions, transactionsErr := json.Marshal(transactionsTestData)
	if oneErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	if transactionsErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	pythonCommand := exec.Command("python3", "../data/knn.py", string(testTransaction), string(transactions))
	classificationResp, pythonErr := pythonCommand.CombinedOutput()
	if pythonErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error running python script: ", pythonErr)
	}

	w.Write(classificationResp)

}

func starlingTransactions(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	transactions := getTransactionsForMonth(getStarlingAccountAndCategoryUid().AccountUid, getStarlingAccountAndCategoryUid().CategoryUid)

	transactionsResp, err := json.Marshal(TransactionResp{Transactions: transactions.FeedItems})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error marshalling resp: ", err)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(transactionsResp)
	}
}

func starlingUser(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	accessToken, exists := os.LookupEnv("ACCESS_TOKEN")

	if !exists {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Println("ERROR: ACCESS_TOKEN not set")
	} else {
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", starlingAPIBaseUrl+"account-holder/name", nil)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("ERROR: ", err)
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)

	res, _ := client.Do(req)

	if err != nil {
		fmt.Println("ERROR: ", err)
	} else {
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
		}
		var res StarlingUser
		json.Unmarshal(body, &res)
		fmt.Printf("%+v\n", res)

	}

}

func getSpaces(accountId string) StarlingSpaces {
	accessToken, exists := os.LookupEnv("ACCESS_TOKEN")

	if !exists {
		fmt.Println("ERROR: ACCESS_TOKEN not set")
	} else {
		fmt.Println("ACCESS_TOKEN: ", accessToken)
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
	accessToken, exists := os.LookupEnv("ACCESS_TOKEN")

	if !exists {
		fmt.Println("ERROR: ACCESS_TOKEN not set")
	} else {
		fmt.Println("ACCESS_TOKEN: ", accessToken)
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

	accessToken, exists := os.LookupEnv("ACCESS_TOKEN")

	if !exists {
		fmt.Println("ERROR: ACCESS_TOKEN not set")
	} else {
		fmt.Println("ACCESS_TOKEN: ", accessToken)
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
		fmt.Println(res)
		fmt.Println("default cat uid" + res.Accounts[0].AccountUid)

		return AccountAndCategoryUid{AccountUid: res.Accounts[0].AccountUid, CategoryUid: res.Accounts[0].DefaultCategory}

	}

}

func getTransactionsForMonth(accountUid string, categoryUid string) Transactions {
	accessToken, exists := os.LookupEnv("ACCESS_TOKEN")

	if !exists {
		fmt.Println("ERROR: ACCESS_TOKEN not set")
	} else {
		fmt.Println("ACCESS_TOKEN Transactions: ", accessToken)
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.starlingbank.com/api/v2/feed/account/%s/category/%s/transactions-between", accountUid, categoryUid), nil)
	if err != nil {
		fmt.Println("ERROR: ", err)
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)

	now := time.Now()
	currentYear, currentMonth, _ := now.Date()
	currentLocation := now.Location()

	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)

	q := req.URL.Query()
	q.Add("minTransactionTimestamp", firstOfMonth.Format(time.RFC3339))

	q.Add("maxTransactionTimestamp", now.Format(time.RFC3339))
	req.URL.RawQuery = q.Encode()
	fmt.Println("query", req.URL.RawQuery)
	fmt.Println("full url", req.URL.String())

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
		fmt.Println(string(body))

		return res
	}

}
