package utils

import "github.com/sirupsen/logrus"

func DefaultLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)

	return logger
}
