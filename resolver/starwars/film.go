package starwars

import (
	"context"

	"github.com/peterhellberg/swapi"

	"gqlgen-starwars/utils"
)

type filmResolver struct {
	client *swapi.Client
}

func NewFilmResolver(client *swapi.Client) *filmResolver {
	return &filmResolver{client}
}

func (*filmResolver) ID(ctx context.Context, film *swapi.Film) (string, error) {
	return film.URL, nil
}

func (*filmResolver) Name(ctx context.Context, film *swapi.Film) (string, error) {
	return film.Title, nil
}

func (r *filmResolver) Characters(ctx context.Context, film *swapi.Film) ([]*swapi.Person, error) {
	characters := make([]*swapi.Person, 0, len(film.CharacterURLs))

	for _, url := range film.CharacterURLs {
		id, err := utils.Id(string(url))
		if err != nil {
			return nil, err
		}

		p, err := r.client.Person(id)
		if err != nil {
			return nil, err
		}

		characters = append(characters, &p)
	}

	return characters, nil
}
