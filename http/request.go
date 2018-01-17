package http

import (
	"github.com/levigross/grequests"
)

//Request exposes request request methods
type Request interface {
	SetHeaders(map[string]string)
	Headers() map[string]string
	SetParams(map[string]string)
	Params() map[string]string
	SetJSON(interface{})
	JSON() interface{}
}

//NewRequest constructs implementation of Request
func NewRequest() Request {
	return new(request)
}

//request is the struct wrapper of grequests struct
type request struct {
	grequests.RequestOptions
}

func (r *request) SetHeaders(headers map[string]string) {
	r.RequestOptions.Headers = headers
}

func (r *request) Headers() map[string]string {
	return r.RequestOptions.Headers
}

func (r *request) SetParams(params map[string]string) {
	r.RequestOptions.Params = params
}

func (r *request) Params() map[string]string {
	return r.RequestOptions.Params
}

func (r *request) SetJSON(data interface{}) {
	r.RequestOptions.JSON = data
}

func (r *request) JSON() interface{} {
	return r.RequestOptions.JSON
}
