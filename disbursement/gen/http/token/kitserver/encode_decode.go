// Code generated by goa v3.1.2, DO NOT EDIT.
//
// token go-kit HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/wondenge/momo-go/disbursement/design

package server

import (
	"context"
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/wondenge/momo-go/disbursement/gen/http/token/server"
	goahttp "goa.design/goa/v3/http"
)

// EncodeCreateResponse returns a go-kit EncodeResponseFunc suitable for
// encoding token create responses.
func EncodeCreateResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodeCreateResponse(encoder)
}

// DecodeCreateRequest returns a go-kit DecodeRequestFunc suitable for decoding
// token create requests.
func DecodeCreateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc {
	dec := server.DecodeCreateRequest(mux, decoder)
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		r = r.WithContext(ctx)
		return dec(r)
	}
}

// EncodeCreateError returns a go-kit EncodeResponseFunc suitable for encoding
// errors returned by the token create endpoint.
func EncodeCreateError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) kithttp.ErrorEncoder {
	enc := server.EncodeCreateError(encoder, formatter)
	return func(ctx context.Context, err error, w http.ResponseWriter) {
		enc(ctx, w, err)
	}
}
