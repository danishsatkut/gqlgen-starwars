package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"gqlgen-starwars/server/handlers"
	"gqlgen-starwars/server/middlewares"
	"gqlgen-starwars/utils"

	"github.com/99designs/gqlgen/handler"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	logger := utils.DefaultLogger()

	router := chi.NewRouter()

	middleware.DefaultLogger = middleware.RequestLogger(&middleware.DefaultLogFormatter{
		Logger: logger,
		NoColor: false,
	})
	router.Use(middleware.Logger)

	router.Handle("/", handler.Playground("GraphQL playground", "/query"))
	router.With(middlewares.RequestID, middlewares.Auth).Handle("/query", handlers.NewGraphQlHandler(handlers.Logger(logger)))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
