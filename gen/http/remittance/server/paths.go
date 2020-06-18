// Code generated by goa v3.1.3, DO NOT EDIT.
//
// HTTP request path constructors for the Remittance service.
//
// Command:
// $ goa gen github.com/wondenge/momo-go/design

package server

import (
	"fmt"
)

// NewTokenRemittancePath returns the URL path to the Remittance service NewToken HTTP endpoint.
func NewTokenRemittancePath() string {
	return "/remittance/token"
}

// GetBalanceRemittancePath returns the URL path to the Remittance service GetBalance HTTP endpoint.
func GetBalanceRemittancePath() string {
	return "/remittance/v1_0/account/balance"
}

// RetrieveAccountHolderRemittancePath returns the URL path to the Remittance service RetrieveAccountHolder HTTP endpoint.
func RetrieveAccountHolderRemittancePath(accountHolderIDType string, accountHolderID string) string {
	return fmt.Sprintf("/remittance/v1_0/accountholder/%v/%v/active", accountHolderIDType, accountHolderID)
}

// TransferRemittancePath returns the URL path to the Remittance service Transfer HTTP endpoint.
func TransferRemittancePath() string {
	return "/remittance/v1_0/transfer"
}

// TransferStatusRemittancePath returns the URL path to the Remittance service TransferStatus HTTP endpoint.
func TransferStatusRemittancePath(referenceID string) string {
	return fmt.Sprintf("/remittance/v1_0/transfer/%v", referenceID)
}
