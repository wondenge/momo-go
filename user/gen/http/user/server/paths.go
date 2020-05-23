// Code generated by goa v3.1.2, DO NOT EDIT.
//
// HTTP request path constructors for the user service.
//
// Command:
// $ goa gen github.com/wondenge/momo-go/user/design

package server

import (
	"fmt"
)

// CreateuserUserPath returns the URL path to the user service createuser HTTP endpoint.
func CreateuserUserPath() string {
	return "/v1_0/apiuser/"
}

// CreatekeyUserPath returns the URL path to the user service createkey HTTP endpoint.
func CreatekeyUserPath() string {
	return "/v1_0/apiuser/{X-Reference-Id}/apikey"
}

// ListUserPath returns the URL path to the user service list HTTP endpoint.
func ListUserPath() string {
	return "/v1_0/apiuser/{X-Reference-Id}"
}

// ShowUserPath returns the URL path to the user service show HTTP endpoint.
func ShowUserPath(aPIUser string) string {
	return fmt.Sprintf("/v1_0/apiuser/apiuser/%v", aPIUser)
}
