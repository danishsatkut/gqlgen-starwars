package starwars

import (
	"context"
	"strconv"

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
