package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"starling/services"
	"starling/types"
)

func starlingAccount(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	accountUid := services.GetStarlingAccountAndCategoryUid()

	balance := services.GetAccountBalance(accountUid.AccountUid)

	spaces := services.GetSpaces(accountUid.AccountUid)

	balanceResp, err := json.Marshal(types.StarlingBalanceAndSpacesResp{Balance: balance.EffectiveBalance, Spaces: spaces})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error marshalling resp: ", err)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(balanceResp)
	}

}