package repositories

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

type MongodbRepository struct {
    Collection *mongo.Collection
}

func NewRepository(collection *mongo.Collection) *MongodbRepository {
    return &MongodbRepository{Collection: collection}
}


func (r *MongodbRepository) InsertDocument(ctx context.Context, doc interface{}) (interface{}, error) {
    result, err :=  r.Collection.InsertOne(ctx, doc)
    if err != nil {
        fmt.Println("error in the insertion ", err.Error())
        return "",err
    }
    return result, err
}


func (r *MongodbRepository) FindDocument(ctx context.Context, filter interface{}) (interface{}, error) {
    return r.Collection.FindOne(ctx, filter), nil
}
