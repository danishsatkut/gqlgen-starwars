package middlewares

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/sirupsen/logrus"

	"gqlgen-starwars/utils"
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
			ctx = utils.WithLogger(ctx, logger)

			next.ServeHTTP(w, r.WithContext(ctx))
		}

		return http.HandlerFunc(fn)
	}

	return mwr
}
