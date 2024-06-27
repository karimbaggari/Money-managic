package services

import (
	"context"
	"money-managic/model"
	"money-managic/repositories/interfaces"
	"money-managic/helpers"
)

type Services struct {
	Repo interfaces.DocumentRepository
}

func NewServices(repo interfaces.DocumentRepository) *Services {
	return &Services{Repo: repo}
}

func  (s *Services) EnterIncomesService(finance model.UserIncomes) (interface{}, error) {
	return s.Repo.InsertDocument(context.Background(), finance)
}

func (s *Services) EnterExpensesService(expenses model.UserExpenses) (interface{}, error) {
	return s.Repo.InsertDocument(context.Background(), expenses)
}

func (s *Services) InsertLivingBudgetService(livingBudget model.UserLivingBudget) (interface{}, error) {
	userId := "abc"
	livingBudgetCalculation := helpers.CalculateLivingBudget(livingBudget, userId)
	return s.Repo.InsertDocument(context.Background(), livingBudgetCalculation)
}

func (s *Services) InsertSavingsService(savings model.UserSavings) (interface{}, error) {
	return s.Repo.InsertDocument(context.Background(), savings)
}