package helpers

import (
	"context"
	// "encoding/json"
	"fmt"
	"money-managic/database"
	"money-managic/model"
	// "money-managic/repositories/MongoRepo"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CalculateLivingBudget(incomesId string, expensesId string) (model.UserLivingBudget, error) {
	database.Connect()
	dbName := os.Getenv("DATABASE_NAME")
	dbCollection1 := os.Getenv("DATABASE_COLLECTION1")
	//dbCollection2 := os.Getenv("DATABASE_COLLECTION2")

	collection := database.GetCollection(dbName, dbCollection1)

	 // Construct ObjectID from string
	 id := "667ddd0cccca77f5aaf54ef8"
	 objectID, err := primitive.ObjectIDFromHex(id)
	 if err != nil {
		 fmt.Println("Invalid ObjectID:", err)
	 }
 
	// Query document by ObjectID
	filter := bson.M{"_id": objectID}
	var result model.UserIncomes
	err = collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		fmt.Println("issue into finding the data of the user check the user Id if it's valid ")
	}
	fmt.Println(result)

	// collection = database.GetCollection(dbName, dbCollection2)

	// expenses, err := repositories.NewRepository(collection).FindDocument(context.Background(), expensesId)
	// if err != nil {
	// 	fmt.Println("issue into finding the data of the user check the user Id if it's valid ")
	// }

	// expensesValue, ok := expenses.(float64)
	// if !ok {
	// 	fmt.Println("expenses is not of type float64")
	// }
	// savingPercentage := 0.2
	// result := incomes. - expensesValue - (incomesValue.MinimalLivingBudget * savingPercentage)
	// var budget model.UserLivingBudget
	// if result >= 3000 {
	// 	budget.MinimalLivingBudget = result
	// 	budget.MaximalLivingBudget = result
	// 	return budget, nil
	// } else {
	// 	for result >= 3000 {
	// 		savingPercentage = savingPercentage - 0.01
	// 		result = incomesValue.MinimalLivingBudget - expensesValue - (incomesValue.MinimalLivingBudget * savingPercentage)
	// 	}
	// 	budget.MinimalLivingBudget = result
	// 	return budget, nil
	// }
	return model.UserLivingBudget{}, nil
}
