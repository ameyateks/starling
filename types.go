package main

type StarlingAccountInfo struct {
	Accounts []StarlingAccount `json:"accounts"`
}

type StarlingAccount struct {
	AccountUid      string `json:"accountUid"`
	AccountType     string `json:"accountType"`
	DefaultCategory string `json:"defaultCategory"`
	Currency        string `json:"currency"`
	CreatedAt       string `json:"createdAt"`
	Name            string `json:"name"`
}

type StarlingAccountResp struct {
	AccountUid string `json:"accountUid"`
}
