package starwars

import (
	"context"
	"strconv"

	"github.com/99designs/gqlgen/graphql"

	"gqlgen-starwars/errors"
	"gqlgen-starwars/server/middlewares"
	"gqlgen-starwars/swapi"
)

// ID returns graphql ID field value for a url
func ID(ctx context.Context, url string) (string, error) {
	id, err := swapi.ResourceId(url)
	if err != nil {
		middlewares.GetLogEntry(ctx).Error("Failed to extract id from resource url")

		return "", errors.NewParsingError(err)
	}

	return strconv.Itoa(id), nil
}

func isFieldRequested(ctx context.Context, field string) bool {
	for _, f := range graphql.CollectAllFields(ctx) {
		if f == field {
			return true
		}
	}

	return false
}
