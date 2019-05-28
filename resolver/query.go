package resolver

import (
	"context"
	"fmt"
	"strconv"

	"github.com/peterhellberg/swapi"

	"gqlgen-starwars/errors"
	"gqlgen-starwars/utils"
)

type queryResolver struct{ *Resolver }

func (r *queryResolver) Character(ctx context.Context, id string) (*swapi.Person, error) {
	logger := utils.GetLogger(ctx)

	personId, err := parseId(id)
	if err != nil {
		logger.WithError(err).Error("Failed to parse character id")

		return nil, err
	}

	person, err := r.client.Person(personId)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to fetch person with id: %v", personId)
	}

	return &person, nil
}

func (r *queryResolver) Film(ctx context.Context, id string) (*swapi.Film, error) {
	logger := utils.GetLogger(ctx)

	filmId, err := parseId(id)
	if err != nil {
		logger.WithError(err).Error("Failed to parse film id")

		return nil, err
	}

	film, err := r.client.Film(filmId)
	if err != nil {
		logger.WithError(err).Error("Failed to fetch film")

		return nil, errors.New("Something went wrong!")
	}

	if film.URL == "" {
		return nil, errors.New("Film not found!")
	}

	return &film, nil
}

func parseId(id string) (int, error) {
	resourceId, err := strconv.Atoi(id)
	if err != nil {
		return 0, errors.NewUserInputError(fmt.Sprintf("Invalid id: %v", id), "id")
	}

	return resourceId, nil
}
