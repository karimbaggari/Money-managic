package helpers

import (
	"context"
	"fmt"
	"money-managic/database"
	"money-managic/model"
	"money-managic/repositories"
	"os"
)

func CalculateLivingBudget(livingBudget model.UserLivingBudget, userId string) float64 {
	database.Connect()
	dbName := os.Getenv("DATABASE_NAME")
	dbCollection1 := os.Getenv("DATABASE_COLLECTION1")
	dbCollection2 := os.Getenv("DATABASE_COLLECTION2")

	collection := database.GetCollection(dbName, dbCollection1)

	incomes, err := repositories.NewRepository(collection).FindDocument(context.Background(), userId)
	if err != nil {
		fmt.Println("issue into finding the data of the user check the user Id if it's valid ")
		return 0.0
	}

	collection = database.GetCollection(dbName, dbCollection2)

	expenses, err := repositories.NewRepository(collection).FindDocument(context.Background(), userId)
	if err != nil {
		fmt.Println("issue into finding the data of the user check the user Id if it's valid ")
		return 0.0
	}

	// Type assertion
	incomesValue, ok := incomes.(float64)
	if !ok {
		fmt.Println("incomes is not of type float64")
		return 0.0
	}

	expensesValue, ok := expenses.(float64)
	if !ok {
		fmt.Println("expenses is not of type float64")
		return 0.0
	}
	savingPercentage := 0.2
	result := incomesValue - expensesValue - (incomesValue * savingPercentage)
	if result >= 3000 {
		return result
	} else {
		for result >= 3000 {
			savingPercentage = savingPercentage - 0.01
			result = incomesValue - expensesValue - (incomesValue * savingPercentage)
		}
		return result
	}
}
