package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/wigwamwam/CRUD_app/handlers"
)

func main() {

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/banks", handlers.HandlerIndexBanks())
	r.Post("/banks", handlers.CreateBank())
	r.Get("/banks/{id}", handlers.ShowBank())
	r.Delete("/banks/{id}", handlers.DeleteBank())
	r.Put("/banks/{id}", handlers.UpdateBank())

	http.ListenAndServe(":3000", r)

}
