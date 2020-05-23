// Code generated by goa v3.1.2, DO NOT EDIT.
//
// token HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/wondenge/momo-go/collection/design

package server

import (
	"context"
	"net/http"

	token "github.com/wondenge/momo-go/collection/gen/token"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeCreateResponse returns an encoder for responses returned by the token
// create endpoint.
func EncodeCreateResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*token.OAuthTokenError)
		enc := encoder(ctx, w)
		body := NewCreateResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeCreateRequest returns a decoder for requests sent to the token create
// endpoint.
func DecodeCreateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			ocpApimSubscriptionKey string
			err                    error
		)
		ocpApimSubscriptionKey = r.Header.Get(" Ocp-Apim-Subscription-Key")
		if ocpApimSubscriptionKey == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError(" Ocp-Apim-Subscription-Key", "header"))
		}
		if err != nil {
			return nil, err
		}
		payload := NewCreatePayload(ocpApimSubscriptionKey)
		user, pass, ok := r.BasicAuth()
		if !ok {
			return nil, goa.MissingFieldError("Authorization", "header")
		}
		payload.APIKey = user
		payload.APISecret = pass

		return payload, nil
	}
}

// EncodeCreateError returns an encoder for errors returned by the create token
// endpoint.
func EncodeCreateError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "internal_error":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewCreateInternalErrorResponseBody(res)
			}
			w.Header().Set("goa-error", "internal_error")
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}
