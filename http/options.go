package http

import (
	"github.com/levigross/grequests"
)

//Options exposes request options methods
type Options interface {
	SetHeaders(map[string]string)
	Headers() map[string]string
	SetParams(map[string]string)
	Params() map[string]string
	SetJSON(interface{})
	JSON() interface{}
}

//NewOptions constructs implementation of Options
func NewOptions() Options {
	return new(options)
}

//options is the struct wrapper of grequests struct
type options struct {
	grequests.RequestOptions
}

func (r *options) SetHeaders(headers map[string]string) {
	r.RequestOptions.Headers = headers
}

func (r *options) Headers() map[string]string {
	return r.RequestOptions.Headers
}

func (r *options) SetParams(params map[string]string) {
	r.RequestOptions.Params = params
}

func (r *options) Params() map[string]string {
	return r.RequestOptions.Params
}

func (r *options) SetJSON(data interface{}) {
	r.RequestOptions.JSON = data
}

func (r *options) JSON() interface{} {
	return r.RequestOptions.JSON
}
