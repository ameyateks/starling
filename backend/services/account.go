package services

import (
	"starling/starlingapi"
)

func GetStarlingAccountDetails() (starlingapi.StarlingBalanceAndSpacesResp, error){
	accountUid := starlingapi.GetStarlingAccountAndCategoryUid()

	balance, balanceErr := starlingapi.GetAccountBalance(accountUid.AccountUid)
	if balanceErr != nil { 
		return starlingapi.StarlingBalanceAndSpacesResp{}, balanceErr
	}

	spaces, spacesErr := starlingapi.GetSpaces(accountUid.AccountUid)
	if spacesErr != nil {
		return starlingapi.StarlingBalanceAndSpacesResp{}, spacesErr
	}

	return starlingapi.StarlingBalanceAndSpacesResp{Balance: balance.EffectiveBalance, Spaces: spaces}, nil
}