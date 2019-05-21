package starwars

import (
	"net/http/httptest"
	"testing"

	"gqlgen-starwars/tests/testutils"
)

func TestFilmQuery(t *testing.T) {
	t.Run("when film is returned from api", func(t *testing.T) {
		// Arrange
		query := `query { film(id: \"1\") { name } }`
		req := testutils.NewGraphQLRequest(t, query)

		response := httptest.NewRecorder()

		// Act
		testutils.PerformGraphQLRequest(response, req)

		// Assert
		testutils.AssertSuccess(t, response)

		expected := `{"film":{"name":"A New Hope"}}`
		testutils.AssertGraphQLData(t, response, expected)
	})

	t.Run("when api returns error", func(t *testing.T) {
		t.Skip("not implemented")
	})
}
