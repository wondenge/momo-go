// Code generated by goa v3.1.2, DO NOT EDIT.
//
// validate client
//
// Command:
// $ goa gen github.com/wondenge/momo-go/disbursement/design

package validate

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// Client is the "validate" service client.
type Client struct {
	ShowEndpoint endpoint.Endpoint
}

// NewClient initializes a "validate" service client given the endpoints.
func NewClient(show endpoint.Endpoint) *Client {
	return &Client{
		ShowEndpoint: show,
	}
}

// Show calls the "show" endpoint of the "validate" service.
func (c *Client) Show(ctx context.Context, p *ShowPayload) (res *Errorreason, err error) {
	var ires interface{}
	ires, err = c.ShowEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*Errorreason), nil
}
