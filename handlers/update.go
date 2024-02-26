package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/me/activities/models"
	"github.com/me/activities/repository"
)

func Update(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Error parsing id: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		log.Printf("Error decoding json: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	rows, err := repository.Update(uint(id), todo)
	if err != nil {
		log.Printf("Error decoding json: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if rows > 1 {
		log.Printf("Error: Updated %d registers", rows)
	}

	rsp := map[string]any{
		"Error":   false,
		"Message": "Updated with success!",
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rsp)

}
