// Code generated by goa v3.1.2, DO NOT EDIT.
//
// token HTTP client types
//
// Command:
// $ goa gen github.com/wondenge/momo-go/remittance/design

package client

import (
	token "github.com/wondenge/momo-go/remittance/gen/token"
	goa "goa.design/goa/v3/pkg"
)

// CreateResponseBody is the type of the "token" service "create" endpoint HTTP
// response body.
type CreateResponseBody struct {
	// An error code.
	Error *string `form:"error,omitempty" json:"error,omitempty" xml:"error,omitempty"`
}

// CreateInternalErrorResponseBody is the type of the "token" service "create"
// endpoint HTTP response body for the "internal_error" error.
type CreateInternalErrorResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// NewCreateOAuthTokenErrorOK builds a "token" service "create" endpoint result
// from a HTTP "OK" response.
func NewCreateOAuthTokenErrorOK(body *CreateResponseBody) *token.OAuthTokenError {
	v := &token.OAuthTokenError{
		Error: body.Error,
	}

	return v
}

// NewCreateInternalError builds a token service create endpoint internal_error
// error.
func NewCreateInternalError(body *CreateInternalErrorResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// ValidateCreateInternalErrorResponseBody runs the validations defined on
// create_internal_error_response_body
func ValidateCreateInternalErrorResponseBody(body *CreateInternalErrorResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}
