package testutils

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

var baseURL = &url.URL{Scheme: "http", Host: "localhost:9999"}

type MockRequest struct {
	StatusCode      int
	Method          string
	Path            string
	RequestHeaders  map[string]string
	ResponseHeaders map[string]string
}

func NewMockRequest(method string, path string, status int) *MockRequest {
	return &MockRequest{
		Method:     method,
		Path:       path,
		StatusCode: status,
	}
}

// RespondWith registers a responder for the mock request path that will return the specified response.
func (m *MockRequest) RespondWith(t *testing.T, jsonResponse interface{}) {
	t.Helper()

	httpmock.RegisterResponder(m.Method, m.URL().String(), func(req *http.Request) (*http.Response, error) {
		// Assert Request Headers
		if m.RequestHeaders != nil {
			for key, value := range m.RequestHeaders {
				assert.Equal(t, value, req.Header.Get(key))
			}
		}

		// Create json response
		res, err := httpmock.NewJsonResponse(m.StatusCode, jsonResponse)
		if err != nil {
			t.Fatal("Failed to marshal response", err)
		}

		// Set Response Headers
		if m.ResponseHeaders != nil {
			for key, value := range m.ResponseHeaders {
				res.Header.Set(key, value)
			}
		}

		return res, nil
	})
}

func (m *MockRequest) URL() *url.URL {
	return baseURL.ResolveReference(&url.URL{Path: m.Path})
}

func (m *MockRequest) ResourceURL() string {
	// "http://example.com/%v/"

	return fmt.Sprintf("%v/%v/", baseURL.String(), m.Path)
}
