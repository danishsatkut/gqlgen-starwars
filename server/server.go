package main

import (
	"github.com/peterhellberg/swapi"

	"gqlgen-starwars"
	"gqlgen-starwars/resolver"

	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	config := gqlgen_starwars.Config{Resolvers: resolver.NewRootResolver(swapi.DefaultClient)}

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(gqlgen_starwars.NewExecutableSchema(config)))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
