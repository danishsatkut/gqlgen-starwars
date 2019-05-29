package resolver

import (
	"context"

	"github.com/peterhellberg/swapi"

	"gqlgen-starwars/errors"
	"gqlgen-starwars/utils"
)

type queryResolver struct{ *Resolver }

func (r *queryResolver) Character(ctx context.Context, id string) (*swapi.Person, error) {
	entry := utils.GetLogEntry(ctx)

	personId, err := utils.ParseId(ctx, id)
	if err != nil {
		return nil, err
	}

	person, err := r.client.Person(personId)
	if err != nil {
		entry.WithError(err).Error("Failed to fetch person")

		return nil, errors.NewAPIError(err)
	}

	if person.URL == "" {
		return nil, errors.NewResourceNotFoundError("Character not found", "Character", id)
	}

	return &person, nil
}

func (r *queryResolver) Film(ctx context.Context, id string) (*swapi.Film, error) {
	entry := utils.GetLogEntry(ctx)

	filmId, err := utils.ParseId(ctx, id)
	if err != nil {
		return nil, err
	}

	film, err := r.client.Film(filmId)
	if err != nil {
		entry.WithError(err).Error("Failed to fetch film")

		return nil, errors.NewAPIError(err)
	}

	if film.URL == "" {
		return nil, errors.NewResourceNotFoundError("Film not found", "Film", id)
	}

	return &film, nil
}
