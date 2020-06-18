// Code generated by goa v3.1.3, DO NOT EDIT.
//
// Remittance go-kit HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/wondenge/momo-go/design

package server

import (
	"context"
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/wondenge/momo-go/gen/http/Remittance/server"
	goahttp "goa.design/goa/v3/http"
)

// EncodeNewTokenResponse returns a go-kit EncodeResponseFunc suitable for
// encoding Remittance NewToken responses.
func EncodeNewTokenResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodeNewTokenResponse(encoder)
}

// DecodeNewTokenRequest returns a go-kit DecodeRequestFunc suitable for
// decoding Remittance NewToken requests.
func DecodeNewTokenRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc {
	dec := server.DecodeNewTokenRequest(mux, decoder)
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		r = r.WithContext(ctx)
		return dec(r)
	}
}

// EncodeNewTokenError returns a go-kit EncodeResponseFunc suitable for
// encoding errors returned by the Remittance NewToken endpoint.
func EncodeNewTokenError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) kithttp.ErrorEncoder {
	enc := server.EncodeNewTokenError(encoder, formatter)
	return func(ctx context.Context, err error, w http.ResponseWriter) {
		enc(ctx, w, err)
	}
}

// EncodeGetBalanceResponse returns a go-kit EncodeResponseFunc suitable for
// encoding Remittance GetBalance responses.
func EncodeGetBalanceResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodeGetBalanceResponse(encoder)
}

// DecodeGetBalanceRequest returns a go-kit DecodeRequestFunc suitable for
// decoding Remittance GetBalance requests.
func DecodeGetBalanceRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc {
	dec := server.DecodeGetBalanceRequest(mux, decoder)
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		r = r.WithContext(ctx)
		return dec(r)
	}
}

// EncodeGetBalanceError returns a go-kit EncodeResponseFunc suitable for
// encoding errors returned by the Remittance GetBalance endpoint.
func EncodeGetBalanceError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) kithttp.ErrorEncoder {
	enc := server.EncodeGetBalanceError(encoder, formatter)
	return func(ctx context.Context, err error, w http.ResponseWriter) {
		enc(ctx, w, err)
	}
}

// EncodeRetrieveAccountHolderResponse returns a go-kit EncodeResponseFunc
// suitable for encoding Remittance RetrieveAccountHolder responses.
func EncodeRetrieveAccountHolderResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodeRetrieveAccountHolderResponse(encoder)
}

// DecodeRetrieveAccountHolderRequest returns a go-kit DecodeRequestFunc
// suitable for decoding Remittance RetrieveAccountHolder requests.
func DecodeRetrieveAccountHolderRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc {
	dec := server.DecodeRetrieveAccountHolderRequest(mux, decoder)
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		r = r.WithContext(ctx)
		return dec(r)
	}
}

// EncodeRetrieveAccountHolderError returns a go-kit EncodeResponseFunc
// suitable for encoding errors returned by the Remittance
// RetrieveAccountHolder endpoint.
func EncodeRetrieveAccountHolderError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) kithttp.ErrorEncoder {
	enc := server.EncodeRetrieveAccountHolderError(encoder, formatter)
	return func(ctx context.Context, err error, w http.ResponseWriter) {
		enc(ctx, w, err)
	}
}

// EncodeTransferResponse returns a go-kit EncodeResponseFunc suitable for
// encoding Remittance Transfer responses.
func EncodeTransferResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodeTransferResponse(encoder)
}

// DecodeTransferRequest returns a go-kit DecodeRequestFunc suitable for
// decoding Remittance Transfer requests.
func DecodeTransferRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc {
	dec := server.DecodeTransferRequest(mux, decoder)
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		r = r.WithContext(ctx)
		return dec(r)
	}
}

// EncodeTransferError returns a go-kit EncodeResponseFunc suitable for
// encoding errors returned by the Remittance Transfer endpoint.
func EncodeTransferError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) kithttp.ErrorEncoder {
	enc := server.EncodeTransferError(encoder, formatter)
	return func(ctx context.Context, err error, w http.ResponseWriter) {
		enc(ctx, w, err)
	}
}

// EncodeTransferStatusResponse returns a go-kit EncodeResponseFunc suitable
// for encoding Remittance TransferStatus responses.
func EncodeTransferStatusResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodeTransferStatusResponse(encoder)
}

// DecodeTransferStatusRequest returns a go-kit DecodeRequestFunc suitable for
// decoding Remittance TransferStatus requests.
func DecodeTransferStatusRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc {
	dec := server.DecodeTransferStatusRequest(mux, decoder)
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		r = r.WithContext(ctx)
		return dec(r)
	}
}

// EncodeTransferStatusError returns a go-kit EncodeResponseFunc suitable for
// encoding errors returned by the Remittance TransferStatus endpoint.
func EncodeTransferStatusError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) kithttp.ErrorEncoder {
	enc := server.EncodeTransferStatusError(encoder, formatter)
	return func(ctx context.Context, err error, w http.ResponseWriter) {
		enc(ctx, w, err)
	}
}
