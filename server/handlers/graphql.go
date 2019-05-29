package handlers

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"runtime/debug"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/handler"
	"github.com/peterhellberg/swapi"
	"github.com/sirupsen/logrus"

	"gqlgen-starwars"
	"gqlgen-starwars/errors"
	"gqlgen-starwars/loaders"
	"gqlgen-starwars/resolver"
	"gqlgen-starwars/utils"
)

type Config struct {
	swapiClient *swapi.Client
	logger      *logrus.Logger
}

type Option func(cfg *Config)

func SwapiClient(c *swapi.Client) Option {
	return func(cfg *Config) {
		cfg.swapiClient = c
	}
}

func Logger(l *logrus.Logger) Option {
	return func(cfg *Config) {
		cfg.logger = l
	}
}

func NewGraphQlHandler(options ...Option) http.Handler {
	cfg := &Config{
		swapiClient: swapi.DefaultClient,
		logger:      utils.DefaultLogger,
	}

	for _, option := range options {
		option(cfg)
	}

	config := gqlgen_starwars.Config{
		Resolvers: resolver.NewRootResolver(cfg.swapiClient),
	}

	return handler.GraphQL(
		gqlgen_starwars.NewExecutableSchema(config),
		loggerMiddleware(),
		panicMiddleware(),
		dataloaderMiddleware(cfg.swapiClient))
}

func loggerMiddleware() handler.Option {
	return handler.RequestMiddleware(func(ctx context.Context, next func(ctx context.Context) []byte) []byte {
		logger := utils.GetLogEntry(ctx)
		rctx := graphql.GetRequestContext(ctx)

		logger.WithField("query", rctx.RawQuery).WithField("variables", rctx.Variables).Info("Executing GraphQL query")
		res := next(ctx)
		logger.WithField("errors", len(rctx.Errors)).Info("Finished query execution")

		return res
	})
}

func panicMiddleware() handler.Option {
	return handler.RecoverFunc(func(ctx context.Context, err interface{}) (userMessage error) {
		fmt.Fprintln(os.Stderr, err)
		fmt.Fprintln(os.Stderr)

		errStack := string(debug.Stack())
		fmt.Fprintln(os.Stderr, errStack)

		return errors.NewServerError(err, errStack)
	})
}

func dataloaderMiddleware(client *swapi.Client) handler.Option {
	fn := func(ctx context.Context, next func(ctx context.Context) []byte) []byte {
		ctx = loaders.Initialize(client).Attach(ctx)

		return next(ctx)
	}

	return handler.RequestMiddleware(fn)
}
