// Code generated by goa v3.1.3, DO NOT EDIT.
//
// HTTP request path constructors for the User service.
//
// Command:
// $ goa gen github.com/wondenge/momo-go/design

package client

import (
	"fmt"
)

// NewUserUserPath returns the URL path to the User service NewUser HTTP endpoint.
func NewUserUserPath() string {
	return "/v1_0/apiuser"
}

// NewKeyUserPath returns the URL path to the User service NewKey HTTP endpoint.
func NewKeyUserPath() string {
	return "/v1_0/apiuser/{X-Reference-Id}/apikey"
}

// GetUserUserPath returns the URL path to the User service GetUser HTTP endpoint.
func GetUserUserPath() string {
	return "/v1_0/apiuser/{X-Reference-Id}"
}

// GetUserDetailsUserPath returns the URL path to the User service GetUserDetails HTTP endpoint.
func GetUserDetailsUserPath(aPIUser string) string {
	return fmt.Sprintf("/apiuser/%v", aPIUser)
}
