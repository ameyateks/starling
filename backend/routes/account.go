package routes

import (
	"encoding/json"
	"net/http"
	"starling/services"
	"starling/utils"
)

func starlingAccount(w http.ResponseWriter, r *http.Request) {

	balanceAndSpaces, err := services.GetStarlingAccountDetails()

	if err != nil { 
		utils.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	balanceResp, err := json.Marshal(balanceAndSpaces)

	if err != nil {
		utils.WriteError(w, err, http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write(balanceResp)
	}

}