package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/me/activities/repository"
)

func List(w http.ResponseWriter, r *http.Request) {
	todos, err := repository.ReadAll()
	if err != nil {
		log.Printf("Error getting registries")
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}
