package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/stephan-lopes/to-do/models"
)

func List(w http.ResponseWriter, r *http.Request) {

	todos, err := models.GetAll()
	if err != nil {
		log.Printf("Error getting todos: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}
