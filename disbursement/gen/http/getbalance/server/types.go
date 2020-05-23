// Code generated by goa v3.1.2, DO NOT EDIT.
//
// getbalance HTTP server types
//
// Command:
// $ goa gen github.com/wondenge/momo-go/disbursement/design

package server

import (
	getbalance "github.com/wondenge/momo-go/disbursement/gen/getbalance"
	getbalanceviews "github.com/wondenge/momo-go/disbursement/gen/getbalance/views"
	goa "goa.design/goa/v3/pkg"
)

// ShowResponseBody is the type of the "getbalance" service "show" endpoint
// HTTP response body.
type ShowResponseBody struct {
	Code *string `form:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	// message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ShowBadRequestResponseBody is the type of the "getbalance" service "show"
// endpoint HTTP response body for the "bad_request" error.
type ShowBadRequestResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// NewShowResponseBody builds the HTTP response body from the result of the
// "show" endpoint of the "getbalance" service.
func NewShowResponseBody(res *getbalanceviews.ErrorreasonView) *ShowResponseBody {
	body := &ShowResponseBody{
		Code:    res.Code,
		Message: res.Message,
	}
	return body
}

// NewShowBadRequestResponseBody builds the HTTP response body from the result
// of the "show" endpoint of the "getbalance" service.
func NewShowBadRequestResponseBody(res *goa.ServiceError) *ShowBadRequestResponseBody {
	body := &ShowBadRequestResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewShowPayload builds a getbalance service show endpoint payload.
func NewShowPayload(authorization *string, xTargetEnvironment string, ocpApimSubscriptionKey string) *getbalance.ShowPayload {
	v := &getbalance.ShowPayload{}
	v.Authorization = authorization
	v.XTargetEnvironment = xTargetEnvironment
	v.OcpApimSubscriptionKey = ocpApimSubscriptionKey

	return v
}
