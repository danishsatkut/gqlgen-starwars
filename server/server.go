package server

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/sirupsen/logrus"

	"github.com/99designs/gqlgen/handler"

	"gqlgen-starwars/server/handlers"
	"gqlgen-starwars/server/middlewares"
)

// Tweak configuration values here.
const (
	addr              = ":8080"
	readHeaderTimeout = 1 * time.Second
	writeTimeout      = 300 * time.Second
	idleTimeout       = 300 * time.Second
	maxHeaderBytes    = http.DefaultMaxHeaderBytes
)

func NewServer(logger *logrus.Logger) *http.Server {
	// Change middleware logger to use logrus logger
	middlewares.SetDefaultLogger(logger)

	router := chi.NewRouter()

	// Use logger and heartbeat middlewares
	router.Use(middleware.Heartbeat("/ping"))

	// GraphQL Playground
	router.Handle("/playground", handler.Playground("GraphQL playground", "/"))

	// GraphQL Query Endpoint
	router.With(
		middlewares.RequestID,
		middlewares.Auth,
		middleware.Logger,
		middlewares.LogEntry(logger),
	).Handle("/", handlers.NewGraphQlHandler(handlers.Logger(logger)))

	// Configure the HTTP server.
	s := &http.Server{
		Addr:              addr,
		Handler:           router,
		ReadHeaderTimeout: readHeaderTimeout,
		WriteTimeout:      writeTimeout,
		IdleTimeout:       idleTimeout,
		MaxHeaderBytes:    maxHeaderBytes,
	}

	return s
}
