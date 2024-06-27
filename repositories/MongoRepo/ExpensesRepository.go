package repositories

import (
"go.mongodb.org/mongo-driver/mongo"
"fmt"
"context"
)

type ExpensesRepository struct {
	Collection *mongo.Collection	
}

func NewExpensesRepository(collection *mongo.Collection) *ExpensesRepository {
	return &ExpensesRepository{Collection: collection}
}

func (r *ExpensesRepository) InsertDocument(ctx context.Context, doc interface{}) (interface{}, error) {
	result, err := r.Collection.InsertOne(ctx, doc)
	if err != nil {
		fmt.Println("error in the insertion : ", err.Error())
		return "", err
	}

	return result, err 
}