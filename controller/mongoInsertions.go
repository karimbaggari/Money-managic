package controllers

// import (
// 	"encoding/json"
// 	"money-managic/model"
// 	"money-managic/services"
// 	"net/http"
// )

// type Handlers struct {
// 	Service *services.Services
// }

// func NewHandlers(service *services.Services) *Handlers {
// 	return &Handlers{Service: service}
// }



// func (h *Handlers) InsertLivingBudget(w http.ResponseWriter, r *http.Request) {
// 	var LivingBudget model.UserLivingBudget

// 	err := json.NewDecoder(r.Body).Decode(&LivingBudget)
// 	if err!= nil {
//             http.Error(w, err.Error(), http.StatusInternalServerError)
//             return
//     }

// 	result, err := h.Service.InsertLivingBudgetService(LivingBudget)
// 	if err!= nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// }

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(result)
// }

// func (h *Handlers) InsertSavings(w http.ResponseWriter, r *http.Request) {
// 	var savings model.UserSavings 
// 	err := json.NewDecoder(r.Body).Decode(savings)
// 	if err!= nil {
//         http.Error(w, err.Error(), http.StatusInternalServerError)
//         return
//     }

// 	result, err := h.Service.InsertSavingsService(savings)
// 	if err!= nil {
//         http.Error(w, err.Error(), http.StatusInternalServerError)
//         return
// }

// w.Header().Set("Content-Type", "application/json")
// w.WriteHeader(http.StatusOK)
// json.NewEncoder(w).Encode(result)
// }