package utils

import (
	"strconv"
	"strings"
)

// Id returns resource id from url
// eg: https://swapi.co/api/films/1/
func Id(url string) (int, error) {
	s := strings.Split(url, "/")

	id, err := strconv.Atoi(s[len(s) - 2])
	if err != nil {
		return 0, err
	}

	return id, err
}
