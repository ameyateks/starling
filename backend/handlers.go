package main

import (
	"encoding/json"

	"fmt"
	"io"
	"net/http"
	"os"
)

const starlingAPIBaseUrl = "https://api.starlingbank.com/api/v2/"

func starlingAccount(w http.ResponseWriter, r *http.Request) {

	accountUid := getStarlingAccountUid()

	balance := getAccountBalance(accountUid)

	balanceResp, err := json.Marshal(StarlingBalanceResp{Balance: balance})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error marshalling resp: ", err)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(balanceResp)
	}

}

func starlingUser(w http.ResponseWriter, r *http.Request) {
	accessToken, exists := os.LookupEnv("ACCESS_TOKEN")

	if !exists {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Println("ERROR: ACCESS_TOKEN not set")
	} else {
		fmt.Println("ACCESS_TOKEN: ", accessToken)
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

func getStarlingAccountUid() string {

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
		return ""
	} else {
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
		}
		var res StarlingAccountInfo
		json.Unmarshal(body, &res)
		fmt.Println("account uid" + res.Accounts[0].AccountUid)

		return res.Accounts[0].AccountUid
	}

}
