// Code generated by goa v3.1.2, DO NOT EDIT.
//
// getbalance HTTP client CLI support package
//
// Command:
// $ goa gen github.com/wondenge/momo-go/remittance/design

package client

import (
	getbalance "github.com/wondenge/momo-go/remittance/gen/getbalance"
)

// BuildShowPayload builds the payload for the getbalance show endpoint from
// CLI flags.
func BuildShowPayload(getbalanceShowAuthorization string, getbalanceShowXTargetEnvironment string, getbalanceShowOcpApimSubscriptionKey string) (*getbalance.ShowPayload, error) {
	var authorization *string
	{
		if getbalanceShowAuthorization != "" {
			authorization = &getbalanceShowAuthorization
		}
	}
	var xTargetEnvironment string
	{
		xTargetEnvironment = getbalanceShowXTargetEnvironment
	}
	var ocpApimSubscriptionKey string
	{
		ocpApimSubscriptionKey = getbalanceShowOcpApimSubscriptionKey
	}
	v := &getbalance.ShowPayload{}
	v.Authorization = authorization
	v.XTargetEnvironment = xTargetEnvironment
	v.OcpApimSubscriptionKey = ocpApimSubscriptionKey

	return v, nil
}
