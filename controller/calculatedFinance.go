package controllers

import (
	"money-managic/services"
	"net/http"
	"money-managic/model"
	"encoding/json"
)

type CalculationHandler struct {
	Service *services.Services
}

func NewCalculationHandler(service *services.Services) *CalculationHandler {
    return &CalculationHandler{Service: service}
}

func (h *CalculationHandler) EnterIncomes(w http.ResponseWriter, r *http.Request) {
	var finance model.UserIncomes

	err := json.NewDecoder(r.Body).Decode(&finance)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := h.Service.EnterIncomesService(finance)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)

}