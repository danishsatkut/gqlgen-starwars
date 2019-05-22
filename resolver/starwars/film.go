package starwars

import (
	"context"
	"log"

	"github.com/peterhellberg/swapi"

	"gqlgen-starwars/utils"
)

type filmResolver struct {
	client *swapi.Client
}

func NewFilmResolver(client *swapi.Client) *filmResolver {
	return &filmResolver{client}
}

func (*filmResolver) ID(ctx context.Context, f *swapi.Film) (string, error) {
	return utils.ID(f.URL)
}

func (*filmResolver) Name(ctx context.Context, f *swapi.Film) (string, error) {
	return f.Title, nil
}

func (r *filmResolver) Characters(ctx context.Context, f *swapi.Film) ([]*swapi.Person, error) {
	characters := make([]*swapi.Person, 0, len(f.CharacterURLs))

	for _, url := range f.CharacterURLs {
		id, err := utils.ResourceId(string(url))
		if err != nil {
			return nil, err
		}

		log.Print("Fetching character: ", id)

		p, err := r.client.Person(id)
		if err != nil {
			return nil, err
		}

		characters = append(characters, &p)
	}

	return characters, nil
}
