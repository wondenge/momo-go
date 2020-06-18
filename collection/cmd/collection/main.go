package main

import (
	"context"
	"flag"
	"fmt"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"sync"

	"github.com/go-kit/kit/log"
	collection "github.com/wondenge/momo-go/collection"
	getbalance "github.com/wondenge/momo-go/collection/gen/getbalance"
	requestopay "github.com/wondenge/momo-go/collection/gen/requestopay"
	token "github.com/wondenge/momo-go/collection/gen/token"
	validate "github.com/wondenge/momo-go/collection/gen/validate"
)

func main() {
	// Define command line flags, add any other flag required to configure the
	// service.
	var (
		hostF     = flag.String("host", "local", "Server host (valid values: local, development)")
		domainF   = flag.String("domain", "", "Host domain name (overrides host domain specified in service design)")
		httpPortF = flag.String("http-port", "", "HTTP port (overrides host HTTP port specified in service design)")
		secureF   = flag.Bool("secure", false, "Use secure scheme (https or grpcs)")
		dbgF      = flag.Bool("debug", false, "Log request and response bodies")
	)
	flag.Parse()

	// Setup gokit logger.
	var (
		logger log.Logger
	)
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	// Initialize the services.
	var (
		tokenSvc       token.Service
		validateSvc    validate.Service
		getbalanceSvc  getbalance.Service
		requestopaySvc requestopay.Service
	)
	{
		tokenSvc = collection.NewToken(logger)
		validateSvc = collection.NewValidate(logger)
		getbalanceSvc = collection.NewGetbalance(logger)
		requestopaySvc = collection.NewRequestopay(logger)
	}

	// Wrap the services in endpoints that can be invoked from other services
	// potentially running in different processes.
	var (
		tokenEndpoints       *token.Endpoints
		validateEndpoints    *validate.Endpoints
		getbalanceEndpoints  *getbalance.Endpoints
		requestopayEndpoints *requestopay.Endpoints
	)
	{
		tokenEndpoints = token.NewEndpoints(tokenSvc)
		validateEndpoints = validate.NewEndpoints(validateSvc)
		getbalanceEndpoints = getbalance.NewEndpoints(getbalanceSvc)
		requestopayEndpoints = requestopay.NewEndpoints(requestopaySvc)
	}

	// Create channel used by both the signal handler and server goroutines
	// to notify the main goroutine when to stop the server.
	errc := make(chan error)

	// Setup interrupt handler. This optional step configures the process so
	// that SIGINT and SIGTERM signals cause the services to stop gracefully.
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		errc <- fmt.Errorf("%s", <-c)
	}()

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	// Start the servers and send errors (if any) to the error channel.
	switch *hostF {
	case "local":
		{
			addr := "http://localhost:3000"
			u, err := url.Parse(addr)
			if err != nil {
				fmt.Fprintf(os.Stderr, "invalid URL %#v: %s\n", addr, err)
				os.Exit(1)
			}
			if *secureF {
				u.Scheme = "https"
			}
			if *domainF != "" {
				u.Host = *domainF
			}
			if *httpPortF != "" {
				h := strings.Split(u.Host, ":")[0]
				u.Host = h + ":" + *httpPortF
			} else if u.Port() == "" {
				u.Host += ":80"
			}
			handleHTTPServer(ctx, u, tokenEndpoints, validateEndpoints, getbalanceEndpoints, requestopayEndpoints, &wg, errc, logger, *dbgF)
		}

	case "development":
		{
			addr := "https://sandbox.momodeveloper.mtn.com"
			u, err := url.Parse(addr)
			if err != nil {
				fmt.Fprintf(os.Stderr, "invalid URL %#v: %s\n", addr, err)
				os.Exit(1)
			}
			if *secureF {
				u.Scheme = "https"
			}
			if *domainF != "" {
				u.Host = *domainF
			}
			if *httpPortF != "" {
				h := strings.Split(u.Host, ":")[0]
				u.Host = h + ":" + *httpPortF
			} else if u.Port() == "" {
				u.Host += ":443"
			}
			handleHTTPServer(ctx, u, tokenEndpoints, validateEndpoints, getbalanceEndpoints, requestopayEndpoints, &wg, errc, logger, *dbgF)
		}

	default:
		fmt.Fprintf(os.Stderr, "invalid host argument: %q (valid hosts: local|development)\n", *hostF)
	}

	// Wait for signal.
	logger.Log("info", fmt.Sprintf("exiting (%v)", <-errc))

	// Send cancellation signal to the goroutines.
	cancel()

	wg.Wait()
	logger.Log("info", fmt.Sprintf("exited"))
}