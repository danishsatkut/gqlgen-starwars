package main

import (
	"context"
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
		wait   = 15 * time.Second
	)

	logger.Infof("Listening for requests on %s", srv.Addr)

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		// Begin listening for requests.
		if err := srv.ListenAndServe(); err != nil {
			logger.WithError(err).Error("Failed to listen and serve")
		}
	}()

	c := make(chan os.Signal, 1)
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
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	logger.Info("Shutting down.")
	os.Exit(0)
}
