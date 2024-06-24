package main

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
	"money-managic/controller"
)


func main() {
    r := mux.NewRouter()
    r.HandleFunc("/", controller.HomeHandler)
    r.HandleFunc("/about", controller.AboutHandler)

    fmt.Println("Starting server at port 8080")
    if err := http.ListenAndServe(":8080", r); err != nil {
        fmt.Println(err)
    }
}
