package services

import (
	"context"
	"money-managic/model"
	"money-managic/repositories/MongoRepo"
)

type LivingBudgetService struct {
	Repo *repositories.LivingBudgetRepository
}

func NewLivingBudgetServices(repo *repositories.LivingBudgetRepository) *LivingBudgetService {
	return &LivingBudgetService{Repo: repo}
}

func (s *LivingBudgetService) InsertLivingBudgetService(livingBudget model.UserLivingBudget) (interface{}, error) {
	return s.Repo.InsertDocument(context.Background(), livingBudget)
}
