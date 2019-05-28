package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"gqlgen-starwars/server/handlers"
	"gqlgen-starwars/server/middlewares"

	"github.com/99designs/gqlgen/handler"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Handle("/", handler.Playground("GraphQL playground", "/query"))
	router.With(middlewares.RequestID).Handle("/query", handlers.NewGraphQlHandler())

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
