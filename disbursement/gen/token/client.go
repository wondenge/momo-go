// Code generated by goa v3.1.2, DO NOT EDIT.
//
// token client
//
// Command:
// $ goa gen github.com/wondenge/momo-go/disbursement/design

package token

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// Client is the "token" service client.
type Client struct {
	CreateEndpoint endpoint.Endpoint
}

// NewClient initializes a "token" service client given the endpoints.
func NewClient(create endpoint.Endpoint) *Client {
	return &Client{
		CreateEndpoint: create,
	}
}

// Create calls the "create" endpoint of the "token" service.
func (c *Client) Create(ctx context.Context, p *CreatePayload) (res *OAuthTokenError, err error) {
	var ires interface{}
	ires, err = c.CreateEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*OAuthTokenError), nil
}