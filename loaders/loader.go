package loaders

import (
	"context"
	"log"
	"sync"

	"github.com/peterhellberg/swapi"
)

type contextKey string

const (
	filmLoaderKey   contextKey = "film_loader"
	personLoaderKey contextKey = "person_loader"
)

type Collection struct {
	lookup map[contextKey]interface{}
}

func (c *Collection) Attach(ctx context.Context) context.Context {
	for key, loader := range c.lookup {
		ctx = context.WithValue(ctx, key, loader)
	}

	return ctx
}

func Initialize(client *swapi.Client) *Collection {
	return &Collection{
		lookup: map[contextKey]interface{}{
			filmLoaderKey:   newFilmLoader(client),
			personLoaderKey: newPersonLoader(client),
		},
	}
}

func GetFilmLoader(ctx context.Context) *FilmLoader {
	if l, ok := ctx.Value(filmLoaderKey).(*FilmLoader); ok {
		return l
	}

	return nil
}

func GetPersonLoader(ctx context.Context) *PersonLoader {
	if l, ok := ctx.Value(personLoaderKey).(*PersonLoader); ok {
		return l
	}

	return nil
}

func newFilmLoader(client *swapi.Client) *FilmLoader {
	return NewFilmLoader(FilmLoaderConfig{
		Fetch: func(ids []int) ([]*swapi.Film, []error) {
			var (
				n      = len(ids)
				films  = make([]*swapi.Film, n)
				errors = make([]error, n)
				wg     sync.WaitGroup
			)

			wg.Add(n)

			for i, id := range ids {
				go func(index, id int) {
					defer wg.Done()

					log.Print("Fetching film: ", id)

					film, err := client.Film(id)
					if err != nil {
						errors[index] = err
					}

					films[index] = &film
				}(i, id)
			}

			wg.Wait()

			return films, errors
		},
	})
}

func newPersonLoader(client *swapi.Client) *PersonLoader {
	return NewPersonLoader(PersonLoaderConfig{
		Fetch: func(ids []int) ([]*swapi.Person, []error) {
			var (
				n      = len(ids)
				people = make([]*swapi.Person, n)
				errs   = make([]error, n)
				wg     sync.WaitGroup
			)

			wg.Add(n)

			for i, id := range ids {
				go func(index, id int) {
					defer wg.Done()

					log.Println("Fetching person: ", id)

					person, err := client.Person(id)
					if err != nil {
						errs[index] = err

						return
					}

					people[index] = &person
				}(i, id)
			}

			wg.Wait()

			return people, errs
		},
	})
}
