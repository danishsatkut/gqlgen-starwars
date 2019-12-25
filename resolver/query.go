package resolver

import (
	"context"
	"fmt"
	"strconv"

	"github.com/peterhellberg/swapi"

	"gqlgen-starwars/errors"
	"gqlgen-starwars/loaders"
	"gqlgen-starwars/server/middlewares"
)

type queryResolver struct{ *Resolver }

func (r *queryResolver) Character(ctx context.Context, id string) (*swapi.Person, error) {
	entry := middlewares.GetLogEntry(ctx)

	personId, err := parseId(id)
	if err != nil {
		middlewares.GetLogEntry(ctx).WithError(err).Error("Failed to parse character id")

		return nil, err
	}

	person, err := loaders.GetPersonLoader(ctx).Load(personId)
	if err != nil {
		entry.WithError(err).Error("Failed to fetch person")

		return nil, errors.NewAPIError(err)
	}

	if person.URL == "" {
		return nil, errors.NewResourceNotFoundError("Character not found", "Character", id)
	}

	return person, nil
}

func (r *queryResolver) Film(ctx context.Context, id string) (*swapi.Film, error) {
	entry := middlewares.GetLogEntry(ctx)

	filmId, err := parseId(id)
	if err != nil {
		middlewares.GetLogEntry(ctx).WithError(err).Error("Failed to parse film id")

		return nil, err
	}

	film, err := loaders.GetFilmLoader(ctx).Load(filmId)
	if err != nil {
		entry.WithError(err).Error("Failed to fetch film")

		return nil, errors.NewAPIError(err)
	}

	if film.URL == "" {
		return nil, errors.NewResourceNotFoundError("Film not found", "Film", id)
	}

	return film, nil
}

func parseId(id string) (int, error) {
	resourceId, err := strconv.Atoi(id)
	if err != nil {
		return 0, errors.NewUserInputError(fmt.Sprintf("Invalid id: %v", id), "id")
	}

	return resourceId, nil
}
