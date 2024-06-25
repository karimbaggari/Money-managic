package repositories

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type DocumentRepository interface {
    InsertDocument(ctx context.Context, doc interface{}) (*mongo.InsertOneResult, error)
    FindDocument(ctx context.Context, filter interface{}) (*mongo.SingleResult, error)
}

type Repository struct {
    Collection *mongo.Collection
}

func NewRepository(collection *mongo.Collection) *Repository {
    return &Repository{Collection: collection}
}


func (r *Repository) InsertDocument(ctx context.Context, doc interface{}) (*mongo.InsertOneResult, error) {
    return r.Collection.InsertOne(ctx, doc)
}

func (r *Repository) FindDocument(ctx context.Context, filter interface{}) (*mongo.SingleResult, error) {
    return r.Collection.FindOne(ctx, filter), nil
}
