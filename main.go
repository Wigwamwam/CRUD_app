package main

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/wigwamwam/CRUD_app/controllers"
	"github.com/wigwamwam/CRUD_app/initializers"
)

func init() {
	initializers.LoadEnvVairables()
	initializers.ConnectToDB()
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/banks", controllers.IndexBanks())
	r.Post("/banks", controllers.CreateBank())
	r.Get("/banks/{id}", controllers.ShowBank())
	r.Delete("/banks/{id}", controllers.DeleteBank())
	r.Put("/banks/{id}", controllers.UpdateBank())

	http.ListenAndServe(":3000", r)

}
