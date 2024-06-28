package repositories

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

type LivingBudgetRepository struct {
	Collection *mongo.Collection
}

func NewLivingBudgetRepository(collection *mongo.Collection) *LivingBudgetRepository {
	return &LivingBudgetRepository{Collection: collection}
}

func (r *LivingBudgetRepository) InsertDocument(ctx context.Context, doc interface{}) (interface{}, error) {
	result, err := r.Collection.InsertOne(ctx, doc)
	if err != nil {
		fmt.Println("error in the insertion : ", err.Error())
		return "", err
	}

	return result, err
}