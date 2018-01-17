package errors

import "strings"

//Separator is the string separator of error messages join
var Separator = " | "

//Error is the custom interface
type Error interface {
	error
	Messages() []string
}

//HTTPError is the http errors
type HTTPError interface {
	Error
	StatusCode() int
}

type errorWrapper struct {
	messages []string
}

func (e *errorWrapper) Error() string {
	return strings.Join(e.messages, Separator)
}

func (e *errorWrapper) Messages() []string {
	return e.messages
}

type httpError struct {
	statusCode int
	*errorWrapper
}

func (e *httpError) StatusCode() int {
	return e.statusCode
}

func new(defaultMessage string, messages ...string) *errorWrapper {
	err := errorWrapper{
		messages: messages,
	}

	if len(messages) == 0 {
		err.messages = append(err.messages, defaultMessage)
	} else {
		err.messages = messages
	}

	return &err
}
