package errors

import "net/http"

//NotFound represents a HTTP status code 404 error
type NotFound struct {
	httpError
}

//NewNotFound constructs a NotFound HTTP error
func NewNotFound(messages ...string) *NotFound {
	return &NotFound{
		httpError{
			http.StatusNotFound,
			new("Resource not found.", messages...),
		},
	}
}

//BadRequest represents a HTTP status code 400 error
type BadRequest struct {
	httpError
}

//NewBadRequest constructs a BadRequest HTTP error
func NewBadRequest(messages ...string) *BadRequest {
	return &BadRequest{
		httpError{
			http.StatusBadRequest,
			new("Invalid request sent.", messages...),
		},
	}
}

//InternalServer represents a HTTP status code 500 error
type InternalServer struct {
	httpError
}

//NewInternalServer constructs a InternalServer HTTP error
func NewInternalServer(messages ...string) *InternalServer {
	return &InternalServer{
		httpError{
			http.StatusInternalServerError,
			new("Something was wrong with server.", messages...),
		},
	}
}

//Unauthorized represents a HTTP status code 401 error
type Unauthorized struct {
	httpError
}

//NewUnauthorized constructs a Unauthorized HTTP error
func NewUnauthorized(messages ...string) *Unauthorized {
	return &Unauthorized{
		httpError{
			http.StatusUnauthorized,
			new("Unauthorized.", messages...),
		},
	}
}

//Forbidden represents a HTTP status code 403 error
type Forbidden struct {
	httpError
}

//NewForbidden constructs a Forbidden HTTP error
func NewForbidden(messages ...string) *Forbidden {
	return &Forbidden{
		httpError{
			http.StatusForbidden,
			new("Forbidden.", messages...),
		},
	}
}

//UnprocessableEntity represents a HTTP status code 422 error
type UnprocessableEntity struct {
	httpError
}

//NewUnprocessableEntity constructs a UnprocessableEntity HTTP error
func NewUnprocessableEntity(messages ...string) *UnprocessableEntity {
	return &UnprocessableEntity{
		httpError{
			http.StatusUnprocessableEntity,
			new("Unprocessable entity.", messages...),
		},
	}
}

//Unavailable represents a HTTP status code 503 error
type Unavailable struct {
	httpError
}

//NewUnavailable constructs a Unavailable HTTP error
func NewUnavailable(messages ...string) *Unavailable {
	return &Unavailable{
		httpError{
			http.StatusServiceUnavailable,
			new("Service is currently unavailable.", messages...),
		},
	}
}
