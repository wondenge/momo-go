// Code generated by goa v3.1.2, DO NOT EDIT.
//
// requestopay service
//
// Command:
// $ goa gen github.com/wondenge/momo-go/collection/design

package requestopay

import (
	"context"

	requestopayviews "github.com/wondenge/momo-go/collection/gen/requestopay/views"
	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// Service is the requestopay service interface.
type Service interface {
	// Request a payment from a consumer (Payer).
	Create(context.Context, *CreatePayload) (res *Errorreason, err error)
	// Get the status of a request to pay.
	Get(context.Context, *GetPayload) (res *Errorreason, err error)
}

// Auther defines the authorization functions to be implemented by the service.
type Auther interface {
	// JWTAuth implements the authorization logic for the JWT security scheme.
	JWTAuth(ctx context.Context, token string, schema *security.JWTScheme) (context.Context, error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "requestopay"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [2]string{"create", "get"}

// CreatePayload is the payload type of the requestopay service create method.
type CreatePayload struct {
	// Authorization header used for Basic authentication
	Authorization *string
	// URL to the server where the callback should be sent.
	XCallbackURL *string
	// Resource ID of the created request to pay transaction.
	XReferenceID string
	// Identifier of the EWP system.
	XTargetEnvironment string
	// Subscription Key
	OcpApimSubscriptionKey string
}

// Errorreason is the result type of the requestopay service create method.
type Errorreason struct {
	Code *string
	// message
	Message *string
}

// GetPayload is the payload type of the requestopay service get method.
type GetPayload struct {
	// UUID of transaction to get result
	ReferenceID *string
	// Authorization header used for Basic authentication
	Authorization *string
	// Identifier of the EWP system.
	XTargetEnvironment string
	// Subscription Key.
	OcpApimSubscriptionKey string
}

// MakeInternalError builds a goa.ServiceError from an error.
func MakeInternalError(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "internal_error",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// MakeBadRequest builds a goa.ServiceError from an error.
func MakeBadRequest(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "bad_request",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// NewErrorreason initializes result type Errorreason from viewed result type
// Errorreason.
func NewErrorreason(vres *requestopayviews.Errorreason) *Errorreason {
	return newErrorreason(vres.Projected)
}

// NewViewedErrorreason initializes viewed result type Errorreason from result
// type Errorreason using the given view.
func NewViewedErrorreason(res *Errorreason, view string) *requestopayviews.Errorreason {
	p := newErrorreasonView(res)
	return &requestopayviews.Errorreason{Projected: p, View: "default"}
}

// newErrorreason converts projected type Errorreason to service type
// Errorreason.
func newErrorreason(vres *requestopayviews.ErrorreasonView) *Errorreason {
	res := &Errorreason{
		Code:    vres.Code,
		Message: vres.Message,
	}
	return res
}

// newErrorreasonView projects result type Errorreason to projected type
// ErrorreasonView using the "default" view.
func newErrorreasonView(res *Errorreason) *requestopayviews.ErrorreasonView {
	vres := &requestopayviews.ErrorreasonView{
		Code:    res.Code,
		Message: res.Message,
	}
	return vres
}
