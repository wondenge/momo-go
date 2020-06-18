// Code generated by goa v3.1.3, DO NOT EDIT.
//
// remittance service
//
// Command:
// $ goa gen github.com/wondenge/momo-go/design

package remittance

import (
	"context"

	remittanceviews "github.com/wondenge/momo-go/gen/remittance/views"
	goa "goa.design/goa/v3/pkg"
)

// Service is the remittance service interface.
type Service interface {
	// Creates an Access Token.
	NewToken(context.Context, string) (res *TokenPost200ApplicationJSONResponse, err error)
	// Get the balance of the account
	GetBalance(context.Context, string) (res *Balance, err error)
	// Checks if an account holder is registered and active in the system
	RetrieveAccountHolder(context.Context, *RetrieveAccountHolderPayload) (res string, err error)
	// Request a payment from a consumer (Payer).
	Transfer(context.Context, *Transfer1) (res string, err error)
	// Get the status of a request to pay.
	TransferStatus(context.Context, *TransferStatusPayload) (res *TransferResult, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "remittance"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [5]string{"NewToken", "GetBalance", "RetrieveAccountHolder", "Transfer", "TransferStatus"}

// TokenPost200ApplicationJSONResponse is the result type of the remittance
// service NewToken method.
type TokenPost200ApplicationJSONResponse struct {
	// A JWT token which can be used to authorize against the other API end-points.
	AccessToken *string
	// The token type.
	TokenType *string
	// The validity time in seconds of the token.
	ExpiresIn *int32
}

// Balance is the result type of the remittance service GetBalance method.
type Balance struct {
	// The available balance of the account
	AvailableBalance *string
	// ISO4217 Currency
	Currency *string
}

// RetrieveAccountHolderPayload is the payload type of the remittance service
// RetrieveAccountHolder method.
type RetrieveAccountHolderPayload struct {
	// Specifies the type of the party ID
	AccountHolderIDType string
	// The party number.
	AccountHolderID string
}

// Transfer1 is the payload type of the remittance service Transfer method.
type Transfer1 struct {
	// Amount that will be debited from the payer account.
	Amount *string
	// ISO4217 Currency
	Currency *string
	// External id is used as a reference to the transaction.
	ExternalID *string
	Payee      *Party
	// Message that will be written in the payer transaction history message field.
	PayerMessage *string
	// Message that will be written in the payee transaction history note field.
	PayeeNote *string
}

// TransferStatusPayload is the payload type of the remittance service
// TransferStatus method.
type TransferStatusPayload struct {
	// UUID of transaction to get result
	ReferenceID string
}

// TransferResult is the result type of the remittance service TransferStatus
// method.
type TransferResult struct {
	// Amount that will be debited from the payer account.
	Amount *string
	// ISO4217 Currency
	Currency *string
	// Financial transactionIdd from mobile money manager
	FinancialTransactionID *string
	// External id is used as a reference to the transaction
	ExternalID *string
	Payee      *Party
	// Message that will be written in the payer transaction history message field.
	PayerMessage *string
	// Message that will be written in the payee transaction history note field.
	PayeeNote *string
	// Status
	Status *string
	Reason *ErrorReason
}

// Party identifies a account holder in the wallet platform.
type Party struct {
	// PartyIdType
	PartyIDType *string
	PartyID     *string
}

// Error Reason
type ErrorReason struct {
	// Code
	Code *string
	// message
	Message *string
}

// Token Post401 Application Json Response
type TokenPost401ApplicationJSONResponse struct {
	// An error code.
	TokenError *string
}

// Error returns an error description.
func (e *TokenPost401ApplicationJSONResponse) Error() string {
	return "Token Post401 Application Json Response"
}

// ErrorName returns "TokenPost401ApplicationJsonResponse".
func (e *TokenPost401ApplicationJSONResponse) ErrorName() string {
	return "unauthorized"
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

// MakeConflict builds a goa.ServiceError from an error.
func MakeConflict(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "conflict",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// MakeNotFound builds a goa.ServiceError from an error.
func MakeNotFound(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "not_found",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// NewTokenPost200ApplicationJSONResponse initializes result type
// TokenPost200ApplicationJSONResponse from viewed result type
// TokenPost200ApplicationJSONResponse.
func NewTokenPost200ApplicationJSONResponse(vres *remittanceviews.TokenPost200ApplicationJSONResponse) *TokenPost200ApplicationJSONResponse {
	return newTokenPost200ApplicationJSONResponse(vres.Projected)
}

// NewViewedTokenPost200ApplicationJSONResponse initializes viewed result type
// TokenPost200ApplicationJSONResponse from result type
// TokenPost200ApplicationJSONResponse using the given view.
func NewViewedTokenPost200ApplicationJSONResponse(res *TokenPost200ApplicationJSONResponse, view string) *remittanceviews.TokenPost200ApplicationJSONResponse {
	p := newTokenPost200ApplicationJSONResponseView(res)
	return &remittanceviews.TokenPost200ApplicationJSONResponse{Projected: p, View: "default"}
}

// NewBalance initializes result type Balance from viewed result type Balance.
func NewBalance(vres *remittanceviews.Balance) *Balance {
	return newBalance(vres.Projected)
}

// NewViewedBalance initializes viewed result type Balance from result type
// Balance using the given view.
func NewViewedBalance(res *Balance, view string) *remittanceviews.Balance {
	p := newBalanceView(res)
	return &remittanceviews.Balance{Projected: p, View: "default"}
}

// NewTransferResult initializes result type TransferResult from viewed result
// type TransferResult.
func NewTransferResult(vres *remittanceviews.TransferResult) *TransferResult {
	return newTransferResult(vres.Projected)
}

// NewViewedTransferResult initializes viewed result type TransferResult from
// result type TransferResult using the given view.
func NewViewedTransferResult(res *TransferResult, view string) *remittanceviews.TransferResult {
	p := newTransferResultView(res)
	return &remittanceviews.TransferResult{Projected: p, View: "default"}
}

// newTokenPost200ApplicationJSONResponse converts projected type
// TokenPost200ApplicationJSONResponse to service type
// TokenPost200ApplicationJSONResponse.
func newTokenPost200ApplicationJSONResponse(vres *remittanceviews.TokenPost200ApplicationJSONResponseView) *TokenPost200ApplicationJSONResponse {
	res := &TokenPost200ApplicationJSONResponse{
		AccessToken: vres.AccessToken,
		TokenType:   vres.TokenType,
		ExpiresIn:   vres.ExpiresIn,
	}
	return res
}

// newTokenPost200ApplicationJSONResponseView projects result type
// TokenPost200ApplicationJSONResponse to projected type
// TokenPost200ApplicationJSONResponseView using the "default" view.
func newTokenPost200ApplicationJSONResponseView(res *TokenPost200ApplicationJSONResponse) *remittanceviews.TokenPost200ApplicationJSONResponseView {
	vres := &remittanceviews.TokenPost200ApplicationJSONResponseView{
		AccessToken: res.AccessToken,
		TokenType:   res.TokenType,
		ExpiresIn:   res.ExpiresIn,
	}
	return vres
}

// newBalance converts projected type Balance to service type Balance.
func newBalance(vres *remittanceviews.BalanceView) *Balance {
	res := &Balance{
		AvailableBalance: vres.AvailableBalance,
		Currency:         vres.Currency,
	}
	return res
}

// newBalanceView projects result type Balance to projected type BalanceView
// using the "default" view.
func newBalanceView(res *Balance) *remittanceviews.BalanceView {
	vres := &remittanceviews.BalanceView{
		AvailableBalance: res.AvailableBalance,
		Currency:         res.Currency,
	}
	return vres
}

// newTransferResult converts projected type TransferResult to service type
// TransferResult.
func newTransferResult(vres *remittanceviews.TransferResultView) *TransferResult {
	res := &TransferResult{
		Amount:                 vres.Amount,
		Currency:               vres.Currency,
		FinancialTransactionID: vres.FinancialTransactionID,
		ExternalID:             vres.ExternalID,
		PayeeNote:              vres.PayeeNote,
		Status:                 vres.Status,
	}
	if vres.Payee != nil {
		res.Payee = transformRemittanceviewsPartyViewToParty(vres.Payee)
	}
	if vres.Reason != nil {
		res.Reason = newErrorReason(vres.Reason)
	}
	return res
}

// newTransferResultView projects result type TransferResult to projected type
// TransferResultView using the "default" view.
func newTransferResultView(res *TransferResult) *remittanceviews.TransferResultView {
	vres := &remittanceviews.TransferResultView{
		Amount:                 res.Amount,
		Currency:               res.Currency,
		FinancialTransactionID: res.FinancialTransactionID,
		ExternalID:             res.ExternalID,
		PayeeNote:              res.PayeeNote,
		Status:                 res.Status,
	}
	if res.Payee != nil {
		vres.Payee = transformPartyToRemittanceviewsPartyView(res.Payee)
	}
	if res.Reason != nil {
		vres.Reason = newErrorReasonView(res.Reason)
	}
	return vres
}

// newErrorReason converts projected type ErrorReason to service type
// ErrorReason.
func newErrorReason(vres *remittanceviews.ErrorReasonView) *ErrorReason {
	res := &ErrorReason{
		Code:    vres.Code,
		Message: vres.Message,
	}
	return res
}

// newErrorReasonView projects result type ErrorReason to projected type
// ErrorReasonView using the "default" view.
func newErrorReasonView(res *ErrorReason) *remittanceviews.ErrorReasonView {
	vres := &remittanceviews.ErrorReasonView{
		Code:    res.Code,
		Message: res.Message,
	}
	return vres
}

// transformRemittanceviewsPartyViewToParty builds a value of type *Party from
// a value of type *remittanceviews.PartyView.
func transformRemittanceviewsPartyViewToParty(v *remittanceviews.PartyView) *Party {
	if v == nil {
		return nil
	}
	res := &Party{
		PartyIDType: v.PartyIDType,
		PartyID:     v.PartyID,
	}

	return res
}

// transformPartyToRemittanceviewsPartyView builds a value of type
// *remittanceviews.PartyView from a value of type *Party.
func transformPartyToRemittanceviewsPartyView(v *Party) *remittanceviews.PartyView {
	if v == nil {
		return nil
	}
	res := &remittanceviews.PartyView{
		PartyIDType: v.PartyIDType,
		PartyID:     v.PartyID,
	}

	return res
}
