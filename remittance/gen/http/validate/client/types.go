// Code generated by goa v3.1.2, DO NOT EDIT.
//
// validate HTTP client types
//
// Command:
// $ goa gen github.com/wondenge/momo-go/remittance/design

package client

import (
	validateviews "github.com/wondenge/momo-go/remittance/gen/validate/views"
	goa "goa.design/goa/v3/pkg"
)

// ShowResponseBody is the type of the "validate" service "show" endpoint HTTP
// response body.
type ShowResponseBody struct {
	Code *string `form:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	// message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ShowBadRequestResponseBody is the type of the "validate" service "show"
// endpoint HTTP response body for the "bad_request" error.
type ShowBadRequestResponseBody struct {
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

// NewShowErrorreasonOK builds a "validate" service "show" endpoint result from
// a HTTP "OK" response.
func NewShowErrorreasonOK(body *ShowResponseBody) *validateviews.ErrorreasonView {
	v := &validateviews.ErrorreasonView{
		Code:    body.Code,
		Message: body.Message,
	}

	return v
}

// NewShowBadRequest builds a validate service show endpoint bad_request error.
func NewShowBadRequest(body *ShowBadRequestResponseBody) *goa.ServiceError {
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

// ValidateShowBadRequestResponseBody runs the validations defined on
// show_bad_request_response_body
func ValidateShowBadRequestResponseBody(body *ShowBadRequestResponseBody) (err error) {
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
