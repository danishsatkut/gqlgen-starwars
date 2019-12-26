package starwars

import (
	"context"

	"github.com/peterhellberg/swapi"

	"gqlgen-starwars/errors"
	"gqlgen-starwars/loaders"
	"gqlgen-starwars/server/middlewares"
	swapihelper "gqlgen-starwars/swapi"
)

type filmResolver struct {
	client *swapi.Client
}

func NewFilmResolver(client *swapi.Client) *filmResolver {
	return &filmResolver{client}
}

func (*filmResolver) ID(ctx context.Context, f *swapi.Film) (string, error) {
	return ID(ctx, f.URL)
}

func (*filmResolver) Name(ctx context.Context, f *swapi.Film) (string, error) {
	return f.Title, nil
}

func (r *filmResolver) Characters(ctx context.Context, f *swapi.Film) ([]*swapi.Person, error) {
	urls := make([]string, 0, len(f.CharacterURLs))
	for _, url := range f.CharacterURLs {
		urls = append(urls, string(url))
	}

	characters, err := getCharacters(ctx, urls)
	if err != nil {
		return nil, err
	}

	if isFieldRequested(ctx, "films") && len(characters) > 0 {
		urlMap := make(map[string]interface{}, 0)
		for _, character := range characters {
			for _, url := range character.FilmURLs {
				urlMap[string(url)] = nil
			}
		}

		urls := make([]string, 0, len(urlMap))
		for url, _ := range urlMap {
			urls = append(urls, url)
		}

		// Load and Prime cache
		// NOTE: This is blocking call
		getFilms(ctx, urls)
	}

	return characters, err
}

func getFilms(ctx context.Context, urls []string) ([]*swapi.Film, error) {
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

	films, errs := loaders.GetFilmLoader(ctx).LoadAll(ids)
	if len(errs) > 0 && errs[0] != nil {
		return nil, errors.NewAPIError(errs[0])
	}

	return films, nil
}
