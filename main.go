package main

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to the Home Page!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "This is the About Page!")
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/", homeHandler)
    r.HandleFunc("/about", aboutHandler)

    fmt.Println("Starting server at port 8080")
    if err := http.ListenAndServe(":8080", r); err != nil {
        fmt.Println(err)
    }
}
