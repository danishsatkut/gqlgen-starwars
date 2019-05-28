package starwars

import (
	"context"

	"github.com/peterhellberg/swapi"

	"gqlgen-starwars/errors"
	"gqlgen-starwars/utils"
)

type filmResolver struct {
	client *swapi.Client
}

func NewFilmResolver(client *swapi.Client) *filmResolver {
	return &filmResolver{client}
}

func (*filmResolver) ID(ctx context.Context, f *swapi.Film) (string, error) {
	return utils.ID(ctx, f.URL)
}

func (*filmResolver) Name(ctx context.Context, f *swapi.Film) (string, error) {
	return f.Title, nil
}

func (r *filmResolver) Characters(ctx context.Context, f *swapi.Film) ([]*swapi.Person, error) {
	logger := utils.GetLogger(ctx)
	characters := make([]*swapi.Person, 0, len(f.CharacterURLs))

	for _, url := range f.CharacterURLs {
		id, err := utils.ResourceId(string(url))
		if err != nil {
			logger.WithError(err).Error("Failed to parse id from url")

			return nil, errors.NewParsingError(err)
		}

		logger.WithField("id", id).Debug("Fetching character")

		p, err := r.client.Person(id)
		if err != nil {
			logger.WithError(err).Error("Failed to fetch person")

			return nil, errors.NewAPIError(err)
		}

		characters = append(characters, &p)
	}

	return characters, nil
}
