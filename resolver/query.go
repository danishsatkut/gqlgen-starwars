package resolver

import (
	"context"
	"strconv"

	"github.com/peterhellberg/swapi"
	"github.com/pkg/errors"

	"gqlgen-starwars/utils"
)

type queryResolver struct{ *Resolver }

func (r *queryResolver) Character(ctx context.Context, id string) (*swapi.Person, error) {
	personId, err := strconv.Atoi(id)
	if err != nil {
		return nil, errors.WithMessagef(err, "Invalid id: %v", id)
	}

	person, err := r.client.Person(personId)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to fetch person with id: %v", personId)
	}

	return &person, nil
}

func (r *queryResolver) Film(ctx context.Context, id string) (*swapi.Film, error) {
	logger := utils.GetLogger(ctx)

	filmId, err := strconv.Atoi(id)
	if err != nil {
		logger.WithError(err).Error("Failed to parse film id")

		return nil, errors.New("Invalid film id")
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
