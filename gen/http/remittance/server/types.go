// Code generated by goa v3.1.3, DO NOT EDIT.
//
// remittance HTTP server types
//
// Command:
// $ goa gen github.com/wondenge/momo-go/design

package server

import (
	remittance "github.com/wondenge/momo-go/gen/remittance"
	remittanceviews "github.com/wondenge/momo-go/gen/remittance/views"
	goa "goa.design/goa/v3/pkg"
)

// TransferRequestBody is the type of the "remittance" service "Transfer"
// endpoint HTTP request body.
type TransferRequestBody struct {
	// Amount that will be debited from the payer account.
	Amount *string `form:"amount,omitempty" json:"amount,omitempty" xml:"amount,omitempty"`
	// ISO4217 Currency
	Currency *string `form:"currency,omitempty" json:"currency,omitempty" xml:"currency,omitempty"`
	// External id is used as a reference to the transaction.
	ExternalID *string           `form:"externalId,omitempty" json:"externalId,omitempty" xml:"externalId,omitempty"`
	Payee      *PartyRequestBody `form:"payee,omitempty" json:"payee,omitempty" xml:"payee,omitempty"`
	// Message that will be written in the payer transaction history message field.
	PayerMessage *string `form:"payerMessage,omitempty" json:"payerMessage,omitempty" xml:"payerMessage,omitempty"`
	// Message that will be written in the payee transaction history note field.
	PayeeNote *string `form:"payeeNote,omitempty" json:"payeeNote,omitempty" xml:"payeeNote,omitempty"`
}

// NewTokenResponseBody is the type of the "remittance" service "NewToken"
// endpoint HTTP response body.
type NewTokenResponseBody struct {
	// A JWT token which can be used to authorize against the other API end-points.
	AccessToken *string `form:"access_token,omitempty" json:"access_token,omitempty" xml:"access_token,omitempty"`
	// The token type.
	TokenType *string `form:"token_type,omitempty" json:"token_type,omitempty" xml:"token_type,omitempty"`
	// The validity time in seconds of the token.
	ExpiresIn *int32 `form:"expires_in,omitempty" json:"expires_in,omitempty" xml:"expires_in,omitempty"`
}

// GetBalanceResponseBody is the type of the "remittance" service "GetBalance"
// endpoint HTTP response body.
type GetBalanceResponseBody struct {
	// The available balance of the account
	AvailableBalance *string `form:"availableBalance,omitempty" json:"availableBalance,omitempty" xml:"availableBalance,omitempty"`
	// ISO4217 Currency
	Currency *string `form:"currency,omitempty" json:"currency,omitempty" xml:"currency,omitempty"`
}

