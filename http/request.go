package http

import (
	"github.com/levigross/grequests"
)

//RequestOptions exposes request request methods
type RequestOptions interface {
	SetHeaders(map[string]string)
	Headers() map[string]string
	SetParams(map[string]string)
	Params() map[string]string
	SetJSON(interface{})
	JSON() interface{}
}

//NewRequestOptions constructs implementation of RequestOptions
func NewRequestOptions() RequestOptions {
	return new(requestOptions)
}

//requestOptions is the struct wrapper of grequests struct
type requestOptions struct {
	grequests.RequestOptions
}

func (r *requestOptions) SetHeaders(headers map[string]string) {
	r.RequestOptions.Headers = headers
}

func (r *requestOptions) Headers() map[string]string {
	return r.RequestOptions.Headers
}

func (r *requestOptions) SetParams(params map[string]string) {
	r.RequestOptions.Params = params
}

func (r *requestOptions) Params() map[string]string {
	return r.RequestOptions.Params
}

func (r *requestOptions) SetJSON(data interface{}) {
	r.RequestOptions.JSON = data
}

func (r *requestOptions) JSON() interface{} {
	return r.RequestOptions.JSON
}
