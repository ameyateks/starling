package starlingapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"starling/utils"
)

func UpdateCategoryForTransactions(feedItemUid string, category string, accountUid string, categoryUid string) error {
	accessToken, accessTokenErr := utils.SourceAccessToken()
	if(accessTokenErr != nil) {
		return accessTokenErr
	}

	postBody, marshalErr := json.Marshal(CategoryUpdateReq{SpendingCategory: category, PermanentSpendingCategoryUpdate: false, PreviousSpendingCategoryReferencesUpdate: false})
	if marshalErr != nil {
		fmt.Println("ERROR: ", marshalErr)
		return marshalErr
	}

	client := &http.Client{}

	req, err := http.NewRequest(
		"PUT",
		fmt.Sprintf(
			"https://api.starlingbank.com/api/v2/feed/account/%s/category/%s/%s/spending-category", accountUid, categoryUid, feedItemUid),
		bytes.NewBuffer(postBody))
	if err != nil {
		return &RequestError{StatusCode: http.StatusInternalServerError, Message: err.Error()}
	}
	req.Header.Add("Content-Type", "application/json;charset=UTF-8")

	req.Header.Set("Authorization", "Bearer "+accessToken)

	res, clientErr := client.Do(req)

	fmt.Print("Status",res.StatusCode)

	if clientErr != nil {
		fmt.Println("ERROR: ", clientErr)
		return &RequestError{StatusCode: res.StatusCode, Message: clientErr.Error()}
	} else if res.StatusCode > 299 && res.StatusCode < 200 {
		return &RequestError{StatusCode: res.StatusCode, Message: "Failed to update category"}
	} else {
		defer res.Body.Close()

		return nil
	}
}

func GetTransactionsForTimePeriod(accountUid string, categoryUid string, firstDate string, secondDate string) (Transactions, error) {
	if utils.IsDemo() {
		return TransactionsTestData, nil
	}

	accessToken, accessTokenErr := utils.SourceAccessToken()
	if(accessTokenErr != nil) {
		return Transactions{}, accessTokenErr
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.starlingbank.com/api/v2/feed/account/%s/category/%s/transactions-between", accountUid, categoryUid), nil)
	if err != nil {
		return Transactions{}, &RequestError{StatusCode: http.StatusInternalServerError, Message: err.Error()}
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)

	q := req.URL.Query()
	q.Add("minTransactionTimestamp", firstDate)

	q.Add("maxTransactionTimestamp", secondDate)
	req.URL.RawQuery = q.Encode()

	res, clientErr := client.Do(req)
	fmt.Print(req.URL)
	fmt.Print("response",res)
	if clientErr != nil {
		return Transactions{}, &RequestError{StatusCode: res.StatusCode, Message: clientErr.Error()}
	}  else if res.StatusCode != 200 {
		return Transactions{}, &RequestError{StatusCode: res.StatusCode, Message: "failed to get transactions"}
	} else {
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)

		if err != nil {
			return Transactions{}, &RequestError{StatusCode: res.StatusCode, Message: err.Error()}
		}
		var res Transactions
		json.Unmarshal(body, &res)

		return res, nil
	}

}