package repositories

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongodbRepository struct {
    Collection *mongo.Collection
}

func NewRepository(collection *mongo.Collection) *MongodbRepository {
    return &MongodbRepository{Collection: collection}
}


// func (r *MongodbRepository) FindDocument(ctx context.Context, filter interface{}) (interface{}, error) {
//     return r.Collection.FindOne(ctx, filter), nil
// }


func (r *MongodbRepository) FindDocument(ctx context.Context, id string) (interface{}, error) {
    var result bson.M  // Assuming you are fetching a BSON document

    // Construct filter to find document by ID
    filter := bson.M{"_id": id}

    // Perform query to find document
    err := r.Collection.FindOne(ctx, filter).Decode(&result)
    if err != nil {
        return nil, err
    }

    return result, nil
}
