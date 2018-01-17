package http

import (
	"reflect"

	"github.com/levigross/grequests"
	"github.com/stone-payments/goclienttools/errors"
)

//Requestable is the interface that handles with http requests
type Requestable interface {
	Get(url string, request Request) (Response, error)
	Post(url string, request Request) (Response, error)
	Put(url string, request Request) (Response, error)
	Delete(url string, request Request) (Response, error)
}

type requester struct{}

//Get is the method wrapper of grequests
func (r requester) Get(url string, request Request) (Response, error) {
	opt, err := r.requestOptions(request)

	if err != nil {
		return nil, err
	}

	resp, err := grequests.Get(url, opt)
	return &response{resp}, err
}

//Post is the method wrapper of grequests
func (r requester) Post(url string, request Request) (Response, error) {
	opt, err := r.requestOptions(request)

	if err != nil {
		return nil, err
	}

	resp, err := grequests.Post(url, opt)
	return &response{resp}, err
}

//Put is the method wrapper of grequests
func (r requester) Put(url string, request Request) (Response, error) {
	opt, err := r.requestOptions(request)

	if err != nil {
		return nil, err
	}

	resp, err := grequests.Put(url, opt)
	return &response{resp}, err
}

//Delete is the method wrapper of grequests
func (r requester) Delete(url string, request Request) (Response, error) {
	opt, err := r.requestOptions(request)

	if err != nil {
		return nil, err
	}

	resp, err := grequests.Delete(url, opt)
	return &response{resp}, err
}

func (r requester) requestOptions(req Request) (*grequests.RequestOptions, error) {
	if r, ok := req.(*request); ok {
		return &r.RequestOptions, nil
	}

	return nil, errors.NewInvalidType(reflect.TypeOf(req), reflect.TypeOf(new(request)))
}
