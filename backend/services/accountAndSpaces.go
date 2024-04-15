package services

import (
	"encoding/json"
	"starling/types"
	"starling/utils"

	"fmt"
	"io"
	"net/http"
	"os"
)


const StarlingAPIBaseUrl = "https://api.starlingbank.com/api/v2/"

func GetSpaces(accountId string) (types.StarlingSpaces, error) {
	if utils.IsDemo() {
		return types.StarlingSpacesTestData, nil
	}

	accessToken, accessTokenErr := utils.SourceAccessToken()
	if(accessTokenErr != nil) {
		return types.StarlingSpaces{}, accessTokenErr
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", StarlingAPIBaseUrl+"account/"+accountId+"/spaces", nil)
	if err != nil {
		return types.StarlingSpaces{}, &types.RequestError{StatusCode: http.StatusInternalServerError, Message: err.Error()}
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)

	res, clientErr := client.Do(req)

	if clientErr != nil {
		fmt.Println("ERROR: ", err)
		return types.StarlingSpaces{}, &types.RequestError{StatusCode: res.StatusCode, Message: clientErr.Error()}
	} else {
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			return types.StarlingSpaces{}, &types.RequestError{StatusCode: res.StatusCode, Message: err.Error()}
			}
		var res types.StarlingSpaces
		json.Unmarshal(body, &res)

		return res, nil
	}

}

func GetAccountBalance(accountId string) types.StarlingBalance {
	if utils.IsDemo() {
		return types.StarlingBalanceTestData
	}

	accessToken, exists := os.LookupEnv("ACCESS_TOKEN")

	if !exists {
		fmt.Println("ERROR: ACCESS_TOKEN not set")
	} else {
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", StarlingAPIBaseUrl+"accounts/"+accountId+"/balance", nil)
	if err != nil {
		fmt.Println("ERROR: ", err)
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)

	res, _ := client.Do(req)

	if err != nil {
		fmt.Println("ERROR: ", err)
		return types.StarlingBalance{}
	} else {
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
		}
		var res types.StarlingBalance
		json.Unmarshal(body, &res)
		fmt.Printf("%+v\n", res)

		return res
	}

}

func GetStarlingAccountAndCategoryUid() types.AccountAndCategoryUid {
	if utils.IsDemo() {
		return types.AccountAndCatUidTestData
	}

	accessToken, exists := os.LookupEnv("ACCESS_TOKEN")

	if !exists {
		fmt.Println("ERROR: ACCESS_TOKEN not set")
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", StarlingAPIBaseUrl+"accounts", nil)
	if err != nil {
		fmt.Println("ERROR: ", err)
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)

	res, _ := client.Do(req)

	if err != nil {
		fmt.Println("ERROR: ", err)
		return types.AccountAndCategoryUid{}
	} else {
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
		}
		var res types.StarlingAccountInfo
		json.Unmarshal(body, &res)

		return types.AccountAndCategoryUid{AccountUid: res.Accounts[0].AccountUid, CategoryUid: res.Accounts[0].DefaultCategory, CreatedAt: res.Accounts[0].CreatedAt}

	}

}

