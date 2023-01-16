package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/wigwamwam/CRUD_app/config"
	"github.com/wigwamwam/CRUD_app/handlers"
	"github.com/wigwamwam/CRUD_app/repository"
)

func main() {
	config.LoadEnvVairables()
	dbUrl := os.Getenv("DB_URL")
	pool, err := pgxpool.New(context.Background(), dbUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer pool.Close()

	db := repository.NewDb(pool)
	handler := handlers.NewHandler(db)

	fmt.Println("Successfully connected")

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/banks", handler.HandlerIndexBanks())
	r.Post("/banks", handler.CreateBank())
	r.Get("/banks/{id}", handler.ShowBank())
	r.Delete("/banks/{id}", handler.DeleteBank())
	r.Put("/banks/{id}", handler.UpdateBank())

	http.ListenAndServe(":3000", r)

}
