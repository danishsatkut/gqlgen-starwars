package starwars

import (
	"context"

	"github.com/peterhellberg/swapi"
)

type personResolver struct {
	client *swapi.Client
}

func (personResolver) ID(ctx context.Context, obj *swapi.Person) (string, error) {
	return obj.URL, nil
}

func NewPersonResolver(client *swapi.Client) *personResolver {
	return &personResolver{client}
}
