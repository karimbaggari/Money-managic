package controllers

import (
	"encoding/json"
	"fmt"
	"money-managic/model"
	"money-managic/services"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

type IncomesHandler struct { 
	IncomesService *services.IncomesService
}

func NewIncomesHandler(incomesHandler *services.IncomesService) *IncomesHandler {
	return &IncomesHandler{IncomesService: incomesHandler}
}


func (h *IncomesHandler) EnterIncomes(w http.ResponseWriter, r *http.Request) {
	var finance model.UserIncomes

	err := json.NewDecoder(r.Body).Decode(&finance)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Printf("Received finance object: %+v\n", finance)

	result, err := h.IncomesService.EnterIncomesService(finance)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	insertResult, ok := result.(*mongo.InsertOneResult)
	if !ok {
		http.Error(w, "Failed to insert document", http.StatusInternalServerError)
		return
	}


	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(insertResult)
}