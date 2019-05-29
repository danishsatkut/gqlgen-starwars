package utils

import (
	"context"

	"github.com/sirupsen/logrus"
)

type contextKey string

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
