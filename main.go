package main

import (
    "log"
    "net/http"
	"os"
    "money-managic/controller"
    "money-managic/database"
    "money-managic/repositories"
    "github.com/joho/godotenv"
    "github.com/gorilla/mux"
)

func main() {
	err := godotenv.Load()
	if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }
    
	database.Connect()
	dbName := os.Getenv("DATABASE-NAME")
	dbCollection1 := os.Getenv("DATABASE-COLLECTION1")


    collection := database.GetCollection(dbName, dbCollection1)
    repo := repositories.NewRepository(collection)
    handlers := controllers.NewHandlers(repo)

    r := mux.NewRouter()
    r.HandleFunc("/", handlers.HomeHandler).Methods("GET")
    r.HandleFunc("/about", handlers.AboutHandler).Methods("GET")

    log.Println("Starting server at port 8080")
    if err := http.ListenAndServe(":8080", r); err != nil {
        log.Println(err)
    }
}
