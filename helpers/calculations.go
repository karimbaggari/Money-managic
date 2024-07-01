package helpers

import (
	"context"
	"fmt"
	"money-managic/database"
	"money-managic/model"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CalculateLivingBudget(incomesId string, expensesId string) (model.UserLivingBudget, error) {
	database.Connect()
	dbName := os.Getenv("DATABASE_NAME")
	dbCollection1 := os.Getenv("DATABASE_COLLECTION1")
	dbCollection2 := os.Getenv("DATABASE_COLLECTION2")

	collection := database.GetCollection(dbName, dbCollection1)

	objectID, err := primitive.ObjectIDFromHex(incomesId)
	if err != nil {
		fmt.Println("Invalid ObjectID:", err)
	}

	// Query document by ObjectID
	var incomesModel model.UserIncomes
	err = collection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&incomesModel)
	if err != nil {
		fmt.Println("issue into finding the data of the user check the user Id if it's valid ")
	}

	collection = database.GetCollection(dbName, dbCollection2)

	objectID, err = primitive.ObjectIDFromHex(expensesId)
	if err != nil {
		fmt.Println("Invalid ObjectID:", err)
	}

	// Query document by ObjectID
	var expensesModel model.UserExpenses
	err = collection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&expensesModel)
	if err != nil {
		fmt.Println("issue into finding the data of the user check the user Id if it's valid ")
	}


	savingPercentage := 0.2
	result := (incomesModel.Incomes + incomesModel.PassiveIncome) - (expensesModel.ConstantExpenses + expensesModel.ExtraExpenses)
	fmt.Println("the rest is ", result)
	var budget model.UserLivingBudget
	if result >= 3000 {
		amountToSave := result * savingPercentage
		budget.MaximalLivingBudget = result - amountToSave
		fmt.Println("the value of savings is ", budget.MaximalLivingBudget)
		return budget, nil
	} else {
		for result >= 3000 {
			savingPercentage = savingPercentage - 0.01
			budget.MinimalLivingBudget = budget.MinimalLivingBudget - (expensesModel.ConstantExpenses + expensesModel.ExtraExpenses) - (budget.MinimalLivingBudget * savingPercentage)
		}
		fmt.Println("the value is ", budget.MinimalLivingBudget)
		return budget, nil
	}

}
