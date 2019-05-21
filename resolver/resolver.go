package resolver

import (
	"github.com/peterhellberg/swapi"

	"gqlgen-starwars"
	"gqlgen-starwars/resolver/starwars"
)

type Resolver struct{
	client *swapi.Client
}

func NewRootResolver(client *swapi.Client) *Resolver {
	return &Resolver{client}
}

func (r *Resolver) Query() gqlgen_starwars.QueryResolver {
	return &queryResolver{r}
}

func (r *Resolver) Film() gqlgen_starwars.FilmResolver {
	return starwars.NewFilmResolver(r.client)
}

func (r *Resolver) Person() gqlgen_starwars.PersonResolver {
	return starwars.NewPersonResolver(r.client)
}
