package main

import (
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"money-managic/controller"
	"money-managic/database"
	"money-managic/repositories"
	"money-managic/services"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	database.Connect()
	dbName := os.Getenv("DATABASE_NAME")
	dbCollection1 := os.Getenv("DATABASE_COLLECTION1")

	collection := database.GetCollection(dbName, dbCollection1)
	repo := repositories.NewRepository(collection)
	logic := services.NewServices(repo)
	handlers := controllers.NewHandlers(logic)

	r := mux.NewRouter()
	r.HandleFunc("/enter-finance", handlers.EnterFinances).Methods("POST")

	log.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Println(err)
	}
}
