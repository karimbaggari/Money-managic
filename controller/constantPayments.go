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

func (h *Handlers) EnterFinances(w http.ResponseWriter, r *http.Request) {
	var finance model.UserFinance

	err := json.NewDecoder(r.Body).Decode(&finance)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := h.Service.EnterFinances(finance)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)

}

func (h *Handlers) AboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the About Page!")
}
