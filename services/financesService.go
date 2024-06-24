package services

import (
	"context"
	"money-managic/model"
	"money-managic/repositories"

	"go.mongodb.org/mongo-driver/mongo"
)

type Services struct {
	Repo *repositories.Repository
}

func NewServices(repo *repositories.Repository) *Services {
	return &Services{Repo: repo}
}

func  (s *Services) EnterFinances(finance model.UserFinance) (*mongo.InsertOneResult, error) {
	return s.Repo.InsertDocument(context.Background(), finance)
}