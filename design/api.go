package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

// API describes the global properties of the API server.
var _ = API("momo", func() {
	Title("MTN MoMo APIs")
	Description("HTTP Service for accessing MTN's MoMo APIs.")
	Version("1.0")
	TermsOfService("https://github.com/wondenge/momo-go/blob/master/LICENCE")
	Contact(func() {
		Name("William Ondenge")
		Email("ondengew@gmail.com")
		URL("https://www.ondenge.me")
	})
	License(func() {
		Name("Apache License")
		URL("https://github.com/wondenge/momo-go/blob/master/LICENCE")
	})
	Docs(func() {
		Description("Library Usage")
		URL("https://github.com/wondenge/momo-go/blob/master/README.md")
	})
	Server("momo", func() {
		Description("momo hosts all MTN's MoMo Services.")
		Services("collection", "disbursement", "remittance", "user")
		Host("local", func() {
			Description("Localhost")
			URI("http://localhost:3000")
		})
		Host("development", func() {
			Description("Development hosts.")
			URI("https://sandbox.momodeveloper.mtn.com")
		})
	})
})