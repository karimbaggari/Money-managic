package interfaces

import (
	"context"
	
)

type DocumentRepository interface {
    InsertDocument(ctx context.Context, doc interface{}) (interface{}, error)
    FindDocument(ctx context.Context, filter interface{}) (interface{}, error)
}