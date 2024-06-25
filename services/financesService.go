package services

import (
	"context"
	"money-managic/model"
	"money-managic/repositories"

	"go.mongodb.org/mongo-driver/mongo"
)

type Services struct {
	Repo repositories.DocumentRepository
}

func NewServices(repo repositories.DocumentRepository) *Services {
	return &Services{Repo: repo}
}

func  (s *Services) EnterIncomesService(finance model.UserIncomes) (*mongo.InsertOneResult, error) {
	return s.Repo.InsertDocument(context.Background(), finance)
}

func (s *Services) EnterExpensesService(expenses model.UserExpenses) (*mongo.InsertOneResult, error) {
	return s.Repo.InsertDocument(context.Background(), expenses)
}

func (s *Services) InsertLivingBudgetService(livingBudget model.UserLivingBudget) (*mongo.InsertOneResult, error) {
	return s.Repo.InsertDocument(context.Background(), livingBudget)
}

func (s *Services) InsertSavingsService(savings model.UserSavings) (*mongo.InsertOneResult, error) {
	return s.Repo.InsertDocument(context.Background(), savings)
}