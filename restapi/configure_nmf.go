// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	swag "github.com/go-openapi/swag"
	cors "github.com/rs/cors"

	"github.com/pl0q1n/No_More_Flex/transactions"
	"github.com/pl0q1n/No_More_Flex/models"
	"github.com/pl0q1n/No_More_Flex/restapi/operations"
)

var transactionsStorage transactions.Storage = transactions.NewMemoryStorage()

//go:generate swagger generate server --target ../../nmf --name Nmf --spec ../swagger.yml

func configureFlags(api *operations.NmfAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.NmfAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.AddTransactionHandler = operations.AddTransactionHandlerFunc(func(params operations.AddTransactionParams) middleware.Responder {
		if err := transactionsStorage.Add(transactions.UserID(0), *params.Body); err != nil {
			return operations.NewAddTransactionDefault(500).WithPayload(&models.Error{Message: swag.String(err.Error())})
		}
		return operations.NewAddTransactionCreated()
	})

	api.FilterTransactionsHandler = operations.FilterTransactionsHandlerFunc(func(params operations.FilterTransactionsParams) middleware.Responder {
		filter := transactions.NewFilter()
		if params.Category != nil {
			filter.Category = *params.Category
		}

		if params.From != nil {
			filter.From = *params.From
		}

		if params.To != nil {
			filter.To = *params.To
		}

		if params.Sender != nil {
			filter.Sender = *params.Sender
		}

		if params.Receiver != nil {
			filter.Receiver = *params.Receiver
		}

		transactions, err := transactionsStorage.Filter(transactions.UserID(0), filter)
		if err != nil {
			return operations.NewFilterTransactionsDefault(500).WithPayload(&models.Error{Message: swag.String(err.Error())})
		}

		return operations.NewFilterTransactionsOK().WithPayload(transactions)
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	handleCORS := cors.Default().Handler
	return handleCORS(handler)
}
