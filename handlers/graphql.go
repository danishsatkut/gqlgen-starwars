package handlers

import (
	"net/http"

	"github.com/99designs/gqlgen/handler"
	"github.com/peterhellberg/swapi"

	"gqlgen-starwars"
	"gqlgen-starwars/resolver"
)

func NewGraphQlHandler() http.Handler {
	config := gqlgen_starwars.Config{Resolvers: resolver.NewRootResolver(swapi.DefaultClient)}

	return handler.GraphQL(gqlgen_starwars.NewExecutableSchema(config))
}
