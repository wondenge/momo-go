// Code generated by goa v3.1.2, DO NOT EDIT.
//
// remittance HTTP client CLI support package
//
// Command:
// $ goa gen github.com/wondenge/momo-go/remittance/design

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/go-kit/kit/endpoint"
	getbalancec "github.com/wondenge/momo-go/remittance/gen/http/getbalance/client"
	tokenc "github.com/wondenge/momo-go/remittance/gen/http/token/client"
	transferc "github.com/wondenge/momo-go/remittance/gen/http/transfer/client"
	validatec "github.com/wondenge/momo-go/remittance/gen/http/validate/client"
	goahttp "goa.design/goa/v3/http"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `getbalance show
token create
transfer (create|get)
validate show
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` getbalance show --authorization "Consectetur sed nobis commodi exercitationem adipisci voluptatem." --x-target-environment "Tempora consequuntur." --ocp-apim-subscription-key "Non dolor assumenda ipsum."` + "\n" +
		os.Args[0] + ` token create --ocp-apim-subscription-key "Illum reprehenderit itaque quis exercitationem sint." --api-key "Non recusandae vel ut voluptatem molestiae." --api-secret "Totam officia placeat dolor quaerat facilis."` + "\n" +
		os.Args[0] + ` transfer create --authorization "Aut consequuntur eius." --x-callback-url "Cupiditate quidem a." --x-reference-id "Magni ut." --x-target-environment "Aut id autem omnis ad." --ocp-apim-subscription-key "Provident itaque."` + "\n" +
		os.Args[0] + ` validate show --account-holder-id-type "Nostrum sit est magnam dolorem." --account-holder-id "Iste harum repellendus sit." --authorization "Voluptatibus voluptatum nemo repellendus voluptas aliquid non." --x-target-environment "Ea voluptatem dolores expedita esse sit praesentium." --ocp-apim-subscription-key "Repellendus quas illum optio."` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (endpoint.Endpoint, interface{}, error) {
	var (
		getbalanceFlags = flag.NewFlagSet("getbalance", flag.ContinueOnError)

		getbalanceShowFlags                      = flag.NewFlagSet("show", flag.ExitOnError)
		getbalanceShowAuthorizationFlag          = getbalanceShowFlags.String("authorization", "", "")
		getbalanceShowXTargetEnvironmentFlag     = getbalanceShowFlags.String("x-target-environment", "REQUIRED", "")
		getbalanceShowOcpApimSubscriptionKeyFlag = getbalanceShowFlags.String("ocp-apim-subscription-key", "REQUIRED", "")

		tokenFlags = flag.NewFlagSet("token", flag.ContinueOnError)

		tokenCreateFlags                      = flag.NewFlagSet("create", flag.ExitOnError)
		tokenCreateOcpApimSubscriptionKeyFlag = tokenCreateFlags.String("ocp-apim-subscription-key", "REQUIRED", "")
		tokenCreateAPIKeyFlag                 = tokenCreateFlags.String("api-key", "REQUIRED", "API Key")
		tokenCreateAPISecretFlag              = tokenCreateFlags.String("api-secret", "REQUIRED", "API Secret")

		transferFlags = flag.NewFlagSet("transfer", flag.ContinueOnError)

		transferCreateFlags                      = flag.NewFlagSet("create", flag.ExitOnError)
		transferCreateAuthorizationFlag          = transferCreateFlags.String("authorization", "", "")
		transferCreateXCallbackURLFlag           = transferCreateFlags.String("x-callback-url", "", "")
		transferCreateXReferenceIDFlag           = transferCreateFlags.String("x-reference-id", "REQUIRED", "")
		transferCreateXTargetEnvironmentFlag     = transferCreateFlags.String("x-target-environment", "REQUIRED", "")
		transferCreateOcpApimSubscriptionKeyFlag = transferCreateFlags.String("ocp-apim-subscription-key", "REQUIRED", "")

		transferGetFlags                      = flag.NewFlagSet("get", flag.ExitOnError)
		transferGetReferenceIDFlag            = transferGetFlags.String("reference-id", "REQUIRED", " UUID of transaction to get result")
		transferGetAuthorizationFlag          = transferGetFlags.String("authorization", "", "")
		transferGetXTargetEnvironmentFlag     = transferGetFlags.String("x-target-environment", "REQUIRED", "")
		transferGetOcpApimSubscriptionKeyFlag = transferGetFlags.String("ocp-apim-subscription-key", "REQUIRED", "")

		validateFlags = flag.NewFlagSet("validate", flag.ContinueOnError)

		validateShowFlags                      = flag.NewFlagSet("show", flag.ExitOnError)
		validateShowAccountHolderIDTypeFlag    = validateShowFlags.String("account-holder-id-type", "REQUIRED", "Specifies the type of the party ID")
		validateShowAccountHolderIDFlag        = validateShowFlags.String("account-holder-id", "REQUIRED", "The party number.")
		validateShowAuthorizationFlag          = validateShowFlags.String("authorization", "", "")
		validateShowXTargetEnvironmentFlag     = validateShowFlags.String("x-target-environment", "REQUIRED", "")
		validateShowOcpApimSubscriptionKeyFlag = validateShowFlags.String("ocp-apim-subscription-key", "REQUIRED", "")
	)
	getbalanceFlags.Usage = getbalanceUsage
	getbalanceShowFlags.Usage = getbalanceShowUsage

	tokenFlags.Usage = tokenUsage
	tokenCreateFlags.Usage = tokenCreateUsage

	transferFlags.Usage = transferUsage
	transferCreateFlags.Usage = transferCreateUsage
	transferGetFlags.Usage = transferGetUsage

	validateFlags.Usage = validateUsage
	validateShowFlags.Usage = validateShowUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "getbalance":
			svcf = getbalanceFlags
		case "token":
			svcf = tokenFlags
		case "transfer":
			svcf = transferFlags
		case "validate":
			svcf = validateFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "getbalance":
			switch epn {
			case "show":
				epf = getbalanceShowFlags

			}

		case "token":
			switch epn {
			case "create":
				epf = tokenCreateFlags

			}

		case "transfer":
			switch epn {
			case "create":
				epf = transferCreateFlags

			case "get":
				epf = transferGetFlags

			}

		case "validate":
			switch epn {
			case "show":
				epf = validateShowFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint endpoint.Endpoint
		err      error
	)
	{
		switch svcn {
		case "getbalance":
			c := getbalancec.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "show":
				endpoint = c.Show()
				data, err = getbalancec.BuildShowPayload(*getbalanceShowAuthorizationFlag, *getbalanceShowXTargetEnvironmentFlag, *getbalanceShowOcpApimSubscriptionKeyFlag)
			}
		case "token":
			c := tokenc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "create":
				endpoint = c.Create()
				data, err = tokenc.BuildCreatePayload(*tokenCreateOcpApimSubscriptionKeyFlag, *tokenCreateAPIKeyFlag, *tokenCreateAPISecretFlag)
			}
		case "transfer":
			c := transferc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "create":
				endpoint = c.Create()
				data, err = transferc.BuildCreatePayload(*transferCreateAuthorizationFlag, *transferCreateXCallbackURLFlag, *transferCreateXReferenceIDFlag, *transferCreateXTargetEnvironmentFlag, *transferCreateOcpApimSubscriptionKeyFlag)
			case "get":
				endpoint = c.Get()
				data, err = transferc.BuildGetPayload(*transferGetReferenceIDFlag, *transferGetAuthorizationFlag, *transferGetXTargetEnvironmentFlag, *transferGetOcpApimSubscriptionKeyFlag)
			}
		case "validate":
			c := validatec.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "show":
				endpoint = c.Show()
				data, err = validatec.BuildShowPayload(*validateShowAccountHolderIDTypeFlag, *validateShowAccountHolderIDFlag, *validateShowAuthorizationFlag, *validateShowXTargetEnvironmentFlag, *validateShowOcpApimSubscriptionKeyFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// getbalanceUsage displays the usage of the getbalance command and its
// subcommands.
func getbalanceUsage() {
	fmt.Fprintf(os.Stderr, `Service is the getbalance service interface.
