package repositories

import (
"go.mongodb.org/mongo-driver/mongo"
"fmt"
"context"
)

type IncomesRepository struct {
	Collection *mongo.Collection	
}

func NewIncomesRepository(collection *mongo.Collection) *IncomesRepository {
	return &IncomesRepository{Collection: collection}
}

func (r *IncomesRepository) InsertDocument(ctx context.Context, doc interface{}) (interface{}, error) {
	result, err := r.Collection.InsertOne(ctx, doc)
	if err != nil {
		fmt.Println("error in the insertion : ", err.Error())
		return "", err
	}

	return result, err 
}