package http

import (
	"fmt"

	"github.com/crowleyfelix/base-client-go/errors"
)

//Manager exposes request manager methods
type Manager interface {
	BuildURL(endpoint string, params ...interface{}) string
	Request(method RequestMethod, url string, request Request, interceptors ...Interceptor) (Response, errors.Error)
}

//NewManager constructs a request manager
func NewManager(baseURL string, interceptors ...Interceptor) Manager {
	return &manager{baseURL, interceptors}
}

type manager struct {
	baseURL      string
	interceptors []Interceptor
}

func (s *manager) BuildURL(endpoint string, params ...interface{}) string {
	return s.baseURL + fmt.Sprintf(endpoint, params...)
}

func (s *manager) Request(method RequestMethod, url string, request Request, interceptors ...Interceptor) (Response, errors.Error) {
	interceptors = s.allInterceptors(interceptors...)

	for _, interceptor := range interceptors {
		request = interceptor.OnRequest(request)
	}

	resp, e := method(url, request)
	if e != nil {
		return nil, errors.NewCallout(e.Error())

	}

	var err errors.Error
	for _, interceptor := range interceptors {
		resp, err = interceptor.OnResponse(resp)

		if err != nil {
			break
		}
	}

	return resp, err
}

func (s *manager) allInterceptors(interceptors ...Interceptor) []Interceptor {
	var all []Interceptor
	all = append(all, Interceptors...)
	all = append(all, s.interceptors...)
	all = append(all, interceptors...)
	return all
}
