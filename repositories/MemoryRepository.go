package repositories

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
)

type InMemoryRepository struct {
    Data map[string]interface{}
}

func NewInMemoryRepository() *InMemoryRepository {
    return &InMemoryRepository{Data: make(map[string]interface{})}
}

func (r *InMemoryRepository) InsertDocument(ctx context.Context, doc interface{}) (*mongo.InsertOneResult, error) {
    id := fmt.Sprintf("doc%d", len(r.Data)+1) // Generate a simple ID
    r.Data[id] = doc
    // Since we are using in-memory storage, return a mock InsertOneResult
    return &mongo.InsertOneResult{InsertedID: id}, nil
}

func (r *InMemoryRepository) FindDocument(ctx context.Context, filter interface{}) (*mongo.SingleResult, error) {
    return nil,nil
}