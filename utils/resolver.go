package utils

import (
	"fmt"
	"strconv"

	"gqlgen-starwars/errors"
)

func ParseId(id string) (int, error) {
	resourceId, err := strconv.Atoi(id)
	if err != nil {
		return 0, errors.NewUserInputError(fmt.Sprintf("Invalid id: %v", id), "id")
	}

	return resourceId, nil
}
