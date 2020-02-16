package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"gqlgen-starwars/tests/testutils"
)

func TestComplexity(t *testing.T) {
	t.Run("when query is too complex", func(t *testing.T) {
		query := `query { film(id: \"1\") { characters { id films { id characters { id } } } } }`
		req := testutils.NewGraphQLRequest(t, query)

		response := httptest.NewRecorder()

		// Act
		testutils.PerformGraphQLRequest(response, req)

		// Assert
		testutils.AssertStatus(t, response, http.StatusUnprocessableEntity)
		testutils.AssertGraphQLData(t, response, "null")
		testutils.AssertGraphQLErrors(t, response, []string{"operation has complexity 178, which exceeds the limit of 150"})
	})
}
