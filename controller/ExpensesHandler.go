package controllers

import (
	"encoding/json"
	"money-managic/model"
	"money-managic/services"
	"net/http"
)

type ExpensesHandlers struct {
	ExpensesService *services.ExpensesServices
}

func NewExpensesHandler(expensesService *services.ExpensesServices) *ExpensesHandlers {
	return &ExpensesHandlers{ExpensesService : expensesService}
}

func (h *ExpensesHandlers) EnterExpenses(w http.ResponseWriter, r *http.Request) {
	var expenses model.UserExpenses

	err := json.NewDecoder(r.Body).Decode(&expenses)
	if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
	}

	result, err := h.ExpensesService.EnterExpensesService(expenses)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
			return
	}	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}