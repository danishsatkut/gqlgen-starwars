package middlewares

import (
	"context"
	"net/http"
)

type contextKey string

const authTokenKey = "auth_token"

func Auth(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		v := r.Header.Get("Authorization")

		ctx := r.Context()
		ctx = context.WithValue(ctx, authTokenKey, v)

		next.ServeHTTP(w, r.WithContext(ctx))
	}

	return http.HandlerFunc(fn)
}

func GetAuthToken(ctx context.Context) string {
	if ctx == nil {
		return ""
	}

	if token, ok := ctx.Value(authTokenKey).(string); ok {
		return token
	}

	return ""
}
