package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/rohinivsenthil/golang-getting-started/handlers"
)


func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", handlers.GetAllTodos).Methods("GET")
	r.HandleFunc("/", handlers.CreateNewTodo).Methods("POST")
	r.HandleFunc("/{todoId}", handlers.UpdateTodo).Methods("PUT")
	r.HandleFunc("/{todoId}", handlers.DeleteTodo).Methods("DELETE")

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Printf("http server crashed: %s", err.Error())
	}
}
