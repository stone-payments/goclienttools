package errors

import (
	"net/http"
)

//Build returns an error based on http status code
func Build(statusCode int, messages ...string) HTTPError {
	switch statusCode {
	default:
		return NewInternalServer(messages...)
	case http.StatusBadRequest:
		return NewBadRequest(messages...)
	case http.StatusUnauthorized:
		return NewUnauthorized(messages...)
	case http.StatusForbidden:
		return NewForbidden(messages...)
	case http.StatusNotFound:
		return NewNotFound(messages...)
	case http.StatusUnprocessableEntity:
		return NewUnprocessableEntity(messages...)
	case http.StatusServiceUnavailable:
		return NewUnavailable(messages...)
	}
}
