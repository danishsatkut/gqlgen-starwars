package starwars

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/peterhellberg/swapi"

	"gqlgen-starwars/handlers"
	"gqlgen-starwars/tests/testutils"
)

func TestFilmQuery(t *testing.T) {
	t.Run("when film is returned from api", func(t *testing.T) {
		// Arrange
		httpmock.Activate()
		defer httpmock.Deactivate()

		m := testutils.NewMockRequest("GET", "/api/films/1", http.StatusOK)
		m.RespondWith(t, swapi.Film{Title: "Good Movie", URL: fmt.Sprintf("http://example.com/%v/", m.Path)})

		c := testutils.SwapiClient(m.URL())

		query := `query { film(id: \"1\") { id name } }`
		req := testutils.NewGraphQLRequest(t, query)

		response := httptest.NewRecorder()

		// Act
		testutils.PerformGraphQLRequest(response, req, handlers.SwapiClient(c))

		// Assert
		testutils.AssertSuccess(t, response)

		expected := `{"film":{"id":"1","name":"Good Movie"}}`
		testutils.AssertGraphQLData(t, response, expected)
		testutils.AssertGraphQLErrors(t, response, []string{})
	})

	t.Run("when api returns error", func(t *testing.T) {
		// Arrange
		httpmock.Activate()
		defer httpmock.Deactivate()

		m := testutils.NewMockRequest("GET", "/api/films/1", http.StatusInternalServerError)
		m.RespondWith(t, "")

		c := testutils.SwapiClient(m.URL())

		query := `query { film(id: \"1\") { name } }`
		req := testutils.NewGraphQLRequest(t, query)

		response := httptest.NewRecorder()

		// Act
		testutils.PerformGraphQLRequest(response, req, handlers.SwapiClient(c))

		// Assert
		testutils.AssertSuccess(t, response)
		testutils.AssertGraphQLData(t, response, "null")
		testutils.AssertGraphQLErrors(t, response, []string{"Something went wrong!"})
	})

	t.Run("when query contains both film and character", func(t *testing.T) {
		// Arrange
		httpmock.Activate()
		defer httpmock.Deactivate()

		m1 := testutils.NewMockRequest("GET", "/api/films/1", http.StatusOK)
		m1.RespondWith(t, swapi.Film{Title: "Good Movie", URL: fmt.Sprintf("http://example.com/%v/", m1.Path)})

		m2 := testutils.NewMockRequest("GET", "/api/people/2", http.StatusOK)
		m2.RespondWith(t, swapi.Person{Name: "John Smith"})

		c := testutils.SwapiClient(m1.URL())

		query := `query { film(id: \"1\") { name } character(id: \"2\") { name } }`
		req := testutils.NewGraphQLRequest(t, query)

		response := httptest.NewRecorder()

		// Act
		testutils.PerformGraphQLRequest(response, req, handlers.SwapiClient(c))

		// Assert
		testutils.AssertSuccess(t, response)

		expected := `{"film":{"name":"Good Movie"},"character":{"name":"John Smith"}}`
		testutils.AssertGraphQLData(t, response, expected)
		testutils.AssertGraphQLErrors(t, response, []string{})
	})
}
