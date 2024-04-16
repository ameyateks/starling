package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"starling/dao/entities"
	"starling/services"
	"starling/starlingapi"
	"starling/utils"

	"github.com/jmoiron/sqlx"
)

func wrapStarlingTransactionsHandler(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		starlingTransactionsHandler(w, r, db)
	}
}

func starlingTransactionsHandler(w http.ResponseWriter, r *http.Request, db *sqlx.DB) {

	transactions, getTransactionsErr  := services.GetTransactions(db)
	if(getTransactionsErr != nil) {
		utils.WriteError(w, getTransactionsErr, http.StatusInternalServerError)
		return
	}

	transactionsResp, err := json.Marshal(transactions)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write(transactionsResp)
	}
}

func classifyTransaction(w http.ResponseWriter, r *http.Request) {

	var requestBody entities.Transaction
	
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		utils.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	classificationResp, err := services.ClassifyTransaction(requestBody)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error running python script: ", err)
	}
	w.Write(classificationResp)
}

func updateCategoryForTransactionsHandler(w http.ResponseWriter, r *http.Request) {

	var postBody starlingapi.CategoryUpdatePostBody
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&postBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := services.UpdateTransactionCategory(postBody.FeedItemUid, postBody.Category)

    resp, jsonErr := json.Marshal(postBody)
	if err != nil {
		utils.WriteError(w, err, http.StatusInternalServerError)
	} else if jsonErr != nil {
		utils.WriteError(w, jsonErr, http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write(resp)
	}
}

