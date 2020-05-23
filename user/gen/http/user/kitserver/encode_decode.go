// Code generated by goa v3.1.2, DO NOT EDIT.
//
// user go-kit HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/wondenge/momo-go/user/design

package server

import (
	"context"
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/wondenge/momo-go/user/gen/http/user/server"
	goahttp "goa.design/goa/v3/http"
)

// EncodeCreateuserResponse returns a go-kit EncodeResponseFunc suitable for
// encoding user createuser responses.
func EncodeCreateuserResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodeCreateuserResponse(encoder)
}

// DecodeCreateuserRequest returns a go-kit DecodeRequestFunc suitable for
// decoding user createuser requests.
func DecodeCreateuserRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc {
	dec := server.DecodeCreateuserRequest(mux, decoder)
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		r = r.WithContext(ctx)
		return dec(r)
	}
}

// EncodeCreateuserError returns a go-kit EncodeResponseFunc suitable for
// encoding errors returned by the user createuser endpoint.
func EncodeCreateuserError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) kithttp.ErrorEncoder {
	enc := server.EncodeCreateuserError(encoder, formatter)
	return func(ctx context.Context, err error, w http.ResponseWriter) {
		enc(ctx, w, err)
	}
}

// EncodeCreatekeyResponse returns a go-kit EncodeResponseFunc suitable for
// encoding user createkey responses.
func EncodeCreatekeyResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodeCreatekeyResponse(encoder)
}

// DecodeCreatekeyRequest returns a go-kit DecodeRequestFunc suitable for
// decoding user createkey requests.
func DecodeCreatekeyRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc {
	dec := server.DecodeCreatekeyRequest(mux, decoder)
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		r = r.WithContext(ctx)
		return dec(r)
	}
}

// EncodeCreatekeyError returns a go-kit EncodeResponseFunc suitable for
// encoding errors returned by the user createkey endpoint.
func EncodeCreatekeyError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) kithttp.ErrorEncoder {
	enc := server.EncodeCreatekeyError(encoder, formatter)
	return func(ctx context.Context, err error, w http.ResponseWriter) {
		enc(ctx, w, err)
	}
}

// EncodeListResponse returns a go-kit EncodeResponseFunc suitable for encoding
// user list responses.
func EncodeListResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodeListResponse(encoder)
}

// DecodeListRequest returns a go-kit DecodeRequestFunc suitable for decoding
// user list requests.
func DecodeListRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc {
	dec := server.DecodeListRequest(mux, decoder)
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		r = r.WithContext(ctx)
		return dec(r)
	}
}

// EncodeListError returns a go-kit EncodeResponseFunc suitable for encoding
// errors returned by the user list endpoint.
func EncodeListError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) kithttp.ErrorEncoder {
	enc := server.EncodeListError(encoder, formatter)
	return func(ctx context.Context, err error, w http.ResponseWriter) {
		enc(ctx, w, err)
	}
}

// EncodeShowResponse returns a go-kit EncodeResponseFunc suitable for encoding
// user show responses.
func EncodeShowResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodeShowResponse(encoder)
}

// DecodeShowRequest returns a go-kit DecodeRequestFunc suitable for decoding
// user show requests.
func DecodeShowRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc {
	dec := server.DecodeShowRequest(mux, decoder)
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		r = r.WithContext(ctx)
		return dec(r)
	}
}

// EncodeShowError returns a go-kit EncodeResponseFunc suitable for encoding
// errors returned by the user show endpoint.
func EncodeShowError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) kithttp.ErrorEncoder {
	enc := server.EncodeShowError(encoder, formatter)
	return func(ctx context.Context, err error, w http.ResponseWriter) {
		enc(ctx, w, err)
	}
}
