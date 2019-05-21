package starwars

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"gqlgen-starwars/handlers"
)

func TestFilmQuery(t *testing.T) {
	t.Run("when film is returned from api", func(t *testing.T) {
		query := `query { film(id: \"1\") { name } }`
		gqlQuery := fmt.Sprintf("{\"query\":\"%v\"}", query)

		//	Make request
		req, err := http.NewRequest("POST", "/query", strings.NewReader(gqlQuery))
		if err != nil {
			t.Fatal(err.Error())
		}

		req.Header.Set("Content-Type", "application/json")

		response := httptest.NewRecorder()

		// Perform
		handler := handlers.NewGraphQlHandler()

		handler.ServeHTTP(response, req)

		if response.Code != http.StatusOK {
			t.Errorf("Wrong status: %v", response.Code)
		}

		resBody := response.Body.String()

		expected := `{"data":{"film":{"name":"A New Hope"}}}`
		if resBody != expected {
			t.Errorf("Unexpected response: %v", resBody)
		}
	})

	t.Run("when api returns error", func(t *testing.T) {
		t.Skip("not implemented")
	})
}
