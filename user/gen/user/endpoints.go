// Code generated by goa v3.1.2, DO NOT EDIT.
//
// user endpoints
//
// Command:
// $ goa gen github.com/wondenge/momo-go/user/design

package user

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints wraps the "user" service endpoints.
type Endpoints struct {
	Createuser endpoint.Endpoint
	Createkey  endpoint.Endpoint
	List       endpoint.Endpoint
	Show       endpoint.Endpoint
}

// NewEndpoints wraps the methods of the "user" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		Createuser: NewCreateuserEndpoint(s),
		Createkey:  NewCreatekeyEndpoint(s),
		List:       NewListEndpoint(s),
		Show:       NewShowEndpoint(s),
	}
}

// Use applies the given middleware to all the "user" service endpoints.
func (e *Endpoints) Use(m func(endpoint.Endpoint) endpoint.Endpoint) {
	e.Createuser = m(e.Createuser)
	e.Createkey = m(e.Createkey)
	e.List = m(e.List)
	e.Show = m(e.Show)
}

// NewCreateuserEndpoint returns an endpoint function that calls the method
// "createuser" of service "user".
func NewCreateuserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*CreateuserPayload)
		res, err := s.Createuser(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedErrorreason(res, "default")
		return vres, nil
	}
}

// NewCreatekeyEndpoint returns an endpoint function that calls the method
// "createkey" of service "user".
func NewCreatekeyEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*CreatekeyPayload)
		res, err := s.Createkey(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedErrorreason(res, "default")
		return vres, nil
	}
}

// NewListEndpoint returns an endpoint function that calls the method "list" of
// service "user".
func NewListEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*ListPayload)
		res, err := s.List(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedErrorreason(res, "default")
		return vres, nil
	}
}

// NewShowEndpoint returns an endpoint function that calls the method "show" of
// service "user".
func NewShowEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*ShowPayload)
		res, err := s.Show(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedErrorreason(res, "default")
		return vres, nil
	}
}