package services

import (
	"context"
	"money-managic/model"
	"money-managic/repositories/MongoRepo"

)


type ExpensesServices struct {
	Repo *repositories.ExpensesRepository
}

func NewExpensesServices(repo *repositories.ExpensesRepository) *ExpensesServices { 
	return &ExpensesServices{Repo: repo}
}

func (s *Services) EnterExpensesService(expenses model.UserExpenses) (interface{}, error) {
	return s.Repo.InsertDocument(context.Background(), expenses)
}