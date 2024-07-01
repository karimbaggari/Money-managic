package controllers

import (
	"encoding/json"
	"fmt"
	"money-managic/helpers"
	"money-managic/model"
	"money-managic/services"
	"net/http"
)

type LivingBudgetHandler struct {
	Service *services.LivingBudgetService
}

func NewLivingBudgetHandler(service *services.LivingBudgetService) *LivingBudgetHandler {
	return &LivingBudgetHandler{Service: service}
}

func (h *LivingBudgetHandler) LivingBudgetMongoHandler(w http.ResponseWriter, r *http.Request) {
	var finance model.UserLivingBudget
	err := json.NewDecoder(r.Body).Decode(&finance)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	result, err := h.Service.InsertLivingBudgetService(finance)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func (h *LivingBudgetHandler) LivingBudgetMemoryHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("1")

	CalculationResult,savings, err := helpers.CalculateLivingBudget("667ddd0cccca77f5aaf54ef8","667ddd3ce75ba1d28706047f")
	fmt.Println(savings)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	result, err := h.Service.InsertLivingBudgetIntoMemoryService(CalculationResult)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