Usage:
    %s [globalflags] getbalance COMMAND [flags]

COMMAND:
    show: Get the balance of the account

Additional help:
    %s getbalance COMMAND --help
`, os.Args[0], os.Args[0])
}
func getbalanceShowUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] getbalance show -authorization STRING -x-target-environment STRING -ocp-apim-subscription-key STRING

Get the balance of the account
    -authorization STRING: 
    -x-target-environment STRING: 
    -ocp-apim-subscription-key STRING: 

Example:
    `+os.Args[0]+` getbalance show --authorization "Consectetur sed nobis commodi exercitationem adipisci voluptatem." --x-target-environment "Tempora consequuntur." --ocp-apim-subscription-key "Non dolor assumenda ipsum."
`, os.Args[0])
}

// tokenUsage displays the usage of the token command and its subcommands.
func tokenUsage() {
	fmt.Fprintf(os.Stderr, `Service is the token service interface.
Usage:
    %s [globalflags] token COMMAND [flags]

COMMAND:
    create: Creates an Access Token.

Additional help:
    %s token COMMAND --help
`, os.Args[0], os.Args[0])
}
func tokenCreateUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] token create -ocp-apim-subscription-key STRING -api-key STRING -api-secret STRING

Creates an Access Token.
    -ocp-apim-subscription-key STRING: 
    -api-key STRING: API Key
    -api-secret STRING: API Secret

