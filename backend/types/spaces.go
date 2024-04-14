package types

type StarlingSpaces struct {
	SavingGoals    []StarlingSavingSpace   `json:"savingsGoals"`
	SpendingSpaces []StarlingSpendingSpace `json:"spendingSpaces"`
}

type StarlingSavingSpace struct {
	SavingsGoalUid  string                  `json:"savingsGoalUid"`
	Name            string                  `json:"name"`
	Target          SignedCurrencyAndAmount `json:"target"`
	TotalSaved      SignedCurrencyAndAmount `json:"totalSaved"`
	SavedPercentage float32                 `json:"savedPercentage"`
	SortOrder       int                     `json:"sortOrder"`
	State           string                  `json:"state"`
}

type StarlingSpendingSpace struct {
	SpaceUid           string                  `json:"spaceUid"`
	Name               string                  `json:"name"`
	SortOrder          int                     `json:"sortOrder"`
	State              string                  `json:"state"`
	SpendingSpaceType  string                  `json:"spendingSpaceType"`
	CardAssociationUid string                  `json:"cardAssociationUid"`
	Balance            SignedCurrencyAndAmount `json:"balance"`
}

type StarlingBalanceAndSpacesResp struct {
	Balance SignedCurrencyAndAmount `json:"balance"`
	Spaces  StarlingSpaces          `json:"spaces"`
}