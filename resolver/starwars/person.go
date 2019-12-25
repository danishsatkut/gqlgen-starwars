package starwars

import (
	"context"

	"github.com/peterhellberg/swapi"

	"gqlgen-starwars/errors"
	"gqlgen-starwars/loaders"
	swapi2 "gqlgen-starwars/swapi"
	"gqlgen-starwars/utils"
)

type personResolver struct {
	client *swapi.Client
}

func NewPersonResolver(client *swapi.Client) *personResolver {
	return &personResolver{client}
}

func (*personResolver) ID(ctx context.Context, p *swapi.Person) (string, error) {
	return utils.ID(ctx, p.URL)
}

func (r *personResolver) Films(ctx context.Context, p *swapi.Person) ([]*swapi.Film, error) {
	entry := utils.GetLogEntry(ctx)
	ids := make([]int, 0, len(p.FilmURLs))

	for _, url := range p.FilmURLs {
		id, err := swapi2.ResourceId(string(url))
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

	return films, nil
}
