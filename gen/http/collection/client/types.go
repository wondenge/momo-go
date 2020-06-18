// Code generated by goa v3.1.3, DO NOT EDIT.
//
// Collection HTTP client types
//
// Command:
// $ goa gen github.com/wondenge/momo-go/design

package client

import (
	collection "github.com/wondenge/momo-go/gen/collection"
	collectionviews "github.com/wondenge/momo-go/gen/collection/views"
	goa "goa.design/goa/v3/pkg"
)

// PaymentRequestRequestBody is the type of the "Collection" service
// "PaymentRequest" endpoint HTTP request body.
type PaymentRequestRequestBody struct {
	// Amount that will be debited from the payer account.
	Amount *string `form:"amount,omitempty" json:"amount,omitempty" xml:"amount,omitempty"`
	// ISO4217 Currency
	Currency *string `form:"currency,omitempty" json:"currency,omitempty" xml:"currency,omitempty"`
	// External id is used as a reference to the transaction.
	ExternalID *string           `form:"externalId,omitempty" json:"externalId,omitempty" xml:"externalId,omitempty"`
	Payer      *PartyRequestBody `form:"payer,omitempty" json:"payer,omitempty" xml:"payer,omitempty"`
	// Message that will be written in the payer transaction history message field.
	PayerMessage *string `form:"payerMessage,omitempty" json:"payerMessage,omitempty" xml:"payerMessage,omitempty"`
	// Message that will be written in the payee transaction history note field.
	PayeeNote *string `form:"payeeNote,omitempty" json:"payeeNote,omitempty" xml:"payeeNote,omitempty"`
}

// NewTokenResponseBody is the type of the "Collection" service "NewToken"
// endpoint HTTP response body.
type NewTokenResponseBody struct {
	// A JWT token which can be used to authorize against the other API end-points.
	AccessToken *string `form:"access_token,omitempty" json:"access_token,omitempty" xml:"access_token,omitempty"`
	// The token type.
	TokenType *string `form:"token_type,omitempty" json:"token_type,omitempty" xml:"token_type,omitempty"`
	// The validity time in seconds of the token.
	ExpiresIn *int32 `form:"expires_in,omitempty" json:"expires_in,omitempty" xml:"expires_in,omitempty"`
}

// GetBalanceResponseBody is the type of the "Collection" service "GetBalance"
// endpoint HTTP response body.
type GetBalanceResponseBody struct {
	// The available balance of the account
	AvailableBalance *string `form:"availableBalance,omitempty" json:"availableBalance,omitempty" xml:"availableBalance,omitempty"`
	// ISO4217 Currency
	Currency *string `form:"currency,omitempty" json:"currency,omitempty" xml:"currency,omitempty"`
}

