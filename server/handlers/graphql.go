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

	"gqlgen-starwars/errors"
	"gqlgen-starwars/generated"
	"gqlgen-starwars/loaders"
	"gqlgen-starwars/resolver"
	"gqlgen-starwars/utils"
)

func NewGraphQlHandler(options ...Option) http.Handler {
	cfg := &config{
		swapiClient: swapi.DefaultClient,
		logger:      utils.DefaultLogger,
	}

	cfg.update(options...)

	config := generated.Config{
		Resolvers: resolver.NewRootResolver(cfg.swapiClient),
	}

	return handler.GraphQL(
		generated.NewExecutableSchema(config),
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
