package testutils

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockRequest struct {
	StatusCode      int
	Method          string
	Path            string
	RequestHeaders  map[string]string
	ResponseHeaders map[string]string
	server          *httptest.Server
}

func NewMockRequest(status int) *MockRequest {
	return &MockRequest{StatusCode: status}
}

// RespondWith starts a local test server that will return the specified response.
// The caller should call Close when finished, to shut it down.
func (m *MockRequest) RespondWith(t *testing.T, jsonResponse interface{}) {
	t.Helper()

	m.server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert := assert.New(t)

		// Validate Method
		if m.Method != "" {
			assert.Equal(m.Method, r.Method)
		}

		// Validate Path
		if m.Path != "" {
			assert.Equal(m.Path, r.URL.Path)
		}

		// Validate Request Headers
		if m.RequestHeaders != nil {
			for key, value := range m.RequestHeaders {
				assert.Equal(value, r.Header.Get(key))
			}
		}

		// Set Response Headers
		if m.ResponseHeaders != nil {
			for key, value := range m.ResponseHeaders {
				w.Header().Set(key, value)
			}
		}

		// Set Response Code
		w.WriteHeader(m.StatusCode)

		b, err := json.Marshal(jsonResponse)
		if err != nil {
			t.Fatal("Failed to marshal response", err, b)
		}

		w.Write(b)
	}))
}

func (m *MockRequest) URL(t *testing.T) *url.URL {
	u, err := url.Parse(m.server.URL)
	if err != nil {
		t.Fatal("Failed to parse url", err)
	}

	return u
}

func (m *MockRequest) Close() {
	m.server.Close()
}
