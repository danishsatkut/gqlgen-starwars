package testutils

import (
	"net/url"

	"github.com/peterhellberg/swapi"
)

func SwapiClient(u *url.URL) *swapi.Client {
	c := swapi.NewClient(nil)
	c.BaseURL = &url.URL{
		Scheme: u.Scheme,
		Host: u.Host,
	}

	return c
}
