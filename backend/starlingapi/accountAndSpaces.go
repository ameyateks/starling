package starlingapi

import (
	"encoding/json"
	"starling/utils"

	"fmt"
	"io"
	"net/http"
	"os"
)


const StarlingAPIBaseUrl = "https://api.starlingbank.com/api/v2/"

func GetSpaces(accountId string) (StarlingSpaces, error) {
	if utils.IsDemo() {
		return StarlingSpacesTestData, nil
	}

	accessToken, accessTokenErr := utils.SourceAccessToken()
	if(accessTokenErr != nil) {
		return StarlingSpaces{}, accessTokenErr
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", StarlingAPIBaseUrl+"account/"+accountId+"/spaces", nil)
	if err != nil {
		return StarlingSpaces{}, &RequestError{StatusCode: http.StatusInternalServerError, Message: err.Error()}
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)

	res, clientErr := client.Do(req)

	if clientErr != nil {
		fmt.Println("ERROR: ", err)
		return StarlingSpaces{}, &RequestError{StatusCode: res.StatusCode, Message: clientErr.Error()}
	} else {
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			return StarlingSpaces{}, &RequestError{StatusCode: res.StatusCode, Message: err.Error()}
			}
		var res StarlingSpaces
		json.Unmarshal(body, &res)

		return res, nil
	}

}

func GetAccountBalance(accountId string) (StarlingBalance, error) {
	if utils.IsDemo() {
		return StarlingBalanceTestData, nil
	}

	accessToken, accessTokenErr := utils.SourceAccessToken()
	if(accessTokenErr != nil) {
		return StarlingBalance{}, accessTokenErr
	}

	client := &http.Client{}
	req, newHttpErr := http.NewRequest("GET", StarlingAPIBaseUrl+"accounts/"+accountId+"/balance", nil)
	if newHttpErr != nil {
		return StarlingBalance{}, newHttpErr
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)

	res, clientErr := client.Do(req)
	if clientErr != nil {
		return StarlingBalance{}, &RequestError{StatusCode: res.StatusCode, Message: clientErr.Error()}
	} else if res.StatusCode != 200 {
		return StarlingBalance{}, &RequestError{StatusCode: res.StatusCode, Message: "failed to get Starling Balance"}
	} else {
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
		}
		var res StarlingBalance
		json.Unmarshal(body, &res)
		fmt.Printf("%+v\n", res)

		return res, nil
	}

}

func GetStarlingAccountAndCategoryUid() AccountAndCategoryUid {
	if utils.IsDemo() {
		return AccountAndCatUidTestData
	}

	accessToken, exists := os.LookupEnv("ACCESS_TOKEN")

	if !exists {
		fmt.Println("ERROR: ACCESS_TOKEN not set")
		return AccountAndCategoryUid{}
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
		return AccountAndCategoryUid{}
	} else {
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
		}
		var res StarlingAccountInfo
		json.Unmarshal(body, &res)

		return AccountAndCategoryUid{AccountUid: res.Accounts[0].AccountUid, CategoryUid: res.Accounts[0].DefaultCategory, CreatedAt: res.Accounts[0].CreatedAt}

	}

}

