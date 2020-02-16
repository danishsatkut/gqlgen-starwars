package testutils

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/99designs/gqlgen/graphql"
	"github.com/stretchr/testify/assert"
)

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

	if len(errs) != len(gqlResponse.Errors) {
		t.Fatal("Error count mismatch: ", gqlResponse.Errors)
	}

	for i, e := range gqlResponse.Errors {
		assert.Equal(t, errs[i], e.Message, "Error mismatch")
	}
}

func AssertStatus(t *testing.T, res *httptest.ResponseRecorder, code int) {
	t.Helper()

	assert.Equal(t, code, res.Code, "Wrong status")
}

func AssertSuccess(t *testing.T, res *httptest.ResponseRecorder) {
	t.Helper()

	AssertStatus(t, res, http.StatusOK)
}

func parseGraphQLResponse(res *httptest.ResponseRecorder) (graphql.Response, error) {
	var r graphql.Response

	err := json.Unmarshal(res.Body.Bytes(), &r)
	if err != nil {
		return r, err
	}

	return r, nil
}
