package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/me/todo-api/models"
	"github.com/me/todo-api/repository"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	var rsp map[string]any

	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		log.Printf("Error decoding json: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	id, err := repository.Insert(todo)
	if err != nil {
		rsp = map[string]any{
			"Error":   true,
			"Message": fmt.Sprintf("Error trying to create: %v", err),
		}
	} else {
		rsp = map[string]any{
			"Error":   false,
			"Message": fmt.Sprintf("Created with success! ID: %d", id),
		}
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rsp)
}
