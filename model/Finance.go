package model

type UserIncomes struct  {
    Incomes       float64 `json:"incomes" bson:"incomes"`
    PassiveIncome float64 `json:"passiveIncome" bson:"passiveIncome"`
}

type UserExpenses struct {
    constantExpenses float64 `json:"constantExpenses" bson:"constantExpenses"`
    extraExpenses float64 `json:"extraExpenses" bson:"extraExpenses"`
}

type UserLivingBudget struct {
    minimalLivingBudget float64 `json:"minimalLivingBudget" bson:"minimalLivingBudget"`
    maximalLivingBudget float64 `json:"maximalLivingBudget" bson:"maximalLivingBudget"`
}

type UserSavings struct {
    minimalSavings float64 `json:"minimalSavings" bson:"minimalSavings"`
    maximalSavings float64 `json:"maximalSavings" bson:"maximalSavings"`
}


