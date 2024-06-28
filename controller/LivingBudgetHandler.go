package controllers

import (
	"encoding/json"
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

func (h *LivingBudgetHandler) LivingBudget(w http.ResponseWriter, r *http.Request) {
	var finance model.UserLivingBudget
	err := json.NewDecoder(r.Body).Decode(&finance)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	CalculationResult, err := helpers.CalculateLivingBudget("123")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	result, err := h.Service.InsertLivingBudgetService(CalculationResult)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
