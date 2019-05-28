package resolver

import (
	"context"

	"github.com/peterhellberg/swapi"

	"gqlgen-starwars/errors"
	"gqlgen-starwars/utils"
)

type queryResolver struct{ *Resolver }

func (r *queryResolver) Character(ctx context.Context, id string) (*swapi.Person, error) {
	logger := utils.GetLogger(ctx)

	personId, err := utils.ParseId(id)
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

	filmId, err := utils.ParseId(id)
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
		return nil, errors.NewResourceNotFoundError("Film not found", "Film", id)
	}

	return &film, nil
}
