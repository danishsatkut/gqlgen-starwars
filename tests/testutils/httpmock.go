package testutils

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

const baseURL = "http://localhost:9999"

type MockRequest struct {
	StatusCode      int
	Method          string
	Path            string
	RequestHeaders  map[string]string
	ResponseHeaders map[string]string
}

func NewMockRequest(method string, path string, status int) *MockRequest {
	return &MockRequest{
		Method: method,
		Path: path,
		StatusCode: status,
	}
}

// RespondWith starts a local test server that will return the specified response.
// The caller should call Close when finished, to shut it down.
func (m *MockRequest) RespondWith(t *testing.T, jsonResponse interface{}) {
	t.Helper()

	httpmock.RegisterResponder(m.Method, m.URL(t).String(), func(req *http.Request) (*http.Response, error) {
		// Validate Request Headers
		if m.RequestHeaders != nil {
			for key, value := range m.RequestHeaders {
				assert.Equal(t, value, req.Header.Get(key))
			}
		}

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

func (m *MockRequest) URL(t *testing.T) *url.URL {
	u, err := url.Parse(baseURL)
	if err != nil {
		t.Errorf("Failed to parse base url: %v", err)
	}

	u.Path = m.Path

	return u
}
