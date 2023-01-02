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
	"3wAA",
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