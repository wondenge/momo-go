// Code generated by goa v3.1.3, DO NOT EDIT.
//
// user client
//
// Command:
// $ goa gen github.com/wondenge/momo-go/design

package user

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// Client is the "user" service client.
type Client struct {
	NewUserEndpoint        endpoint.Endpoint
	NewKeyEndpoint         endpoint.Endpoint
	GetUserEndpoint        endpoint.Endpoint
	GetUserDetailsEndpoint endpoint.Endpoint
}

// NewClient initializes a "user" service client given the endpoints.
func NewClient(newUser, newKey, getUser, getUserDetails endpoint.Endpoint) *Client {
	return &Client{
		NewUserEndpoint:        newUser,
		NewKeyEndpoint:         newKey,
		GetUserEndpoint:        getUser,
		GetUserDetailsEndpoint: getUserDetails,
	}
}

// NewUser calls the "NewUser" endpoint of the "user" service.
// NewUser may return the following errors:
//	- "bad_request" (type *goa.ServiceError): Bad request, e.g. invalid data was sent in the request.
//	- "conflict" (type *ErrorReason): Conflict, duplicated reference id
//	- "internal_error" (type *goa.ServiceError): Internal error. Check log for information.
//	- error: internal error
func (c *Client) NewUser(ctx context.Context, p *APIUser) (res string, err error) {
	var ires interface{}
	ires, err = c.NewUserEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(string), nil
}

// NewKey calls the "NewKey" endpoint of the "user" service.
// NewKey may return the following errors:
//	- "bad_request" (type *goa.ServiceError): Bad request, e.g. invalid data was sent in the request.
//	- "not_found" (type *ErrorReason): Not found, reference id not found or closed in sandbox
//	- "internal_error" (type *goa.ServiceError): Internal error. Check log for information.
//	- error: internal error
func (c *Client) NewKey(ctx context.Context, p *NewKeyPayload) (res *APIUserKeyResult, err error) {
	var ires interface{}
	ires, err = c.NewKeyEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*APIUserKeyResult), nil
}

// GetUser calls the "GetUser" endpoint of the "user" service.
// GetUser may return the following errors:
//	- "bad_request" (type *goa.ServiceError): Bad request, e.g. invalid data was sent in the request.
//	- "not_found" (type *goa.ServiceError): Not found, reference id not found or closed in sandbox
//	- "internal_error" (type *goa.ServiceError): Internal error. Check log for information.
//	- error: internal error
func (c *Client) GetUser(ctx context.Context, p *GetUserPayload) (res string, err error) {
	var ires interface{}
	ires, err = c.GetUserEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(string), nil
}

// GetUserDetails calls the "GetUserDetails" endpoint of the "user" service.
// GetUserDetails may return the following errors:
//	- "bad_request" (type *goa.ServiceError): Bad request, e.g. invalid data was sent in the request.
//	- "internal_error" (type *goa.ServiceError): Internal error. Check log for information.
//	- error: internal error
func (c *Client) GetUserDetails(ctx context.Context, p string) (res *APIUserResult, err error) {
	var ires interface{}
	ires, err = c.GetUserDetailsEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*APIUserResult), nil
}