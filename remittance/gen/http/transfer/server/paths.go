// Code generated by goa v3.1.2, DO NOT EDIT.
//
// HTTP request path constructors for the transfer service.
//
// Command:
// $ goa gen github.com/wondenge/momo-go/remittance/design

package server

import (
	"fmt"
)

// CreateTransferPath returns the URL path to the transfer service create HTTP endpoint.
func CreateTransferPath() string {
	return "/remittance/v1_0/transfer"
}

// GetTransferPath returns the URL path to the transfer service get HTTP endpoint.
func GetTransferPath(referenceID string) string {
	return fmt.Sprintf("/remittance/v1_0/transfer/%v", referenceID)
}
