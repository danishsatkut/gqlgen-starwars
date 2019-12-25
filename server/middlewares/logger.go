package middlewares

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/sirupsen/logrus"
)

func SetDefaultLogger(logger *logrus.Logger) {
	middleware.DefaultLogger = middleware.RequestLogger(&middleware.DefaultLogFormatter{
		Logger:  logger,
		NoColor: false,
	})
}

func LogEntry(l *logrus.Logger) func(http.Handler) http.Handler {
	mwr := func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			logger := l.WithField("request_id", middleware.GetReqID(ctx))
			ctx = WithLogEntry(ctx, logger)

			next.ServeHTTP(w, r.WithContext(ctx))
		}

		return http.HandlerFunc(fn)
	}

	return mwr
}

var loggerKey = contextKey("logger")

func WithLogEntry(ctx context.Context, entry *logrus.Entry) context.Context {
	return context.WithValue(ctx, loggerKey, entry)
}

func GetLogEntry(ctx context.Context) *logrus.Entry {
	if entry, ok := ctx.Value(loggerKey).(*logrus.Entry); ok {
		return entry
	}

	return logrus.NewEntry(logrus.New())
}
