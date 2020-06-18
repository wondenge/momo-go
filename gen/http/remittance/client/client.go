// Code generated by goa v3.1.3, DO NOT EDIT.
//
// remittance client HTTP transport
//
// Command:
// $ goa gen github.com/wondenge/momo-go/design

package client

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	goahttp "goa.design/goa/v3/http"
)

// Client lists the remittance service endpoint HTTP clients.
type Client struct {
	// NewToken Doer is the HTTP client used to make requests to the NewToken
	// endpoint.
	NewTokenDoer goahttp.Doer

	// GetBalance Doer is the HTTP client used to make requests to the GetBalance
	// endpoint.
	GetBalanceDoer goahttp.Doer

	// RetrieveAccountHolder Doer is the HTTP client used to make requests to the
	// RetrieveAccountHolder endpoint.
	RetrieveAccountHolderDoer goahttp.Doer

	// Transfer Doer is the HTTP client used to make requests to the Transfer
	// endpoint.
	TransferDoer goahttp.Doer

	// TransferStatus Doer is the HTTP client used to make requests to the
	// TransferStatus endpoint.
	TransferStatusDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the remittance service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		NewTokenDoer:              doer,
		GetBalanceDoer:            doer,
		RetrieveAccountHolderDoer: doer,
		TransferDoer:              doer,
		TransferStatusDoer:        doer,
		RestoreResponseBody:       restoreBody,
		scheme:                    scheme,
		host:                      host,
		decoder:                   dec,
		encoder:                   enc,
	}
}

// NewToken returns an endpoint that makes HTTP requests to the remittance
// service NewToken server.
func (c *Client) NewToken() endpoint.Endpoint {
	var (
		encodeRequest  = EncodeNewTokenRequest(c.encoder)
		decodeResponse = DecodeNewTokenResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildNewTokenRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.NewTokenDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("remittance", "NewToken", err)
		}
		return decodeResponse(resp)
	}
}

// GetBalance returns an endpoint that makes HTTP requests to the remittance
// service GetBalance server.
func (c *Client) GetBalance() endpoint.Endpoint {
	var (
		encodeRequest  = EncodeGetBalanceRequest(c.encoder)
		decodeResponse = DecodeGetBalanceResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildGetBalanceRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.GetBalanceDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("remittance", "GetBalance", err)
		}
		return decodeResponse(resp)
	}
}

// RetrieveAccountHolder returns an endpoint that makes HTTP requests to the
// remittance service RetrieveAccountHolder server.
func (c *Client) RetrieveAccountHolder() endpoint.Endpoint {
	var (
		decodeResponse = DecodeRetrieveAccountHolderResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildRetrieveAccountHolderRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.RetrieveAccountHolderDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("remittance", "RetrieveAccountHolder", err)
		}
		return decodeResponse(resp)
	}
}

// Transfer returns an endpoint that makes HTTP requests to the remittance
// service Transfer server.
func (c *Client) Transfer() endpoint.Endpoint {
	var (
		encodeRequest  = EncodeTransferRequest(c.encoder)
		decodeResponse = DecodeTransferResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildTransferRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.TransferDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("remittance", "Transfer", err)
		}
		return decodeResponse(resp)
	}
}

// TransferStatus returns an endpoint that makes HTTP requests to the
// remittance service TransferStatus server.
func (c *Client) TransferStatus() endpoint.Endpoint {
	var (
		decodeResponse = DecodeTransferStatusResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildTransferStatusRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.TransferStatusDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("remittance", "TransferStatus", err)
		}
		return decodeResponse(resp)
	}
}