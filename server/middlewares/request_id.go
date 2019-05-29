package middlewares

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/google/uuid"
)

const (
	requestIdHeader = "X-Request-Id"
)

func RequestID(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		requestID := r.Header.Get(requestIdHeader)
		if requestID == "" {
			requestID = uuid.New().String()
		}

		ctx = context.WithValue(ctx, middleware.RequestIDKey, requestID)
		next.ServeHTTP(w, r.WithContext(ctx))
	}

	return http.HandlerFunc(fn)
}
