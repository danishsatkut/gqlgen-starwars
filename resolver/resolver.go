package resolver

import (
	"github.com/peterhellberg/swapi"

	"gqlgen-starwars/generated"
	"gqlgen-starwars/resolver/starwars"
)

type Resolver struct {
	client *swapi.Client
}

func NewRootResolver(client *swapi.Client) *Resolver {
	return &Resolver{client}
}

func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}

func (r *Resolver) Film() generated.FilmResolver {
	return starwars.NewFilmResolver(r.client)
}

func (r *Resolver) Person() generated.PersonResolver {
	return starwars.NewPersonResolver(r.client)
}
