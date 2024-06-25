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
	mongoService := services.NewServices(repo)
	handlers := controllers.NewHandlers(mongoService)

	inMemoryRepo :=repositories.NewInMemoryRepository()
	memoryService := services.NewServices(inMemoryRepo)
	memoryHandler := controllers.NewCalculationHandler(memoryService)

	r := mux.NewRouter()
	r.HandleFunc("/enter-incomes", handlers.EnterIncomes).Methods("POST")
	r.HandleFunc("/enter-expenses", handlers.EnterExpenses).Methods("POST")
	r.HandleFunc("/insert-LivingBudget", handlers.InsertLivingBudget).Methods("POST")
	r.HandleFunc("/insert-Savings", handlers.InsertSavings).Methods("POST")
	r.HandleFunc("/store-in-address", memoryHandler.EnterIncomes).Methods("POST")


	log.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Println(err)
	}
}
