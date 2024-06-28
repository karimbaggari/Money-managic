package helpers

import (
	"context"
	"fmt"
	"money-managic/database"
	"money-managic/model"
	"money-managic/repositories/MongoRepo"
	"os"
)

func CalculateLivingBudget(userId string) (model.UserLivingBudget, error) {
	database.Connect()
	dbName := os.Getenv("DATABASE_NAME")
	dbCollection1 := os.Getenv("DATABASE_COLLECTION1")
	dbCollection2 := os.Getenv("DATABASE_COLLECTION2")

	collection := database.GetCollection(dbName, dbCollection1)

	incomes, err := repositories.NewRepository(collection).FindDocument(context.Background(), userId)
	if err != nil {
		fmt.Println("issue into finding the data of the user check the user Id if it's valid ")
	}

	collection = database.GetCollection(dbName, dbCollection2)

	expenses, err := repositories.NewRepository(collection).FindDocument(context.Background(), userId)
	if err != nil {
		fmt.Println("issue into finding the data of the user check the user Id if it's valid ")
	}

	// Type assertion
	incomesValue, ok := incomes.(float64)
	if !ok {
		fmt.Println("incomes is not of type float64")
	}

	expensesValue, ok := expenses.(float64)
	if !ok {
		fmt.Println("expenses is not of type float64")
	}
	savingPercentage := 0.2
	result := incomesValue - expensesValue - (incomesValue * savingPercentage)
	var budget model.UserLivingBudget
	if result >= 3000 {
		budget.MinimalLivingBudget = result
		budget.MaximalLivingBudget = result
		return budget, nil
	} else {
		for result >= 3000 {
			savingPercentage = savingPercentage - 0.01
			result = incomesValue - expensesValue - (incomesValue * savingPercentage)
		}
		budget.MinimalLivingBudget = result
		return budget, nil
	}
}
