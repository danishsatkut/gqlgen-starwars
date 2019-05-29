package middlewares

import (
	"github.com/go-chi/chi/middleware"
	"github.com/sirupsen/logrus"
)

func SetDefaultLogger(logger *logrus.Logger) {
	middleware.DefaultLogger = middleware.RequestLogger(&middleware.DefaultLogFormatter{
		Logger:  logger,
		NoColor: false,
	})
}