// PaymentStatusResponseBody is the type of the "Collection" service
// "PaymentStatus" endpoint HTTP response body.
type PaymentStatusResponseBody struct {
	// Amount that will be debited from the payer account.
	Amount *string `form:"amount,omitempty" json:"amount,omitempty" xml:"amount,omitempty"`
	// ISO4217 Currency
	Currency *string `form:"currency,omitempty" json:"currency,omitempty" xml:"currency,omitempty"`
	// Financial transactionIdd from mobile money manager.
	FinancialTransactionID *string `form:"financialTransactionId,omitempty" json:"financialTransactionId,omitempty" xml:"financialTransactionId,omitempty"`
	// External id provided in the creation of the requestToPay transaction.
	ExternalID *string            `form:"externalId,omitempty" json:"externalId,omitempty" xml:"externalId,omitempty"`
	Payer      *PartyResponseBody `form:"payer,omitempty" json:"payer,omitempty" xml:"payer,omitempty"`
	// Message that will be written in the payer transaction history message field.
	PayerMessage *string `form:"payerMessage,omitempty" json:"payerMessage,omitempty" xml:"payerMessage,omitempty"`
	// Message that will be written in the payee transaction history note field.
	PayeeNote *string `form:"payeeNote,omitempty" json:"payeeNote,omitempty" xml:"payeeNote,omitempty"`
	// Status
	Status *string                  `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	Reason *ErrorReasonResponseBody `form:"reason,omitempty" json:"reason,omitempty" xml:"reason,omitempty"`
}

// NewTokenUnauthorizedResponseBody is the type of the "Collection" service
// "NewToken" endpoint HTTP response body for the "unauthorized" error.
type NewTokenUnauthorizedResponseBody struct {
	// An error code.
	Error *string `form:"error,omitempty" json:"error,omitempty" xml:"error,omitempty"`
}

// NewTokenInternalErrorResponseBody is the type of the "Collection" service
// "NewToken" endpoint HTTP response body for the "internal_error" error.
type NewTokenInternalErrorResponseBody struct {
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

// GetBalanceBadRequestResponseBody is the type of the "Collection" service
// "GetBalance" endpoint HTTP response body for the "bad_request" error.
type GetBalanceBadRequestResponseBody struct {
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

// GetBalanceInternalErrorResponseBody is the type of the "Collection" service
// "GetBalance" endpoint HTTP response body for the "internal_error" error.
type GetBalanceInternalErrorResponseBody struct {
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

// RetrieveAccountHolderBadRequestResponseBody is the type of the "Collection"
// service "RetrieveAccountHolder" endpoint HTTP response body for the
// "bad_request" error.
type RetrieveAccountHolderBadRequestResponseBody struct {
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

// RetrieveAccountHolderInternalErrorResponseBody is the type of the
// "Collection" service "RetrieveAccountHolder" endpoint HTTP response body for
// the "internal_error" error.
type RetrieveAccountHolderInternalErrorResponseBody struct {
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

// PaymentRequestBadRequestResponseBody is the type of the "Collection" service
// "PaymentRequest" endpoint HTTP response body for the "bad_request" error.
type PaymentRequestBadRequestResponseBody struct {
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

// PaymentRequestConflictResponseBody is the type of the "Collection" service
// "PaymentRequest" endpoint HTTP response body for the "conflict" error.
type PaymentRequestConflictResponseBody struct {
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

// PaymentRequestInternalErrorResponseBody is the type of the "Collection"
// service "PaymentRequest" endpoint HTTP response body for the
// "internal_error" error.
type PaymentRequestInternalErrorResponseBody struct {
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

// PaymentStatusBadRequestResponseBody is the type of the "Collection" service
// "PaymentStatus" endpoint HTTP response body for the "bad_request" error.
type PaymentStatusBadRequestResponseBody struct {
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

// PaymentStatusNotFoundResponseBody is the type of the "Collection" service
// "PaymentStatus" endpoint HTTP response body for the "not_found" error.
type PaymentStatusNotFoundResponseBody struct {
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

// PaymentStatusInternalErrorResponseBody is the type of the "Collection"
// service "PaymentStatus" endpoint HTTP response body for the "internal_error"
// error.
type PaymentStatusInternalErrorResponseBody struct {
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

// PartyRequestBody is used to define fields on request body types.
type PartyRequestBody struct {
	// PartyIdType
	PartyIDType *string `form:"partyIdType,omitempty" json:"partyIdType,omitempty" xml:"partyIdType,omitempty"`
	PartyID     *string `form:"partyId,omitempty" json:"partyId,omitempty" xml:"partyId,omitempty"`
}

// PartyResponseBody is used to define fields on response body types.
type PartyResponseBody struct {
	// PartyIdType
	PartyIDType *string `form:"partyIdType,omitempty" json:"partyIdType,omitempty" xml:"partyIdType,omitempty"`
	PartyID     *string `form:"partyId,omitempty" json:"partyId,omitempty" xml:"partyId,omitempty"`
}

// ErrorReasonResponseBody is used to define fields on response body types.
type ErrorReasonResponseBody struct {
	// Code
	Code *string `form:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	// message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// NewPaymentRequestRequestBody builds the HTTP request body from the payload
// of the "PaymentRequest" endpoint of the "Collection" service.
func NewPaymentRequestRequestBody(p *collection.RequestToPay) *PaymentRequestRequestBody {
	body := &PaymentRequestRequestBody{
		Amount:       p.Amount,
		Currency:     p.Currency,
		ExternalID:   p.ExternalID,
		PayerMessage: p.PayerMessage,
		PayeeNote:    p.PayeeNote,
	}
	if p.Payer != nil {
		body.Payer = marshalCollectionPartyToPartyRequestBody(p.Payer)
	}
	return body
}

// NewNewTokenTokenPost200ApplicationJSONResponseOK builds a "Collection"
// service "NewToken" endpoint result from a HTTP "OK" response.
func NewNewTokenTokenPost200ApplicationJSONResponseOK(body *NewTokenResponseBody) *collectionviews.TokenPost200ApplicationJSONResponseView {
	v := &collectionviews.TokenPost200ApplicationJSONResponseView{
		AccessToken: body.AccessToken,
		TokenType:   body.TokenType,
		ExpiresIn:   body.ExpiresIn,
	}

	return v
}

// NewNewTokenUnauthorized builds a Collection service NewToken endpoint
// unauthorized error.
func NewNewTokenUnauthorized(body *NewTokenUnauthorizedResponseBody) *collection.TokenPost401ApplicationJSONResponse {
	v := &collection.TokenPost401ApplicationJSONResponse{
		Error: body.Error,
	}

	return v
}

// NewNewTokenInternalError builds a Collection service NewToken endpoint
// internal_error error.
func NewNewTokenInternalError(body *NewTokenInternalErrorResponseBody) *goa.ServiceError {
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

// NewGetBalanceBalanceOK builds a "Collection" service "GetBalance" endpoint
// result from a HTTP "OK" response.
func NewGetBalanceBalanceOK(body *GetBalanceResponseBody) *collectionviews.BalanceView {
	v := &collectionviews.BalanceView{
		AvailableBalance: body.AvailableBalance,
		Currency:         body.Currency,
	}

	return v
}

// NewGetBalanceBadRequest builds a Collection service GetBalance endpoint
// bad_request error.
func NewGetBalanceBadRequest(body *GetBalanceBadRequestResponseBody) *goa.ServiceError {
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

// NewGetBalanceInternalError builds a Collection service GetBalance endpoint
// internal_error error.
func NewGetBalanceInternalError(body *GetBalanceInternalErrorResponseBody) *goa.ServiceError {
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

// NewRetrieveAccountHolderBadRequest builds a Collection service
// RetrieveAccountHolder endpoint bad_request error.
func NewRetrieveAccountHolderBadRequest(body *RetrieveAccountHolderBadRequestResponseBody) *goa.ServiceError {
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

// NewRetrieveAccountHolderInternalError builds a Collection service
// RetrieveAccountHolder endpoint internal_error error.
func NewRetrieveAccountHolderInternalError(body *RetrieveAccountHolderInternalErrorResponseBody) *goa.ServiceError {
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

// NewPaymentRequestBadRequest builds a Collection service PaymentRequest
// endpoint bad_request error.
func NewPaymentRequestBadRequest(body *PaymentRequestBadRequestResponseBody) *goa.ServiceError {
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

// NewPaymentRequestConflict builds a Collection service PaymentRequest
// endpoint conflict error.
func NewPaymentRequestConflict(body *PaymentRequestConflictResponseBody) *goa.ServiceError {
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

// NewPaymentRequestInternalError builds a Collection service PaymentRequest
// endpoint internal_error error.
func NewPaymentRequestInternalError(body *PaymentRequestInternalErrorResponseBody) *goa.ServiceError {
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

// NewPaymentStatusRequestToPayResultOK builds a "Collection" service
// "PaymentStatus" endpoint result from a HTTP "OK" response.
func NewPaymentStatusRequestToPayResultOK(body *PaymentStatusResponseBody) *collectionviews.RequestToPayResultView {
	v := &collectionviews.RequestToPayResultView{
		Amount:                 body.Amount,
		Currency:               body.Currency,
		FinancialTransactionID: body.FinancialTransactionID,
		ExternalID:             body.ExternalID,
		PayerMessage:           body.PayerMessage,
		PayeeNote:              body.PayeeNote,
		Status:                 body.Status,
	}
	if body.Payer != nil {
		v.Payer = unmarshalPartyResponseBodyToCollectionviewsPartyView(body.Payer)
	}
	if body.Reason != nil {
		v.Reason = unmarshalErrorReasonResponseBodyToCollectionviewsErrorReasonView(body.Reason)
	}

	return v
}

// NewPaymentStatusBadRequest builds a Collection service PaymentStatus
// endpoint bad_request error.
func NewPaymentStatusBadRequest(body *PaymentStatusBadRequestResponseBody) *goa.ServiceError {
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

// NewPaymentStatusNotFound builds a Collection service PaymentStatus endpoint
// not_found error.
func NewPaymentStatusNotFound(body *PaymentStatusNotFoundResponseBody) *goa.ServiceError {
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

// NewPaymentStatusInternalError builds a Collection service PaymentStatus
// endpoint internal_error error.
func NewPaymentStatusInternalError(body *PaymentStatusInternalErrorResponseBody) *goa.ServiceError {
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

// ValidateNewTokenInternalErrorResponseBody runs the validations defined on
// NewToken_internal_error_Response_Body
func ValidateNewTokenInternalErrorResponseBody(body *NewTokenInternalErrorResponseBody) (err error) {
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

// ValidateGetBalanceBadRequestResponseBody runs the validations defined on
// GetBalance_bad_request_Response_Body
func ValidateGetBalanceBadRequestResponseBody(body *GetBalanceBadRequestResponseBody) (err error) {
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

// ValidateGetBalanceInternalErrorResponseBody runs the validations defined on
// GetBalance_internal_error_Response_Body
func ValidateGetBalanceInternalErrorResponseBody(body *GetBalanceInternalErrorResponseBody) (err error) {
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

// ValidateRetrieveAccountHolderBadRequestResponseBody runs the validations
// defined on RetrieveAccountHolder_bad_request_Response_Body
func ValidateRetrieveAccountHolderBadRequestResponseBody(body *RetrieveAccountHolderBadRequestResponseBody) (err error) {
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

// ValidateRetrieveAccountHolderInternalErrorResponseBody runs the validations
// defined on RetrieveAccountHolder_internal_error_Response_Body
func ValidateRetrieveAccountHolderInternalErrorResponseBody(body *RetrieveAccountHolderInternalErrorResponseBody) (err error) {
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

// ValidatePaymentRequestBadRequestResponseBody runs the validations defined on
// PaymentRequest_bad_request_Response_Body
func ValidatePaymentRequestBadRequestResponseBody(body *PaymentRequestBadRequestResponseBody) (err error) {
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

// ValidatePaymentRequestConflictResponseBody runs the validations defined on
// PaymentRequest_conflict_Response_Body
func ValidatePaymentRequestConflictResponseBody(body *PaymentRequestConflictResponseBody) (err error) {
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

// ValidatePaymentRequestInternalErrorResponseBody runs the validations defined
// on PaymentRequest_internal_error_Response_Body
func ValidatePaymentRequestInternalErrorResponseBody(body *PaymentRequestInternalErrorResponseBody) (err error) {
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

// ValidatePaymentStatusBadRequestResponseBody runs the validations defined on
// PaymentStatus_bad_request_Response_Body
func ValidatePaymentStatusBadRequestResponseBody(body *PaymentStatusBadRequestResponseBody) (err error) {
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

// ValidatePaymentStatusNotFoundResponseBody runs the validations defined on
// PaymentStatus_not_found_Response_Body
func ValidatePaymentStatusNotFoundResponseBody(body *PaymentStatusNotFoundResponseBody) (err error) {
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

// ValidatePaymentStatusInternalErrorResponseBody runs the validations defined
// on PaymentStatus_internal_error_Response_Body
func ValidatePaymentStatusInternalErrorResponseBody(body *PaymentStatusInternalErrorResponseBody) (err error) {
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

// ValidatePartyRequestBody runs the validations defined on PartyRequestBody
func ValidatePartyRequestBody(body *PartyRequestBody) (err error) {
	if body.PartyIDType != nil {
		if !(*body.PartyIDType == "MSISDN" || *body.PartyIDType == "EMAIL" || *body.PartyIDType == "PARTY_CODE") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.partyIdType", *body.PartyIDType, []interface{}{"MSISDN", "EMAIL", "PARTY_CODE"}))
		}
	}
	return
}

// ValidatePartyResponseBody runs the validations defined on PartyResponseBody
func ValidatePartyResponseBody(body *PartyResponseBody) (err error) {
	if body.PartyIDType != nil {
		if !(*body.PartyIDType == "MSISDN" || *body.PartyIDType == "EMAIL" || *body.PartyIDType == "PARTY_CODE") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.partyIdType", *body.PartyIDType, []interface{}{"MSISDN", "EMAIL", "PARTY_CODE"}))
		}
	}
	return
}

// ValidateErrorReasonResponseBody runs the validations defined on
// ErrorReasonResponseBody
func ValidateErrorReasonResponseBody(body *ErrorReasonResponseBody) (err error) {
	if body.Code != nil {
		if !(*body.Code == "PAYEE_NOT_FOUND" || *body.Code == "PAYER_NOT_FOUND" || *body.Code == "NOT_ALLOWED" || *body.Code == "NOT_ALLOWED_TARGET_ENVIRONMENT" || *body.Code == "INVALID_CALLBACK_URL_HOST" || *body.Code == "INVALID_CURRENCY" || *body.Code == "SERVICE_UNAVAILABLE" || *body.Code == "INTERNAL_PROCESSING_ERROR" || *body.Code == "NOT_ENOUGH_FUNDS" || *body.Code == "PAYER_LIMIT_REACHED" || *body.Code == "PAYEE_NOT_ALLOWED_TO_RECEIVE" || *body.Code == "PAYMENT_NOT_APPROVED" || *body.Code == "RESOURCE_NOT_FOUND" || *body.Code == "APPROVAL_REJECTED" || *body.Code == "EXPIRED" || *body.Code == "TRANSACTION_CANCELED" || *body.Code == "RESOURCE_ALREADY_EXIST") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.code", *body.Code, []interface{}{"PAYEE_NOT_FOUND", "PAYER_NOT_FOUND", "NOT_ALLOWED", "NOT_ALLOWED_TARGET_ENVIRONMENT", "INVALID_CALLBACK_URL_HOST", "INVALID_CURRENCY", "SERVICE_UNAVAILABLE", "INTERNAL_PROCESSING_ERROR", "NOT_ENOUGH_FUNDS", "PAYER_LIMIT_REACHED", "PAYEE_NOT_ALLOWED_TO_RECEIVE", "PAYMENT_NOT_APPROVED", "RESOURCE_NOT_FOUND", "APPROVAL_REJECTED", "EXPIRED", "TRANSACTION_CANCELED", "RESOURCE_ALREADY_EXIST"}))
		}
	}
	return
}
