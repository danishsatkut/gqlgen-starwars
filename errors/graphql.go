package errors

type GraphQLError struct {
	message    string
	extensions map[string]interface{}
}

func NewGraphQLError(msg string, extensions map[string]interface{}) GraphQLError {
	return GraphQLError{msg, extensions}
}

func (e GraphQLError) Extensions() map[string]interface{} {
	return e.extensions
}

func (e GraphQLError) Error() string {
	return e.message
}
