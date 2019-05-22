package utils

import (
	"context"

	"github.com/sirupsen/logrus"
)

type contextKey struct {
	name string
}

var loggerKey = &contextKey{name: "logger"}

func WithLogger(ctx context.Context, logger *logrus.Entry) context.Context {
	return context.WithValue(ctx, loggerKey, logger)
}

func GetLogger(ctx context.Context) *logrus.Entry {
	if logger, ok := ctx.Value(loggerKey).(*logrus.Entry); ok {
		return logger
	}

	return logrus.NewEntry(logrus.New())
}
