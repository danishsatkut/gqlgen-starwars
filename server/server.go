package main

import (
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

	logger := utils.DefaultLogger

	// Change middleware logger to use logrus logger
	middlewares.SetDefaultLogger(logger)

	router := chi.NewRouter()

	// Use logger and heartbeat middlewares
	router.Use(middleware.Logger, middleware.Heartbeat("/ping"))

	// GraphQL Playground
	router.Handle("/playground", handler.Playground("GraphQL playground", "/"))

	// GraphQL Query Endpoint
	router.With(middlewares.RequestID, middlewares.Auth).Handle("/", handlers.NewGraphQlHandler(handlers.Logger(logger)))

	// TODO: Gracefully shutdown server
	logger.Infof("connect to http://localhost:%s/playground for GraphQL playground", port)
	logger.Fatal(http.ListenAndServe(":"+port, router))
}
