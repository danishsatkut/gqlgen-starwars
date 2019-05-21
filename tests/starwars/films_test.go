package starwars

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/peterhellberg/swapi"

	"gqlgen-starwars/handlers"
	"gqlgen-starwars/tests/testutils"
)

func TestFilmQuery(t *testing.T) {
	t.Run("when film is returned from api", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			film := swapi.Film{Title: "A New Hope"}
			b, err := json.Marshal(film)
			if err != nil {
				t.Fatal("Failed to marshal film", err)
			}

			w.Write(b)
		}))
		defer server.Close()

		u, err := url.Parse(server.URL)
		if err != nil {
			t.Fatal(err.Error())
		}

		c := swapi.NewClient(nil)
		c.BaseURL = u

		// Arrange
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
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(""))
		}))
		defer server.Close()

		u, err := url.Parse(server.URL)
		if err != nil {
			t.Fatal(err.Error())
		}

		c := swapi.NewClient(nil)
		c.BaseURL = u

		// Arrange
		query := `query { film(id: \"1\") { name } }`
		req := testutils.NewGraphQLRequest(t, query)

		response := httptest.NewRecorder()

		// Act
		testutils.PerformGraphQLRequest(response, req, handlers.SwapiClient(c))

		// Assert
		testutils.AssertSuccess(t, response)
		testutils.AssertGraphQLData(t, response, "null")
		testutils.AssertGraphQLErrors(t, response, []string{"failed to fetch film: error reading response from GET /api/films/1?format=json: EOF"})
	})
}
