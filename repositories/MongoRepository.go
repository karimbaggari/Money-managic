package repositories

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongodbRepository struct {
    Collection *mongo.Collection
}

func NewRepository(collection *mongo.Collection) *MongodbRepository {
    return &MongodbRepository{Collection: collection}
}


func (r *MongodbRepository) InsertDocument(ctx context.Context, doc interface{}) (*mongo.InsertOneResult, error) {
    return r.Collection.InsertOne(ctx, doc)
}


func (r *MongodbRepository) FindDocument(ctx context.Context, filter interface{}) (*mongo.SingleResult, error) {
    return r.Collection.FindOne(ctx, filter), nil
}
