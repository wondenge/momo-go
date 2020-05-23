// Code generated by goa v3.1.2, DO NOT EDIT.
//
// getbalance HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/wondenge/momo-go/collection/design

package server

import (
	"context"
	"net/http"
	"strings"

	getbalanceviews "github.com/wondenge/momo-go/collection/gen/getbalance/views"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeShowResponse returns an encoder for responses returned by the
// getbalance show endpoint.
func EncodeShowResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*getbalanceviews.Errorreason)
		enc := encoder(ctx, w)
		body := NewShowResponseBody(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeShowRequest returns a decoder for requests sent to the getbalance show
// endpoint.
func DecodeShowRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			authorization          *string
			xTargetEnvironment     string
			ocpApimSubscriptionKey string
			err                    error
		)
		authorizationRaw := r.Header.Get(" Authorization")
		if authorizationRaw != "" {
			authorization = &authorizationRaw
		}
		xTargetEnvironment = r.Header.Get(" X-Target-Environment")
		if xTargetEnvironment == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError(" X-Target-Environment", "header"))
		}
		ocpApimSubscriptionKey = r.Header.Get(" Ocp-Apim-Subscription-Key")
		if ocpApimSubscriptionKey == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError(" Ocp-Apim-Subscription-Key", "header"))
		}
		if err != nil {
			return nil, err
		}
		payload := NewShowPayload(authorization, xTargetEnvironment, ocpApimSubscriptionKey)
		if payload.Authorization != nil {
			if strings.Contains(*payload.Authorization, " ") {
				// Remove authorization scheme prefix (e.g. "Bearer")
				cred := strings.SplitN(*payload.Authorization, " ", 2)[1]
				payload.Authorization = &cred
			}
		}

		return payload, nil
	}
}

// EncodeShowError returns an encoder for errors returned by the show
// getbalance endpoint.
func EncodeShowError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "bad_request":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewShowBadRequestResponseBody(res)
			}
			w.Header().Set("goa-error", "bad_request")
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}
