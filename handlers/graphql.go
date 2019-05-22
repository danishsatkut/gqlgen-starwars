package handlers

import (
	"context"
	"net/http"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/handler"
	"github.com/google/uuid"
	"github.com/peterhellberg/swapi"
	"github.com/sirupsen/logrus"

	"gqlgen-starwars"
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
		logger:      logrus.New(),
	}

	for _, option := range options {
		option(cfg)
	}

	config := gqlgen_starwars.Config{
		Resolvers: resolver.NewRootResolver(cfg.swapiClient),
	}

	return handler.GraphQL(
		gqlgen_starwars.NewExecutableSchema(config),
		loggerMiddleware(cfg.logger))
}

func loggerMiddleware(l *logrus.Logger) handler.Option {
	return handler.RequestMiddleware(func(ctx context.Context, next func(ctx context.Context) []byte) []byte {
		rctx := graphql.GetRequestContext(ctx)
		logger := l.WithField("request_id", uuid.New())
		ctx = utils.WithLogger(ctx, logger)

		logger.WithField("query", rctx.RawQuery).WithField("variables", rctx.Variables).Info("Executing GraphQL query")
		res := next(ctx)
		logger.WithField("errors", len(rctx.Errors)).Info("Finished query execution")

		return res
	})
}
