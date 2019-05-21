package testutils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/99designs/gqlgen/graphql"
	"github.com/stretchr/testify/assert"

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

func AssertGraphQLData(t *testing.T, res *httptest.ResponseRecorder, expected string) {
	t.Helper()

	gqlResponse, err := parseGraphQLResponse(res)
	if err != nil {
		t.Errorf("Failed parsing graphql response: %v", err)
	}

	assert.Equal(t, expected, string(gqlResponse.Data), "Response data did not match")
}

func AssertGraphQLErrors(t *testing.T, res *httptest.ResponseRecorder, errs []string) {
	t.Helper()

	gqlResponse, err := parseGraphQLResponse(res)
	if err != nil {
		t.Errorf("Failed parsing graphql response: %v", err)
	}

	assert.Equal(t, len(errs), len(gqlResponse.Errors), "Error count mismatch")

	for i, e := range gqlResponse.Errors {
		assert.Equal(t, errs[i], e.Message, "Error mismatch")
	}
}

func AssertSuccess(t *testing.T, res *httptest.ResponseRecorder) {
	t.Helper()

	assert.Equal(t, http.StatusOK, res.Code, "Wrong status: ", res.Code)
}

func parseGraphQLResponse(res *httptest.ResponseRecorder) (graphql.Response, error) {
	var r graphql.Response

	err := json.Unmarshal(res.Body.Bytes(), &r)
	if err != nil {
		return r, err
	}

	return r, nil
}
