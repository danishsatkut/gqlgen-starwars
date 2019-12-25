package starwars

import (
	"context"

	"github.com/peterhellberg/swapi"

	"gqlgen-starwars/errors"
	"gqlgen-starwars/loaders"
	"gqlgen-starwars/server/middlewares"
	swapi2 "gqlgen-starwars/swapi"
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
	entry := middlewares.GetLogEntry(ctx)
	ids := make([]int, 0, len(f.CharacterURLs))

	for _, url := range f.CharacterURLs {
		id, err := swapi2.ResourceId(string(url))
		if err != nil {
			entry.WithError(err).Error("Failed to parse id from url")

			return nil, errors.NewParsingError(err)
		}

		ids = append(ids, id)
	}

	characters, errs := loaders.GetPersonLoader(ctx).LoadAll(ids)
	if len(errs) > 0 && errs[0] != nil {
		return nil, errs[0]
	}

	return characters, nil
}
