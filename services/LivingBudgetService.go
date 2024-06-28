package services

import (
	"context"
	"money-managic/model"
	"money-managic/repositories/MongoRepo"
	memory "money-managic/repositories/MemoryRepo"
)

type LivingBudgetService struct {
	Repo *repositories.LivingBudgetRepository
	MemoryRepo *memory.InMemoryRepository
}

func NewLivingBudgetServices(repo *repositories.LivingBudgetRepository, memory *memory.InMemoryRepository) *LivingBudgetService {
	return &LivingBudgetService{Repo: repo, MemoryRepo: memory}
}

func (s *LivingBudgetService) InsertLivingBudgetIntoMemoryService(livingBudget model.UserLivingBudget) (interface{}, error) {
	return s.MemoryRepo.InsertDocument(livingBudget)
}

func (s *LivingBudgetService) InsertLivingBudgetService(livingBudget model.UserLivingBudget) (interface{}, error) {
	return s.Repo.InsertDocument(context.Background(), livingBudget)
}


