// Package public provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/algorand/oapi-codegen DO NOT EDIT.
package public

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	. "github.com/algorand/go-algorand/daemon/algod/api/server/v2/generated/model"
	"github.com/algorand/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get a list of unconfirmed transactions currently in the transaction pool by address.
	// (GET /v2/accounts/{address}/transactions/pending)
	GetPendingTransactionsByAddress(ctx echo.Context, address string, params GetPendingTransactionsByAddressParams) error
	// Broadcasts a raw transaction or transaction group to the network.
	// (POST /v2/transactions)
	RawTransaction(ctx echo.Context) error
	// Get a list of unconfirmed transactions currently in the transaction pool.
	// (GET /v2/transactions/pending)
	GetPendingTransactions(ctx echo.Context, params GetPendingTransactionsParams) error
	// Get a specific pending transaction.
	// (GET /v2/transactions/pending/{txid})
	PendingTransactionInformation(ctx echo.Context, txid string, params PendingTransactionInformationParams) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetPendingTransactionsByAddress converts echo context to params.
func (w *ServerInterfaceWrapper) GetPendingTransactionsByAddress(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "address" -------------
	var address string

	err = runtime.BindStyledParameterWithLocation("simple", false, "address", runtime.ParamLocationPath, ctx.Param("address"), &address)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter address: %s", err))
	}

	ctx.Set(Api_keyScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetPendingTransactionsByAddressParams
	// ------------- Optional query parameter "max" -------------

	err = runtime.BindQueryParameter("form", true, false, "max", ctx.QueryParams(), &params.Max)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter max: %s", err))
	}

	// ------------- Optional query parameter "format" -------------

	err = runtime.BindQueryParameter("form", true, false, "format", ctx.QueryParams(), &params.Format)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter format: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetPendingTransactionsByAddress(ctx, address, params)
	return err
}

// RawTransaction converts echo context to params.
func (w *ServerInterfaceWrapper) RawTransaction(ctx echo.Context) error {
	var err error

	ctx.Set(Api_keyScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.RawTransaction(ctx)
	return err
}

// GetPendingTransactions converts echo context to params.
func (w *ServerInterfaceWrapper) GetPendingTransactions(ctx echo.Context) error {
	var err error

	ctx.Set(Api_keyScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetPendingTransactionsParams
	// ------------- Optional query parameter "max" -------------

	err = runtime.BindQueryParameter("form", true, false, "max", ctx.QueryParams(), &params.Max)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter max: %s", err))
	}

	// ------------- Optional query parameter "format" -------------

	err = runtime.BindQueryParameter("form", true, false, "format", ctx.QueryParams(), &params.Format)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter format: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetPendingTransactions(ctx, params)
	return err
}

// PendingTransactionInformation converts echo context to params.
func (w *ServerInterfaceWrapper) PendingTransactionInformation(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "txid" -------------
	var txid string

	err = runtime.BindStyledParameterWithLocation("simple", false, "txid", runtime.ParamLocationPath, ctx.Param("txid"), &txid)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter txid: %s", err))
	}

	ctx.Set(Api_keyScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params PendingTransactionInformationParams
	// ------------- Optional query parameter "format" -------------

	err = runtime.BindQueryParameter("form", true, false, "format", ctx.QueryParams(), &params.Format)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter format: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PendingTransactionInformation(ctx, txid, params)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface, m ...echo.MiddlewareFunc) {
	RegisterHandlersWithBaseURL(router, si, "", m...)
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string, m ...echo.MiddlewareFunc) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/v2/accounts/:address/transactions/pending", wrapper.GetPendingTransactionsByAddress, m...)
	router.POST(baseURL+"/v2/transactions", wrapper.RawTransaction, m...)
	router.GET(baseURL+"/v2/transactions/pending", wrapper.GetPendingTransactions, m...)
	router.GET(baseURL+"/v2/transactions/pending/:txid", wrapper.PendingTransactionInformation, m...)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"0AAA",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
