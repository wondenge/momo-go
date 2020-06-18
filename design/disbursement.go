package design

import (
	. "github.com/wondenge/momo-go/momo"
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

var _ = Service("disbursement", func() {

	HTTP(func() {
		Path("/disbursement")
	})

	// Used to authorize and authenticate other end-points of the API.
	// a) Provider system requests an access token using the API Key and API user as authentication.
	// POST/token (Authorization: Basic)
	// b) Wallet platform authenticates credentials and responds with the access token
	// Response (Access Token and Validity time)
	// c) Provider system will use the access token for any request that is sent to Wallet Platform.
	// e.g. POST /requesttopay (Authentication: Bearer)
	// Response(HTTP: 202: Accepted)
	Method("NewToken", func() {
		Description("Creates an Access Token.")
		Payload(String)
		Result(TokenPost200ApplicationJsonResponse)

		// 401
		Error("unauthorized", TokenPost401ApplicationJsonResponse, "Unauthorized")

		// "500"
		Error("internal_error", ErrorResult, "Error")

		HTTP(func() {
			POST("/token")

			Response(StatusOK)

			// Status 401
			// RFC 7231, 6.5.1
			Response("unauthorized", StatusUnauthorized)

			// Status 500
			// RFC 7231, 6.6.1
			Response("internal_error", StatusInternalServerError)
		})
	})

	// Get the balance of the account.
	Method("GetBalance", func() {
		Description("Get the balance of the account")
		Payload(String)
		Result(Balance)

		// 400
		Error("bad_request", ErrorResult, "Bad request, e.g. invalid data was sent in the request.")

		// "500"
		Error("internal_error", ErrorResult, "Internal error. The returned response contains details.")

		HTTP(func() {
			GET("/v1_0/account/balance")

			Response(StatusOK)

			// Status 400
			// RFC 7231, 6.5.1
			Response("bad_request", StatusBadRequest)

			// Status 500
			// RFC 7231, 6.6.1
			Response("internal_error", StatusInternalServerError)
		})
	})

	// Operation is used  to check if an account holder is registered and active in the system.
	Method("RetrieveAccountHolder", func() {
		Description("Checks if an account holder is registered and active in the system")
		Payload(func() {

			//  Allowed values [msisdn, email, party_code].
			//  accountHolderId should explicitly be in small letters.
			Attribute("accountHolderIdType", String, func() {
				Description("Specifies the type of the party ID")
			})

			// Validated according to the party ID type (case Sensitive).
			// msisdn - Mobile Number validated according to ITU-T E.164.
			// Validated with IsMSISDN email - Validated to be a valid e-mail format.
			// Validated with IsEmail party_code - UUID of the party.
			// Validated with IsUuid
			Attribute("accountHolderId", String, func() {
				Description("The party number.")
			})
		})
		Result(String)

		// 400
		Error("bad_request", ErrorResult, "Bad request, e.g. invalid data was sent in the request.")

		// "500"
		Error("internal_error", ErrorResult, "Internal error. The returned response contains details.")

		HTTP(func() {
			GET("/v1_0/accountholder/{accountHolderIdType}/{accountHolderId}/active")

			// True if account holder is registered and active,
			// false if the account holder is not active or not found
			Response(StatusOK)

			// Status 400
			// RFC 7231, 6.5.1
			Response("bad_request", StatusBadRequest)

			// Status 500
			// RFC 7231, 6.6.1
			Response("internal_error", StatusInternalServerError)
		})
	})

	// The payer will be asked to authorize the payment.
	// The transaction will be executed once the payer has authorized the payment.
	// The requesttopay will be in status PENDING until the transaction is authorized
	// or declined by the payer or it is timed out by the system.
	// Status of the transaction can be validated by using the GET /requesttopay/ <resourceId>
	Method("PaymentRequest", func() {
		Description("Request a payment from a consumer (Payer).")
		Payload(RequestToPay)
		Result(String)

		// 400
		Error("bad_request", ErrorReason, "Bad request, e.g. invalid data was sent in the request.")

		// 409
		Error("conflict", ErrorReason, "Conflict, duplicated reference id")

		// 500
		Error("internal_error", ErrorReason, "Internal error.")

		HTTP(func() {
			POST("/v1_0/requesttopay")

			Response(StatusAccepted)

			// Status 400
			// RFC 7231, 6.5.1
			Response("bad_request", StatusBadRequest)

			// Status 409
			// RFC 7231, 6.5.8
			Response("conflict", StatusConflict)

			// Status 500
			// RFC 7231, 6.6.1
			Response("internal_error", StatusInternalServerError)
		})
	})

	Method("PaymentStatus", func() {
		Description("Get the status of a request to pay.")
		Payload(func() {

			// Reference id  used when creating the request to pay.
			// X-Reference-Id that was passed in the post is
			// used as reference to the request.
			Attribute("referenceId", String, func() {
				Description(" UUID of transaction to get result")
			})
		})
		Result(RequestToPayResult)

		// 400
		Error("bad_request", ErrorResult, "Bad request, e.g. an incorrectly formatted reference id was provided.")

		// 404
		Error("not_found", ErrorReason, "Resource not found.")

		// 500
		Error("internal_error", ErrorReason, "Internal error.")

		HTTP(func() {
			GET("/v1_0/requesttopay/{referenceId}")

			// Note that a  failed request to pay will be returned with this status too.
			// The 'status' of the RequestToPayResult can be used to determine the outcome of the request.
			// The 'reason' field can be used to retrieve a cause in case of failure.
			Response(StatusOK)

			// Status 400
			// RFC 7231, 6.5.1
			Response("bad_request", StatusBadRequest)

			// Status 404
			// RFC 7231, 6.5.4
			Response("not_found", StatusNotFound)

			// Status 500
			// RFC 7231, 6.6.1
			Response("internal_error", StatusInternalServerError)
		})
	})
})
