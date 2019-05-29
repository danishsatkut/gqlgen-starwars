package utils

import "github.com/sirupsen/logrus"

var DefaultLogger = newLogger()

func newLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)

	return logger
}
