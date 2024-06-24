package controllers

import (
    "context"
    "fmt"
    "net/http"
    "money-managic/repositories"
    "go.mongodb.org/mongo-driver/bson"
)

type Handlers struct {
    Repo *repositories.Repository
}

func NewHandlers(repo *repositories.Repository) *Handlers {
    return &Handlers{Repo: repo}
}

func (h *Handlers) HomeHandler(w http.ResponseWriter, r *http.Request) {
    doc := bson.D{{"name", "Home"}, {"message", "Welcome to the Home Page!"}}

    result, err := h.Repo.InsertDocument(context.TODO(), doc)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    fmt.Fprintf(w, "Inserted document with ID: %v", result.InsertedID)
}

func (h *Handlers) AboutHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "This is the About Page!")
}
