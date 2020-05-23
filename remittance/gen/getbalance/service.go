// Code generated by goa v3.1.2, DO NOT EDIT.
//
// getbalance service
//
// Command:
// $ goa gen github.com/wondenge/momo-go/remittance/design

package getbalance

import (
	"context"

	getbalanceviews "github.com/wondenge/momo-go/remittance/gen/getbalance/views"
	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// Service is the getbalance service interface.
type Service interface {
	// Get the balance of the account
	Show(context.Context, *ShowPayload) (res *Errorreason, err error)
}

// Auther defines the authorization functions to be implemented by the service.
type Auther interface {
	// JWTAuth implements the authorization logic for the JWT security scheme.
	JWTAuth(ctx context.Context, token string, schema *security.JWTScheme) (context.Context, error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "getbalance"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [1]string{"show"}

// ShowPayload is the payload type of the getbalance service show method.
type ShowPayload struct {
	// Authorization header for Basic authentication.
	Authorization *string
	// Identifier of the EWP system.
	XTargetEnvironment string
	// Subscription Key
	OcpApimSubscriptionKey string
}

// Errorreason is the result type of the getbalance service show method.
type Errorreason struct {
	Code *string
	// message
	Message *string
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
func NewErrorreason(vres *getbalanceviews.Errorreason) *Errorreason {
	return newErrorreason(vres.Projected)
}

// NewViewedErrorreason initializes viewed result type Errorreason from result
// type Errorreason using the given view.
func NewViewedErrorreason(res *Errorreason, view string) *getbalanceviews.Errorreason {
	p := newErrorreasonView(res)
	return &getbalanceviews.Errorreason{Projected: p, View: "default"}
}

// newErrorreason converts projected type Errorreason to service type
// Errorreason.
func newErrorreason(vres *getbalanceviews.ErrorreasonView) *Errorreason {
	res := &Errorreason{
		Code:    vres.Code,
		Message: vres.Message,
	}
	return res
}

// newErrorreasonView projects result type Errorreason to projected type
// ErrorreasonView using the "default" view.
func newErrorreasonView(res *Errorreason) *getbalanceviews.ErrorreasonView {
	vres := &getbalanceviews.ErrorreasonView{
		Code:    res.Code,
		Message: res.Message,
	}
	return vres
}
