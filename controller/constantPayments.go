package controllers

import (
	"encoding/json"
	"fmt"
	"money-managic/model"
	"money-managic/services"
	"net/http"
)

type Handlers struct {
	Service *services.Services
}

func NewHandlers(service *services.Services) *Handlers {
	return &Handlers{Service: service}
}

func (h *Handlers) EnterIncomes(w http.ResponseWriter, r *http.Request) {
	var finance model.UserIncomes

	err := json.NewDecoder(r.Body).Decode(&finance)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Printf("Received finance object: %+v\n", finance)

	result, err := h.Service.EnterIncomesService(finance)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func (h *Handlers) EnterExpenses(w http.ResponseWriter, r *http.Request) {
	var expenses model.UserExpenses

	err := json.NewDecoder(r.Body).Decode(&expenses)
	if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
	}

	result, err := h.Service.EnterExpensesService(expenses)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
			return
	}	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}


func (h *Handlers) InsertLivingBudget(w http.ResponseWriter, r *http.Request) {
	var LivingBudget model.UserLivingBudget

	err := json.NewDecoder(r.Body).Decode(&LivingBudget)
	if err!= nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
    }

	result, err := h.Service.InsertLivingBudgetService(LivingBudget)
	if err!= nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func (h *Handlers) InsertSavings(w http.ResponseWriter, r *http.Request) {
	var savings model.UserSavings 
	err := json.NewDecoder(r.Body).Decode(savings)
	if err!= nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

	result, err := h.Service.InsertSavingsService(savings)
	if err!= nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
}

w.Header().Set("Content-Type", "application/json")
w.WriteHeader(http.StatusOK)
json.NewEncoder(w).Encode(result)
}