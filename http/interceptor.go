package http

import "github.com/stone-payments/goclienttools/errors"

//Interceptor exposes interceptor methods
type Interceptor interface {
	OnRequest(RequestOptions) RequestOptions
	OnResponse(Response) (Response, errors.Error)
}

type interceptor struct {
	onRequestCallback  func(RequestOptions) RequestOptions
	onResponseCallback func(Response) (Response, errors.Error)
}

//OnRequestCallback define request callback interceptor assignature
type OnRequestCallback func(RequestOptions) RequestOptions

//OnResponseCallback define response callback interceptor assignature
type OnResponseCallback func(Response) (Response, errors.Error)

//NewInterceptor constructs interceptor
func NewInterceptor(onRequestCallback OnRequestCallback, onResponseCallback OnResponseCallback) Interceptor {
	return &interceptor{
		onRequestCallback,
		onResponseCallback,
	}
}

func (p *interceptor) OnRequest(request RequestOptions) RequestOptions {
	if p.onRequestCallback != nil {
		return p.onRequestCallback(request)
	}

	return request
}

func (p *interceptor) OnResponse(response Response) (Response, errors.Error) {
	if p.onResponseCallback != nil {
		return p.onResponseCallback(response)
	}

	return response, nil
}
