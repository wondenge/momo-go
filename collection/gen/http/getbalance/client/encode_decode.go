// Code generated by goa v3.1.2, DO NOT EDIT.
//
// getbalance HTTP client encoders and decoders
//
// Command:
// $ goa gen github.com/wondenge/momo-go/collection/design

package client

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"

	getbalance "github.com/wondenge/momo-go/collection/gen/getbalance"
	getbalanceviews "github.com/wondenge/momo-go/collection/gen/getbalance/views"
	goahttp "goa.design/goa/v3/http"
)

// BuildShowRequest instantiates a HTTP request object with method and path set
// to call the "getbalance" service "show" endpoint
func (c *Client) BuildShowRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ShowGetbalancePath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("getbalance", "show", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeShowRequest returns an encoder for requests sent to the getbalance
// show server.
func EncodeShowRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*getbalance.ShowPayload)
		if !ok {
			return goahttp.ErrInvalidType("getbalance", "show", "*getbalance.ShowPayload", v)
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

// DecodeShowResponse returns a decoder for responses returned by the
// getbalance show endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeShowResponse may return the following errors:
//	- "bad_request" (type *goa.ServiceError): http.StatusBadRequest
//	- error: internal error
func DecodeShowResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
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
				body ShowResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("getbalance", "show", err)
			}
			p := NewShowErrorreasonOK(&body)
			view := "default"
			vres := &getbalanceviews.Errorreason{Projected: p, View: view}
			if err = getbalanceviews.ValidateErrorreason(vres); err != nil {
				return nil, goahttp.ErrValidationError("getbalance", "show", err)
			}
			res := getbalance.NewErrorreason(vres)
			return res, nil
		case http.StatusBadRequest:
			var (
				body ShowBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("getbalance", "show", err)
			}
			err = ValidateShowBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("getbalance", "show", err)
			}
			return nil, NewShowBadRequest(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("getbalance", "show", resp.StatusCode, string(body))
		}
	}
}
