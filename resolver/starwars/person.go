package starwars

import (
	"context"

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
	urls := make([]string, 0, len(p.FilmURLs))
	for _, url := range p.FilmURLs {
		urls = append(urls, string(url))
	}

	films, err := getFilms(ctx, urls)
	if err != nil {
		return nil, err
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

func getCharacters(ctx context.Context, urls []string) ([]*swapi.Person, error) {
	entry := middlewares.GetLogEntry(ctx)
	ids := make([]int, 0, len(urls))

	for _, url := range urls {
		id, err := swapihelper.ResourceId(url)
		if err != nil {
			entry.WithError(err).Error("Failed to parse id from url")

			return nil, errors.NewParsingError(err)
		}

		ids = append(ids, id)
	}

	characters, errs := loaders.GetPersonLoader(ctx).LoadAll(ids)
	if len(errs) > 0 && errs[0] != nil {
		return nil, errors.NewAPIError(errs[0])
	}

	return characters, nil
}
