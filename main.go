package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/stephan-lopes/to-do/configs"
	"github.com/stephan-lopes/to-do/handlers"
)

func main() {
	err := configs.Load()
	if err != nil {
		log.Fatalf("Error loading configs: %v", err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/", handlers.Create).Methods("POST")
	r.HandleFunc("/", handlers.List).Methods("GET")
	r.HandleFunc("/{id}", handlers.Update).Methods("PUT")
	r.HandleFunc("/{id}", handlers.Delete).Methods("DELETE")
	r.HandleFunc("/{id}", handlers.Get).Methods("GET")

	fmt.Println("Server running on port:", configs.GetServerPort())
	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)
}
