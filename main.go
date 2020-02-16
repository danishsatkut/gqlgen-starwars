package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"gqlgen-starwars/server"
	"gqlgen-starwars/server/middlewares"
)

func main() {
	var (
		logger = middlewares.DefaultLogger
		srv    = server.NewServer(logger)
	)

	logger.Infof("Listening for requests on %s", srv.Addr)

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		// Begin listening for requests.
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			logger.WithError(err).Error("Failed to listen and serve")
		}
	}()

	if err := gracefulShutdown(srv); err != nil {
		logger.WithError(err).Error("Failed to cleanly shutdown server")
	}

	logger.Info("Shutting down.")
	os.Exit(0)
}

func gracefulShutdown(srv *http.Server) error {
	var (
		c    = make(chan os.Signal, 1)
		wait = 15 * time.Second
	)

	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	return srv.Shutdown(ctx)
}