Example:
    `+os.Args[0]+` token create --ocp-apim-subscription-key "Illum reprehenderit itaque quis exercitationem sint." --api-key "Non recusandae vel ut voluptatem molestiae." --api-secret "Totam officia placeat dolor quaerat facilis."
`, os.Args[0])
}

// transferUsage displays the usage of the transfer command and its subcommands.
func transferUsage() {
	fmt.Fprintf(os.Stderr, `Service is the transfer service interface.
Usage:
    %s [globalflags] transfer COMMAND [flags]

COMMAND:
    create: Transfers an amount from the owner’s account to a payee account
    get: This operation is used to get the status of a transfer.

Additional help:
    %s transfer COMMAND --help
`, os.Args[0], os.Args[0])
}
func transferCreateUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] transfer create -authorization STRING -x-callback-url STRING -x-reference-id STRING -x-target-environment STRING -ocp-apim-subscription-key STRING

Transfers an amount from the owner’s account to a payee account
    -authorization STRING: 
    -x-callback-url STRING: 
    -x-reference-id STRING: 
    -x-target-environment STRING: 
    -ocp-apim-subscription-key STRING: 

Example:
    `+os.Args[0]+` transfer create --authorization "Aut consequuntur eius." --x-callback-url "Cupiditate quidem a." --x-reference-id "Magni ut." --x-target-environment "Aut id autem omnis ad." --ocp-apim-subscription-key "Provident itaque."
`, os.Args[0])
}

func transferGetUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] transfer get -reference-id STRING -authorization STRING -x-target-environment STRING -ocp-apim-subscription-key STRING

This operation is used to get the status of a transfer.
    -reference-id STRING:  UUID of transaction to get result
    -authorization STRING: 
    -x-target-environment STRING: 
    -ocp-apim-subscription-key STRING: 

Example:
    `+os.Args[0]+` transfer get --reference-id "Commodi autem voluptatibus sunt." --authorization "Eaque distinctio inventore dicta earum sit." --x-target-environment "Voluptatem omnis aut nihil sequi quasi." --ocp-apim-subscription-key "Voluptatum culpa."
`, os.Args[0])
}

// validateUsage displays the usage of the validate command and its subcommands.
func validateUsage() {
	fmt.Fprintf(os.Stderr, `Service is the validate service interface.
Usage:
    %s [globalflags] validate COMMAND [flags]

COMMAND:
    show: Checks if an account holder is registered and active in the system

Additional help:
    %s validate COMMAND --help
`, os.Args[0], os.Args[0])
}
func validateShowUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] validate show -account-holder-id-type STRING -account-holder-id STRING -authorization STRING -x-target-environment STRING -ocp-apim-subscription-key STRING

Checks if an account holder is registered and active in the system
    -account-holder-id-type STRING: Specifies the type of the party ID
    -account-holder-id STRING: The party number.
    -authorization STRING: 
    -x-target-environment STRING: 
    -ocp-apim-subscription-key STRING: 

Example:
    `+os.Args[0]+` validate show --account-holder-id-type "Nostrum sit est magnam dolorem." --account-holder-id "Iste harum repellendus sit." --authorization "Voluptatibus voluptatum nemo repellendus voluptas aliquid non." --x-target-environment "Ea voluptatem dolores expedita esse sit praesentium." --ocp-apim-subscription-key "Repellendus quas illum optio."
`, os.Args[0])
}