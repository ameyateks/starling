package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"starling/types"
	"strings"
)

func IsDemo() bool {
	envValue, exists := os.LookupEnv("IS_DEMO")

	var isDemo bool

	if exists {
		isDemo = strings.ToLower(envValue) == "true" //TODO: use lib helper function to parse boolean better!
	} else {
		isDemo = false
	}

	return isDemo
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func WriteError(w http.ResponseWriter, reqError error, statusCode int)  {
	errResp, err := json.Marshal(types.ErrorResponse{Error: reqError.Error(), StatusCode: statusCode})
	if (err != nil) {
	   os.Exit(1)
   }
   	w.WriteHeader(statusCode)
	w.Write(errResp)
	
}

func SourceAccessToken() (string, error) {
	accessToken, exists := os.LookupEnv("ACCESS_TOKEN")

	if !exists {
		fmt.Println("ERROR: ACCESS_TOKEN not set")
		return "", errors.New("cannot find ACCESS_TOKEN in .env")
	} else {
		return accessToken, nil
	}
}