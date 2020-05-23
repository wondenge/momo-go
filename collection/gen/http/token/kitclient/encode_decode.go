// Code generated by goa v3.1.2, DO NOT EDIT.
//
// token go-kit HTTP client encoders and decoders
//
// Command:
// $ goa gen github.com/wondenge/momo-go/collection/design

package client

import (
	"context"
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/wondenge/momo-go/collection/gen/http/token/client"
	goahttp "goa.design/goa/v3/http"
)

// EncodeCreateRequest returns a go-kit EncodeRequestFunc suitable for encoding
// token create requests.
func EncodeCreateRequest(encoder func(*http.Request) goahttp.Encoder) kithttp.EncodeRequestFunc {
	enc := client.EncodeCreateRequest(encoder)
	return func(_ context.Context, r *http.Request, v interface{}) error {
		return enc(r, v)
	}
}

// DecodeCreateResponse returns a go-kit DecodeResponseFunc suitable for
// decoding token create responses.
func DecodeCreateResponse(decoder func(*http.Response) goahttp.Decoder) kithttp.DecodeResponseFunc {
	dec := client.DecodeCreateResponse(decoder, false)
	return func(ctx context.Context, resp *http.Response) (interface{}, error) {
		return dec(resp)
	}
}
