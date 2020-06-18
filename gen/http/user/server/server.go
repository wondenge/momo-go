// Code generated by goa v3.1.3, DO NOT EDIT.
//
// User HTTP server
//
// Command:
// $ goa gen github.com/wondenge/momo-go/design

package server

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	user "github.com/wondenge/momo-go/gen/user"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Server lists the User service endpoint HTTP handlers.
type Server struct {
	Mounts         []*MountPoint
	NewUser        http.Handler
	NewKey         http.Handler
	GetUser        http.Handler
	GetUserDetails http.Handler
}

// ErrorNamer is an interface implemented by generated error structs that
// exposes the name of the error as defined in the design.
type ErrorNamer interface {
	ErrorName() string
}

// MountPoint holds information about the mounted endpoints.
type MountPoint struct {
	// Method is the name of the service method served by the mounted HTTP handler.
	Method string
	// Verb is the HTTP method used to match requests to the mounted handler.
	Verb string
	// Pattern is the HTTP request path pattern used to match requests to the
	// mounted handler.
	Pattern string
}

// New instantiates HTTP handlers for all the User service endpoints using the
// provided encoder and decoder. The handlers are mounted on the given mux
// using the HTTP verb and path defined in the design. errhandler is called
// whenever a response fails to be encoded. formatter is used to format errors
// returned by the service methods prior to encoding. Both errhandler and
// formatter are optional and can be nil.
func New(
	e *user.Endpoints,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"NewUser", "POST", "/v1_0/apiuser"},
			{"NewKey", "POST", "/v1_0/apiuser/{X-Reference-Id}/apikey"},
			{"GetUser", "GET", "/v1_0/apiuser/{X-Reference-Id}"},
			{"GetUserDetails", "GET", "/apiuser/{APIUser}"},
		},
		NewUser:        NewNewUserHandler(e.NewUser, mux, decoder, encoder, errhandler, formatter),
		NewKey:         NewNewKeyHandler(e.NewKey, mux, decoder, encoder, errhandler, formatter),
		GetUser:        NewGetUserHandler(e.GetUser, mux, decoder, encoder, errhandler, formatter),
		GetUserDetails: NewGetUserDetailsHandler(e.GetUserDetails, mux, decoder, encoder, errhandler, formatter),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "User" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.NewUser = m(s.NewUser)
	s.NewKey = m(s.NewKey)
	s.GetUser = m(s.GetUser)
	s.GetUserDetails = m(s.GetUserDetails)
}

// Mount configures the mux to serve the User endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountNewUserHandler(mux, h.NewUser)
	MountNewKeyHandler(mux, h.NewKey)
	MountGetUserHandler(mux, h.GetUser)
	MountGetUserDetailsHandler(mux, h.GetUserDetails)
}

// MountNewUserHandler configures the mux to serve the "User" service "NewUser"
// endpoint.
func MountNewUserHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/v1_0/apiuser", f)
}

// NewNewUserHandler creates a HTTP handler which loads the HTTP request and
// calls the "User" service "NewUser" endpoint.
func NewNewUserHandler(
	endpoint endpoint.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeNewUserRequest(mux, decoder)
		encodeResponse = EncodeNewUserResponse(encoder)
		encodeError    = EncodeNewUserError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "NewUser")
		ctx = context.WithValue(ctx, goa.ServiceKey, "User")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountNewKeyHandler configures the mux to serve the "User" service "NewKey"
// endpoint.
func MountNewKeyHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/v1_0/apiuser/{X-Reference-Id}/apikey", f)
}

// NewNewKeyHandler creates a HTTP handler which loads the HTTP request and
// calls the "User" service "NewKey" endpoint.
func NewNewKeyHandler(
	endpoint endpoint.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeNewKeyRequest(mux, decoder)
		encodeResponse = EncodeNewKeyResponse(encoder)
		encodeError    = EncodeNewKeyError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "NewKey")
		ctx = context.WithValue(ctx, goa.ServiceKey, "User")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountGetUserHandler configures the mux to serve the "User" service "GetUser"
// endpoint.
func MountGetUserHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/v1_0/apiuser/{X-Reference-Id}", f)
}

// NewGetUserHandler creates a HTTP handler which loads the HTTP request and
// calls the "User" service "GetUser" endpoint.
func NewGetUserHandler(
	endpoint endpoint.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeGetUserRequest(mux, decoder)
		encodeResponse = EncodeGetUserResponse(encoder)
		encodeError    = EncodeGetUserError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "GetUser")
		ctx = context.WithValue(ctx, goa.ServiceKey, "User")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountGetUserDetailsHandler configures the mux to serve the "User" service
// "GetUserDetails" endpoint.
func MountGetUserDetailsHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/apiuser/{APIUser}", f)
}

// NewGetUserDetailsHandler creates a HTTP handler which loads the HTTP request
// and calls the "User" service "GetUserDetails" endpoint.
func NewGetUserDetailsHandler(
	endpoint endpoint.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeGetUserDetailsRequest(mux, decoder)
		encodeResponse = EncodeGetUserDetailsResponse(encoder)
		encodeError    = EncodeGetUserDetailsError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "GetUserDetails")
		ctx = context.WithValue(ctx, goa.ServiceKey, "User")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}
