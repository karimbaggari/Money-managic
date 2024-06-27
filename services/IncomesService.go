package services

import (
	"money-managic/repositories/MongoRepo"
	"money-managic/model"
	"context"
)

type IncomesService struct { 
	Repo  *repositories.IncomesRepository
}

func NewIncomesService(repo *repositories.IncomesRepository) *IncomesService {
	return &IncomesService{Repo: repo}
}

func  (s *Services) EnterIncomesService(finance model.UserIncomes) (interface{}, error) {
	return s.Repo.InsertDocument(context.Background(), finance)
}