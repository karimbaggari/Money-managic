package model

type UserIncomes struct  {
    Incomes       float64 `json:"incomes" bson:"incomes"`
    PassiveIncome float64 `json:"passiveIncome" bson:"passiveIncome"`
}

type UserExpenses struct {
    ConstantExpenses float64 `json:"constantExpenses" bson:"constantExpenses"`
    ExtraExpenses float64 `json:"extraExpenses" bson:"extraExpenses"`
}

type UserLivingBudget struct {
    MinimalLivingBudget float64 `json:"minimalLivingBudget" bson:"minimalLivingBudget"`
    MaximalLivingBudget float64 `json:"maximalLivingBudget" bson:"maximalLivingBudget"`
}

type UserSavings struct {
    MinimalSavings float64 `json:"minimalSavings" bson:"minimalSavings"`
    MaximalSavings float64 `json:"maximalSavings" bson:"maximalSavings"`
}


