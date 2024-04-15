package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"starling/dao"
	"starling/services"
	"starling/types"
	"starling/utils"
	"time"

	"github.com/jmoiron/sqlx"
)

func wrapStarlingTransactionsHandler(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		starlingTransactionsHandler(w, r, db)
	}
}

func starlingTransactionsHandler(w http.ResponseWriter, r *http.Request, db *sqlx.DB) {
	enableCors(&w)

	now := time.Now()

	thirtyDaysAgo := now.AddDate(0, 0, -30)

	transactions, getTransactionsErr  := dao.FetchTransactionsBetween(db, thirtyDaysAgo.Format(time.RFC3339), now.Format(time.RFC3339))
	if(getTransactionsErr != nil) {
		utils.WriteError(w, getTransactionsErr, http.StatusInternalServerError)
	}

	transactionsResp, err := json.Marshal(transactions)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(transactionsResp)
	}
}

func classifyTransaction(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	var requestBody types.Transaction
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Printf("Received Name: %s", requestBody.CounterPartyName)

	transactionFromReq, oneErr := json.Marshal(requestBody)

	if oneErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	pythonCommand := exec.Command("python3", "../data/knn.py", string(transactionFromReq))
	classificationResp, pythonErr := pythonCommand.CombinedOutput()
	if pythonErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error running python script: ", pythonErr)
	}

	w.Write(classificationResp)
}

func updateCategoryForTransactionsHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	var postBody types.CategoryUpdatePostBody
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&postBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	accountUid, categoryUid := services.GetStarlingAccountAndCategoryUid().AccountUid, services.GetStarlingAccountAndCategoryUid().CategoryUid

	updateCategory, putError := services.UpdateCategoryForTransactions(postBody, accountUid, categoryUid)

	resp, err := json.Marshal(types.CategoryUpdatePostBody{FeedItemUid: updateCategory.FeedItemUid, Category: updateCategory.Category})

	if err != nil {
		utils.WriteError(w, err, http.StatusInternalServerError)
	} else if putError != nil {
		utils.WriteError(w, putError, http.StatusInternalServerError)

	} else {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(resp)
	}
}