// TransferStatusResponseBody is the type of the "remittance" service
// "TransferStatus" endpoint HTTP response body.
type TransferStatusResponseBody struct {
	// Amount that will be debited from the payer account.
	Amount *string `form:"amount,omitempty" json:"amount,omitempty" xml:"amount,omitempty"`
	// ISO4217 Currency
	Currency *string `form:"currency,omitempty" json:"currency,omitempty" xml:"currency,omitempty"`
	// Financial transactionIdd from mobile money manager
	FinancialTransactionID *string `form:"financialTransactionId,omitempty" json:"financialTransactionId,omitempty" xml:"financialTransactionId,omitempty"`
	// External id is used as a reference to the transaction
	ExternalID *string            `form:"externalId,omitempty" json:"externalId,omitempty" xml:"externalId,omitempty"`
	Payee      *PartyResponseBody `form:"payee,omitempty" json:"payee,omitempty" xml:"payee,omitempty"`
	// Message that will be written in the payee transaction history note field.
	PayeeNote *string `form:"payeeNote,omitempty" json:"payeeNote,omitempty" xml:"payeeNote,omitempty"`
	// Status
	Status *string                  `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	Reason *ErrorReasonResponseBody `form:"reason,omitempty" json:"reason,omitempty" xml:"reason,omitempty"`
}

// NewTokenInternalErrorResponseBody is the type of the "remittance" service
// "NewToken" endpoint HTTP response body for the "internal_error" error.
type NewTokenInternalErrorResponseBody struct {
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

// NewTokenUnauthorizedResponseBody is the type of the "remittance" service
// "NewToken" endpoint HTTP response body for the "unauthorized" error.
type NewTokenUnauthorizedResponseBody struct {
	// An error code.
	TokenError *string `form:"token_error,omitempty" json:"token_error,omitempty" xml:"token_error,omitempty"`
}

// GetBalanceBadRequestResponseBody is the type of the "remittance" service
// "GetBalance" endpoint HTTP response body for the "bad_request" error.
type GetBalanceBadRequestResponseBody struct {
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

// GetBalanceInternalErrorResponseBody is the type of the "remittance" service
// "GetBalance" endpoint HTTP response body for the "internal_error" error.
type GetBalanceInternalErrorResponseBody struct {
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

// RetrieveAccountHolderBadRequestResponseBody is the type of the "remittance"
// service "RetrieveAccountHolder" endpoint HTTP response body for the
// "bad_request" error.
type RetrieveAccountHolderBadRequestResponseBody struct {
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

// RetrieveAccountHolderInternalErrorResponseBody is the type of the
// "remittance" service "RetrieveAccountHolder" endpoint HTTP response body for
// the "internal_error" error.
type RetrieveAccountHolderInternalErrorResponseBody struct {
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

// TransferBadRequestResponseBody is the type of the "remittance" service
// "Transfer" endpoint HTTP response body for the "bad_request" error.
type TransferBadRequestResponseBody struct {
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

// TransferConflictResponseBody is the type of the "remittance" service
// "Transfer" endpoint HTTP response body for the "conflict" error.
type TransferConflictResponseBody struct {
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

// TransferInternalErrorResponseBody is the type of the "remittance" service
// "Transfer" endpoint HTTP response body for the "internal_error" error.
type TransferInternalErrorResponseBody struct {
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

// TransferStatusBadRequestResponseBody is the type of the "remittance" service
// "TransferStatus" endpoint HTTP response body for the "bad_request" error.
type TransferStatusBadRequestResponseBody struct {
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

// TransferStatusNotFoundResponseBody is the type of the "remittance" service
// "TransferStatus" endpoint HTTP response body for the "not_found" error.
type TransferStatusNotFoundResponseBody struct {
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

// TransferStatusInternalErrorResponseBody is the type of the "remittance"
// service "TransferStatus" endpoint HTTP response body for the
// "internal_error" error.
type TransferStatusInternalErrorResponseBody struct {
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

// PartyRequestBody is used to define fields on request body types.
type PartyRequestBody struct {
	// PartyIdType
	PartyIDType *string `form:"partyIdType,omitempty" json:"partyIdType,omitempty" xml:"partyIdType,omitempty"`
	PartyID     *string `form:"partyId,omitempty" json:"partyId,omitempty" xml:"partyId,omitempty"`
}

// NewNewTokenResponseBody builds the HTTP response body from the result of the
// "NewToken" endpoint of the "remittance" service.
func NewNewTokenResponseBody(res *remittanceviews.TokenPost200ApplicationJSONResponseView) *NewTokenResponseBody {
	body := &NewTokenResponseBody{
		AccessToken: res.AccessToken,
		TokenType:   res.TokenType,
		ExpiresIn:   res.ExpiresIn,
	}
	return body
}

// NewGetBalanceResponseBody builds the HTTP response body from the result of
// the "GetBalance" endpoint of the "remittance" service.
func NewGetBalanceResponseBody(res *remittanceviews.BalanceView) *GetBalanceResponseBody {
	body := &GetBalanceResponseBody{
		AvailableBalance: res.AvailableBalance,
		Currency:         res.Currency,
	}
	return body
}

// NewTransferStatusResponseBody builds the HTTP response body from the result
// of the "TransferStatus" endpoint of the "remittance" service.
func NewTransferStatusResponseBody(res *remittanceviews.TransferResultView) *TransferStatusResponseBody {
	body := &TransferStatusResponseBody{
		Amount:                 res.Amount,
		Currency:               res.Currency,
		FinancialTransactionID: res.FinancialTransactionID,
		ExternalID:             res.ExternalID,
		PayeeNote:              res.PayeeNote,
		Status:                 res.Status,
	}
	if res.Payee != nil {
		body.Payee = marshalRemittanceviewsPartyViewToPartyResponseBody(res.Payee)
	}
	if res.Reason != nil {
		body.Reason = marshalRemittanceviewsErrorReasonViewToErrorReasonResponseBody(res.Reason)
	}
	return body
}

// NewNewTokenInternalErrorResponseBody builds the HTTP response body from the
// result of the "NewToken" endpoint of the "remittance" service.
func NewNewTokenInternalErrorResponseBody(res *goa.ServiceError) *NewTokenInternalErrorResponseBody {
	body := &NewTokenInternalErrorResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewNewTokenUnauthorizedResponseBody builds the HTTP response body from the
// result of the "NewToken" endpoint of the "remittance" service.
func NewNewTokenUnauthorizedResponseBody(res *remittance.TokenPost401ApplicationJSONResponse) *NewTokenUnauthorizedResponseBody {
	body := &NewTokenUnauthorizedResponseBody{
		TokenError: res.TokenError,
	}
	return body
}

// NewGetBalanceBadRequestResponseBody builds the HTTP response body from the
// result of the "GetBalance" endpoint of the "remittance" service.
func NewGetBalanceBadRequestResponseBody(res *goa.ServiceError) *GetBalanceBadRequestResponseBody {
	body := &GetBalanceBadRequestResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewGetBalanceInternalErrorResponseBody builds the HTTP response body from
// the result of the "GetBalance" endpoint of the "remittance" service.
func NewGetBalanceInternalErrorResponseBody(res *goa.ServiceError) *GetBalanceInternalErrorResponseBody {
	body := &GetBalanceInternalErrorResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewRetrieveAccountHolderBadRequestResponseBody builds the HTTP response body
// from the result of the "RetrieveAccountHolder" endpoint of the "remittance"
// service.
func NewRetrieveAccountHolderBadRequestResponseBody(res *goa.ServiceError) *RetrieveAccountHolderBadRequestResponseBody {
	body := &RetrieveAccountHolderBadRequestResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewRetrieveAccountHolderInternalErrorResponseBody builds the HTTP response
// body from the result of the "RetrieveAccountHolder" endpoint of the
// "remittance" service.
func NewRetrieveAccountHolderInternalErrorResponseBody(res *goa.ServiceError) *RetrieveAccountHolderInternalErrorResponseBody {
	body := &RetrieveAccountHolderInternalErrorResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewTransferBadRequestResponseBody builds the HTTP response body from the
// result of the "Transfer" endpoint of the "remittance" service.
func NewTransferBadRequestResponseBody(res *goa.ServiceError) *TransferBadRequestResponseBody {
	body := &TransferBadRequestResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewTransferConflictResponseBody builds the HTTP response body from the
// result of the "Transfer" endpoint of the "remittance" service.
func NewTransferConflictResponseBody(res *goa.ServiceError) *TransferConflictResponseBody {
	body := &TransferConflictResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewTransferInternalErrorResponseBody builds the HTTP response body from the
// result of the "Transfer" endpoint of the "remittance" service.
func NewTransferInternalErrorResponseBody(res *goa.ServiceError) *TransferInternalErrorResponseBody {
	body := &TransferInternalErrorResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewTransferStatusBadRequestResponseBody builds the HTTP response body from
// the result of the "TransferStatus" endpoint of the "remittance" service.
func NewTransferStatusBadRequestResponseBody(res *goa.ServiceError) *TransferStatusBadRequestResponseBody {
	body := &TransferStatusBadRequestResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewTransferStatusNotFoundResponseBody builds the HTTP response body from the
// result of the "TransferStatus" endpoint of the "remittance" service.
func NewTransferStatusNotFoundResponseBody(res *goa.ServiceError) *TransferStatusNotFoundResponseBody {
	body := &TransferStatusNotFoundResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewTransferStatusInternalErrorResponseBody builds the HTTP response body
// from the result of the "TransferStatus" endpoint of the "remittance" service.
func NewTransferStatusInternalErrorResponseBody(res *goa.ServiceError) *TransferStatusInternalErrorResponseBody {
	body := &TransferStatusInternalErrorResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewRetrieveAccountHolderPayload builds a remittance service
// RetrieveAccountHolder endpoint payload.
func NewRetrieveAccountHolderPayload(accountHolderIDType string, accountHolderID string) *remittance.RetrieveAccountHolderPayload {
	v := &remittance.RetrieveAccountHolderPayload{}
	v.AccountHolderIDType = accountHolderIDType
	v.AccountHolderID = accountHolderID

	return v
}

// NewTransfer1 builds a remittance service Transfer endpoint payload.
func NewTransfer1(body *TransferRequestBody) *remittance.Transfer1 {
	v := &remittance.Transfer1{
		Amount:       body.Amount,
		Currency:     body.Currency,
		ExternalID:   body.ExternalID,
		PayerMessage: body.PayerMessage,
		PayeeNote:    body.PayeeNote,
	}
	if body.Payee != nil {
		v.Payee = unmarshalPartyRequestBodyToRemittanceParty(body.Payee)
	}

	return v
}

// NewTransferStatusPayload builds a remittance service TransferStatus endpoint
// payload.
func NewTransferStatusPayload(referenceID string) *remittance.TransferStatusPayload {
	v := &remittance.TransferStatusPayload{}
	v.ReferenceID = referenceID

	return v
}

// ValidateTransferRequestBody runs the validations defined on
// TransferRequestBody
func ValidateTransferRequestBody(body *TransferRequestBody) (err error) {
	if body.Payee != nil {
		if err2 := ValidatePartyRequestBody(body.Payee); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
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
