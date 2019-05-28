package errors

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

const (
	UserInputError = "BAD_USER_INPUT"
	ServerError    = "INTERNAL_SERVER_ERROR"
	NotFoundError  = "NOT_FOUND"
	APIError       = "API_ERROR"
)

func NewUserInputError(message string, argument string) error {
	ext := map[string]interface{}{
		"code":  UserInputError,
		"input": argument,
	}

	return NewGraphQLError(message, ext)
}

func NewServerError(err interface{}, stack string) error {
	ext := map[string]interface{}{
		"code":       ServerError,
		"error":      fmt.Sprint(err),
		"stacktrace": strings.Split(stack, "\n"),
	}

	return NewGraphQLError("Internal Server Error", ext)
}

func NewResourceNotFoundError(message string, resourceName string, resourceId string) error {
	ext := map[string]interface{}{
		"code": NotFoundError,
		"resource": map[string]string{
			"id":   resourceId,
			"name": resourceName,
		},
	}

	return NewGraphQLError(message, ext)
}

func NewAPIError(err error) error {
	ext := map[string]interface{}{
		"code":  APIError,
		"error": err.Error(),
	}

	return NewGraphQLError("Something went wrong!", ext)
}

func New(message string) error {
	return errors.New(message)
}

func Wrapf(err error, message string, args ...interface{}) error {
	return errors.Wrapf(err, message, args)
}
