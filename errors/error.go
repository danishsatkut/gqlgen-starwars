package errors

import (
	"github.com/pkg/errors"
)

const (
	UserInputError = "BAD_USER_INPUT"
)

func NewUserInputError(message string, argument string) error {
	ext := map[string]interface{}{
		"code": UserInputError,
		"input": argument,
	}

	return NewGraphQLError(message, ext)
}

func New(message string) error {
	return errors.New(message)
}

func Wrapf(err error, message string, args ...interface{}) error {
	return errors.Wrapf(err, message)
}
