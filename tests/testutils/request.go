package testutils

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"gqlgen-starwars/handlers"
)

func NewGraphQLRequest(t *testing.T, query string) *http.Request {
	gqlQuery := fmt.Sprintf("{\"query\":\"%v\"}", query)

	req, err := http.NewRequest("POST", "/query", strings.NewReader(gqlQuery))
	if err != nil {
		t.Fatal(err.Error())
	}

	req.Header.Set("Content-Type", "application/json")

	return req
}

func PerformGraphQLRequest(res http.ResponseWriter, req *http.Request, opts ...handlers.Option) {
	handler := handlers.NewGraphQlHandler(opts...)
	handler.ServeHTTP(res, req)
}
