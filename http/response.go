package http

import (
	"net/http"

	"github.com/crowleyfelix/base-client-go/errors"
	"github.com/levigross/grequests"
)

//Response exposes http response methods
type Response interface {
	Bytes() []byte
	Ok() bool
	Error() error
	RawResponse() *http.Response
	StatusCode() int
	Header() http.Header
	JSON(interface{}) errors.Error
	ClearInternalBuffer()
	Close() error
	Read([]byte) (int, error)
}

//response is the struct wrapper of grequests struct
type response struct {
	*grequests.Response
}

func (r *response) Ok() bool {
	return r.Response.Ok
}
func (r *response) Error() error {
	return r.Response.Error
}
func (r *response) RawResponse() *http.Response {
	return r.Response.RawResponse
}
func (r *response) StatusCode() int {
	return r.Response.StatusCode
}
func (r *response) Header() http.Header {
	return r.Response.Header
}
func (r *response) JSON(obj interface{}) errors.Error {
	err := r.Response.JSON(obj)
	if err != nil {
		return errors.NewSerializing(err.Error())
	}
	return nil
}
