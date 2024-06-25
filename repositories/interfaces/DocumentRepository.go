package interfaces

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type DocumentRepository interface {
    InsertDocument(ctx context.Context, doc interface{}) (*mongo.InsertOneResult, error)
    FindDocument(ctx context.Context, filter interface{}) (*mongo.SingleResult, error)
}