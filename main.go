package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/me/todo-api/configs"
	"github.com/me/todo-api/handlers"
)

func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	r.Post("/", handlers.Create)
	r.Put("/{id}", handlers.Update)
	r.Delete("/{id}", handlers.Delete)
	r.Get("/", handlers.Read)
	r.Get("/{id}", handlers.Read)

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)

}
