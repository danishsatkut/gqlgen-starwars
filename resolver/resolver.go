package resolver

import (
	"context"

	"gqlgen-starwars"
	"gqlgen-starwars/model"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Query() gqlgen_starwars.QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Film(ctx context.Context, id string) (*model.Film, error) {
	panic("not implemented")
}
