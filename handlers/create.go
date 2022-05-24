package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/stephan-lopes/to-do/models"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo

	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		log.Printf("Error decoding JSON: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	id, err := models.Insert(todo)
	var resp map[string]any

	if err != nil {
		resp = map[string]any{
			"Error":   true,
			"Message": fmt.Sprintf("Ocorreu um erro ao inserir: %v", err),
		}
	}

	resp = map[string]any{
		"Error":   false,
		"Message": fmt.Sprintf("Todo inserido com sucesso: %d", id),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
