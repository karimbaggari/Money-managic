package model

type UserIncomes struct  {
    UserID   string  `json:"user_id" bson:"user_id"`
    Incomes       float64 `json:"incomes" bson:"incomes"`
    PassiveIncome float64 `json:"passiveIncome" bson:"passiveIncome"`
}

type UserExpenses struct {
    constantExpenses float64
    extraExpenses float64
}

type UserLivingBudget struct {
    minimalLivingBudget float64
    maximalLivingBudget float64
}

type UserSavings struct {
    minimalSavings float64
    maximalSavings float64
}


