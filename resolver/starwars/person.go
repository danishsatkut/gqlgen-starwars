package starwars

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/peterhellberg/swapi"

	"gqlgen-starwars/errors"
	"gqlgen-starwars/loaders"
	"gqlgen-starwars/server/middlewares"
	swapihelper "gqlgen-starwars/swapi"
)

type personResolver struct {
	client *swapi.Client
}

func NewPersonResolver(client *swapi.Client) *personResolver {
	return &personResolver{client}
}

func (*personResolver) ID(ctx context.Context, p *swapi.Person) (string, error) {
	return ID(ctx, p.URL)
}

func (r *personResolver) Films(ctx context.Context, p *swapi.Person) ([]*swapi.Film, error) {
	entry := middlewares.GetLogEntry(ctx)
	ids := make([]int, 0, len(p.FilmURLs))

	entry.Debugf("Resolving films for: %s", p.Name)

	for _, url := range p.FilmURLs {
		id, err := swapihelper.ResourceId(string(url))
		if err != nil {
			entry.WithError(err).Error("Failed to parse id from url")

			return nil, errors.NewParsingError(err)
		}

		ids = append(ids, id)
	}

	films, errs := loaders.GetFilmLoader(ctx).LoadAll(ids)
	if len(errs) > 0 && errs[0] != nil {
		return nil, errs[0]
	}

	if isFieldRequested(ctx, "characters") && len(films) > 0 {
		urls := make([]string, 0, len(films) * len(films[0].CharacterURLs))

		for _, film := range films {
			for _, url := range film.CharacterURLs {
				urls = append(urls, string(url))
			}
		}

		// Load and Prime cache for characters
		// NOTE: This will block until all characters are loaded.
		getCharacters(ctx, urls)
	}

	return films, nil
}

func isFieldRequested(ctx context.Context, field string) bool {
	for _, f := range graphql.CollectAllFields(ctx) {
		if f == field {
			return true
		}
	}

	return false
}
