package utils

import (
	"context"
	"fmt"
	"strconv"

	"gqlgen-starwars/errors"
)

func ParseId(ctx context.Context, id string) (int, error) {
	resourceId, err := strconv.Atoi(id)
	if err != nil {
		GetLogEntry(ctx).WithError(err).Error("Failed to parse id")

		return 0, errors.NewUserInputError(fmt.Sprintf("Invalid id: %v", id), "id")
	}

	return resourceId, nil
}

// ID returns graphql ID field value for a url
func ID(ctx context.Context, url string) (string, error) {
	id, err := ResourceId(url)
	if err != nil {
		GetLogEntry(ctx).Error("Failed to extract id from resource url")

		return "", errors.NewParsingError(err)
	}

	return strconv.Itoa(id), nil
}
