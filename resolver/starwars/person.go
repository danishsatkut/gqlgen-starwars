package starwars

import (
	"context"

	"github.com/peterhellberg/swapi"

	"gqlgen-starwars/utils"
)

type personResolver struct {
	client *swapi.Client
}

func (personResolver) ID(ctx context.Context, p *swapi.Person) (string, error) {
	return utils.ID(p.URL)
}

func NewPersonResolver(client *swapi.Client) *personResolver {
	return &personResolver{client}
}
