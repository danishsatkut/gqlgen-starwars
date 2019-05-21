package handlers

import (
	"net/http"

	"github.com/99designs/gqlgen/handler"
	"github.com/peterhellberg/swapi"

	"gqlgen-starwars"
	"gqlgen-starwars/resolver"
)

type Config struct {
	swapiClient *swapi.Client
}

type Option func(cfg *Config)

func SwapiClient(c *swapi.Client) Option {
	return func(cfg *Config) {
		cfg.swapiClient = c
	}
}

func NewGraphQlHandler(options ...Option) http.Handler {
	cfg := &Config{
		swapiClient: swapi.DefaultClient,
	}

	for _, option := range options {
		option(cfg)
	}

	config := gqlgen_starwars.Config{
		Resolvers: resolver.NewRootResolver(cfg.swapiClient),
	}

	return handler.GraphQL(gqlgen_starwars.NewExecutableSchema(config))
}
