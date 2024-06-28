package main

import (
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"money-managic/controller"
	"money-managic/database"
	"money-managic/repositories/MongoRepo"
	memory "money-managic/repositories/MemoryRepo"
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
	dbCollection2 := os.Getenv("DATABASE_COLLECTION2")
	dbCollection3 := os.Getenv("DATABASE_COLLECTION3")

	collection := database.GetCollection(dbName, dbCollection1)

	incomesRepo := repositories.NewIncomesRepository(collection)
	incomesService := services.NewIncomesService(incomesRepo)
	incomesHandlers := controllers.NewIncomesHandler(incomesService)

	collection = database.GetCollection(dbName, dbCollection2)

	expenseRepo := repositories.NewExpensesRepository(collection)
	expensesService := services.NewExpensesServices(expenseRepo)
	expensesHandlers := controllers.NewExpensesHandler(expensesService)

	collection = database.GetCollection(dbName, dbCollection3)

	livingBudgetRepo := repositories.NewLivingBudgetRepository(collection)
	livingBudgetMemoryRepo := memory.NewInMemoryRepository()
	livingBudgetService := services.NewLivingBudgetServices(livingBudgetRepo, livingBudgetMemoryRepo)
	livingBudgetHandler := controllers.NewLivingBudgetHandler(livingBudgetService)

	//inMemoryRepo :=repositories.NewInMemoryRepository()
	//memoryService := services.NewServices(inMemoryRepo)
	//memoryHandler := controllers.NewCalculationHandler(memoryService)

	r := mux.NewRouter()
	r.HandleFunc("/enter-incomes", incomesHandlers.EnterIncomes).Methods("POST")
	r.HandleFunc("/enter-expenses", expensesHandlers.EnterExpenses).Methods("POST")
	r.HandleFunc("/insert-LivingBudget", livingBudgetHandler.LivingBudget).Methods("POST")
	//r.HandleFunc("/insert-Savings", handlers.InsertSavings).Methods("POST")
	//r.HandleFunc("/store-savings-in-address", memoryHandler.LivingBudget).Methods("POST")
	//r.HandleFunc("/store-living-budget-in-address", memoryHandler.SaveBudget).Methods("POST")



	log.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Println(err)
	}
}
