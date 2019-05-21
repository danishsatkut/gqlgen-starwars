package resolver

import (
	"context"
	"strconv"

	"github.com/peterhellberg/swapi"
	"github.com/pkg/errors"
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
	filmId, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	film, err := r.client.Film(filmId)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to fetch film")
	}

	return &film, nil
}
