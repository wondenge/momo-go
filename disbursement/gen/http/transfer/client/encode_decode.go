// Code generated by goa v3.1.2, DO NOT EDIT.
//
// transfer HTTP client encoders and decoders
//
// Command:
// $ goa gen github.com/wondenge/momo-go/disbursement/design

package client

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"

	transfer "github.com/wondenge/momo-go/disbursement/gen/transfer"
	transferviews "github.com/wondenge/momo-go/disbursement/gen/transfer/views"
	goahttp "goa.design/goa/v3/http"
)

// BuildCreateRequest instantiates a HTTP request object with method and path
// set to call the "transfer" service "create" endpoint
func (c *Client) BuildCreateRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: CreateTransferPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("transfer", "create", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeCreateRequest returns an encoder for requests sent to the transfer
// create server.
func EncodeCreateRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*transfer.CreatePayload)
		if !ok {
			return goahttp.ErrInvalidType("transfer", "create", "*transfer.CreatePayload", v)
		}
		if p.Authorization != nil {
			head := *p.Authorization
			req.Header.Set(" Authorization", head)
		}
		if p.XCallbackURL != nil {
			head := *p.XCallbackURL
			req.Header.Set(" X-Callback-Url", head)
		}
		{
			head := p.XReferenceID
			req.Header.Set(" X-Reference-Id", head)
		}
		{
			head := p.XTargetEnvironment
			req.Header.Set(" X-Target-Environment", head)
		}
		{
			head := p.OcpApimSubscriptionKey
			req.Header.Set(" Ocp-Apim-Subscription-Key", head)
		}
		return nil
	}
}

// DecodeCreateResponse returns a decoder for responses returned by the
// transfer create endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeCreateResponse may return the following errors:
//	- "bad_request" (type *goa.ServiceError): http.StatusBadRequest
//	- error: internal error
func DecodeCreateResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusAccepted:
			var (
				body CreateResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("transfer", "create", err)
			}
			p := NewCreateErrorreasonAccepted(&body)
			view := "default"
			vres := &transferviews.Errorreason{Projected: p, View: view}
			if err = transferviews.ValidateErrorreason(vres); err != nil {
				return nil, goahttp.ErrValidationError("transfer", "create", err)
			}
			res := transfer.NewErrorreason(vres)
			return res, nil
		case http.StatusBadRequest:
			var (
				body CreateBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("transfer", "create", err)
			}
			err = ValidateCreateBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("transfer", "create", err)
			}
			return nil, NewCreateBadRequest(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("transfer", "create", resp.StatusCode, string(body))
		}
	}
}

// BuildGetRequest instantiates a HTTP request object with method and path set
// to call the "transfer" service "get" endpoint
func (c *Client) BuildGetRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		referenceID string
	)
	{
		p, ok := v.(*transfer.GetPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("transfer", "get", "*transfer.GetPayload", v)
		}
		if p.ReferenceID != nil {
			referenceID = *p.ReferenceID
		}
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: GetTransferPath(referenceID)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("transfer", "get", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeGetRequest returns an encoder for requests sent to the transfer get
// server.
func EncodeGetRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*transfer.GetPayload)
		if !ok {
			return goahttp.ErrInvalidType("transfer", "get", "*transfer.GetPayload", v)
		}
		if p.Authorization != nil {
			head := *p.Authorization
			req.Header.Set(" Authorization", head)
		}
		{
			head := p.XTargetEnvironment
			req.Header.Set(" X-Target-Environment", head)
		}
		{
			head := p.OcpApimSubscriptionKey
			req.Header.Set(" Ocp-Apim-Subscription-Key", head)
		}
		return nil
	}
}

// DecodeGetResponse returns a decoder for responses returned by the transfer
// get endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeGetResponse may return the following errors:
//	- "bad_request" (type *goa.ServiceError): http.StatusBadRequest
//	- error: internal error
func DecodeGetResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body GetResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("transfer", "get", err)
			}
			p := NewGetErrorreasonOK(&body)
			view := "default"
			vres := &transferviews.Errorreason{Projected: p, View: view}
			if err = transferviews.ValidateErrorreason(vres); err != nil {
				return nil, goahttp.ErrValidationError("transfer", "get", err)
			}
			res := transfer.NewErrorreason(vres)
			return res, nil
		case http.StatusBadRequest:
			var (
				body GetBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("transfer", "get", err)
			}
			err = ValidateGetBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("transfer", "get", err)
			}
			return nil, NewGetBadRequest(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("transfer", "get", resp.StatusCode, string(body))
		}
	}
}