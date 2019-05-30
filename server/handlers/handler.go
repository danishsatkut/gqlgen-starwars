package handlers

import (
	"github.com/peterhellberg/swapi"
	"github.com/sirupsen/logrus"
)

type config struct {
	swapiClient *swapi.Client
	logger      *logrus.Logger
}

func (c *config) update(options ...Option) {
	for _, option := range options {
		option(c)
	}
}

type Option func(cfg *config)

func SwapiClient(c *swapi.Client) Option {
	return func(cfg *config) {
		cfg.swapiClient = c
	}
}

func Logger(l *logrus.Logger) Option {
	return func(cfg *config) {
		cfg.logger = l
	}
}
