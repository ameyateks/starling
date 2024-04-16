package routes

import (
	"encoding/json"
	"net/http"
	"starling/services"
	"starling/types"
	"starling/utils"
)

func starlingAccount(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	accountUid := services.GetStarlingAccountAndCategoryUid()

	balance, balanceErr := services.GetAccountBalance(accountUid.AccountUid)
	if balanceErr != nil { 
		utils.WriteError(w, balanceErr, http.StatusInternalServerError)
		return
	}

	spaces, spacesErr := services.GetSpaces(accountUid.AccountUid)
	if spacesErr != nil {
		utils.WriteError(w, spacesErr, http.StatusInternalServerError)
		return
	}

	balanceResp, err := json.Marshal(types.StarlingBalanceAndSpacesResp{Balance: balance.EffectiveBalance, Spaces: spaces})

	if err != nil {
		utils.WriteError(w, err, http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(balanceResp)
	}

}