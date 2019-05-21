package starwars

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/peterhellberg/swapi"

	"gqlgen-starwars/handlers"
	"gqlgen-starwars/tests/testutils"
)

func TestFilmQuery(t *testing.T) {
	t.Run("when film is returned from api", func(t *testing.T) {
		// Arrange
		m := testutils.NewMockRequest(http.StatusOK)
		m.RespondWith(t, swapi.Film{Title: "A New Hope"})
		defer m.Close()

		c := swapi.NewClient(nil)
		c.BaseURL = m.URL(t)

		query := `query { film(id: \"1\") { name } }`
		req := testutils.NewGraphQLRequest(t, query)

		response := httptest.NewRecorder()

		// Act
		testutils.PerformGraphQLRequest(response, req, handlers.SwapiClient(c))

		// Assert
		testutils.AssertSuccess(t, response)

		expected := `{"film":{"name":"A New Hope"}}`
		testutils.AssertGraphQLData(t, response, expected)
	})

	t.Run("when api returns error", func(t *testing.T) {
		// Arrange
		m := testutils.NewMockRequest(http.StatusInternalServerError)
		m.RespondWith(t, "")
		defer m.Close()

		c := swapi.NewClient(nil)
		c.BaseURL = m.URL(t)

		query := `query { film(id: \"1\") { name } }`
		req := testutils.NewGraphQLRequest(t, query)

		response := httptest.NewRecorder()

		// Act
		testutils.PerformGraphQLRequest(response, req, handlers.SwapiClient(c))

		// Assert
		testutils.AssertSuccess(t, response)
		testutils.AssertGraphQLData(t, response, "null")
		testutils.AssertGraphQLErrors(t, response, []string{"Failed to fetch film"})
	})
}
