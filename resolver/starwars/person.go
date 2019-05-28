package starwars

import (
	"context"
	"log"

	"github.com/peterhellberg/swapi"

	"gqlgen-starwars/errors"
	"gqlgen-starwars/utils"
)

type personResolver struct {
	client *swapi.Client
}

func NewPersonResolver(client *swapi.Client) *personResolver {
	return &personResolver{client}
}

func (*personResolver) ID(ctx context.Context, p *swapi.Person) (string, error) {
	return utils.ID(p.URL)
}

func (r *personResolver) Films(ctx context.Context, p *swapi.Person) ([]*swapi.Film, error) {
	logger := utils.GetLogger(ctx)
	films := make([]*swapi.Film, 0, len(p.FilmURLs))

	for _, url := range p.FilmURLs {
		id, err := utils.ResourceId(string(url))
		if err != nil {
			logger.WithError(err).Error("Failed to parse id from url")

			return nil, errors.NewParsingError(err)
		}

		log.Print("Fetching film: ", id)

		film, err := r.client.Film(id)
		if err != nil {
			logger.WithError(err).Error("Failed to fetch person")

			return nil, errors.NewAPIError(err)
		}

		films = append(films, &film)
	}

	return films, nil
}
