package main

import (
	"gqlgen-starwars/server"
	"gqlgen-starwars/server/middlewares"
)

func main() {
	logger := middlewares.DefaultLogger

	s := server.NewServer(logger)

	logger.Infof("Listening for requests on %s", s.Addr)

	// Begin listening for requests.
	if err := s.ListenAndServe(); err != nil {
		logger.WithError(err).Error("Failed to listen and serve")
	}

	// TODO: intercept shutdown signals for cleanup of connections.
	logger.Info("Shutting down.")
}
