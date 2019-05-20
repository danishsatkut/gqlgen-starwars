package starwars

import (
	"context"

	"github.com/peterhellberg/swapi"
)

type filmResolver struct {
	client *swapi.Client
}

func (filmResolver) ID(ctx context.Context, film *swapi.Film) (string, error) {
	return film.URL, nil
}

func (filmResolver) Name(ctx context.Context, film *swapi.Film) (string, error) {
	return film.Title, nil
}

func NewFilmResolver(client *swapi.Client) *filmResolver {
	return &filmResolver{client}
}
