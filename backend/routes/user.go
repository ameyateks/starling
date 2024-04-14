package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"starling/services"
	"starling/types"
)

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
	req, err := http.NewRequest("GET", services.StarlingAPIBaseUrl+"account-holder/name", nil)
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
		var res types.StarlingUser
		json.Unmarshal(body, &res)
		fmt.Printf("%+v\n", res)

	}

}