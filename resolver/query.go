package resolver

import (
	"context"
	"strconv"

	"github.com/peterhellberg/swapi"
	"github.com/pkg/errors"
)

type queryResolver struct{ *Resolver }

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
