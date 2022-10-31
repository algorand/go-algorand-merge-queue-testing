// Package private provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/algorand/oapi-codegen DO NOT EDIT.
package private

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/algorand/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Aborts a catchpoint catchup.
	// (DELETE /v2/catchup/{catchpoint})
	AbortCatchup(ctx echo.Context, catchpoint string) error
	// Starts a catchpoint catchup.
	// (POST /v2/catchup/{catchpoint})
	StartCatchup(ctx echo.Context, catchpoint string) error
	// Return a list of participation keys
	// (GET /v2/participation)
	GetParticipationKeys(ctx echo.Context) error
	// Add a participation key to the node
	// (POST /v2/participation)
	AddParticipationKey(ctx echo.Context) error
	// Delete a given participation key by ID
	// (DELETE /v2/participation/{participation-id})
	DeleteParticipationKeyByID(ctx echo.Context, participationId string) error
	// Get participation key info given a participation ID
	// (GET /v2/participation/{participation-id})
	GetParticipationKeyByID(ctx echo.Context, participationId string) error
	// Append state proof keys to a participation key
	// (POST /v2/participation/{participation-id})
	AppendKeys(ctx echo.Context, participationId string) error

	// (POST /v2/shutdown)
	ShutdownNode(ctx echo.Context, params ShutdownNodeParams) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// AbortCatchup converts echo context to params.
func (w *ServerInterfaceWrapper) AbortCatchup(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "catchpoint" -------------
	var catchpoint string

	err = runtime.BindStyledParameterWithLocation("simple", false, "catchpoint", runtime.ParamLocationPath, ctx.Param("catchpoint"), &catchpoint)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter catchpoint: %s", err))
	}

	ctx.Set(Api_keyScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AbortCatchup(ctx, catchpoint)
	return err
}

// StartCatchup converts echo context to params.
func (w *ServerInterfaceWrapper) StartCatchup(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "catchpoint" -------------
	var catchpoint string

	err = runtime.BindStyledParameterWithLocation("simple", false, "catchpoint", runtime.ParamLocationPath, ctx.Param("catchpoint"), &catchpoint)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter catchpoint: %s", err))
	}

	ctx.Set(Api_keyScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.StartCatchup(ctx, catchpoint)
	return err
}

// GetParticipationKeys converts echo context to params.
func (w *ServerInterfaceWrapper) GetParticipationKeys(ctx echo.Context) error {
	var err error

	ctx.Set(Api_keyScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetParticipationKeys(ctx)
	return err
}

// AddParticipationKey converts echo context to params.
func (w *ServerInterfaceWrapper) AddParticipationKey(ctx echo.Context) error {
	var err error

	ctx.Set(Api_keyScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AddParticipationKey(ctx)
	return err
}

// DeleteParticipationKeyByID converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteParticipationKeyByID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "participation-id" -------------
	var participationId string

	err = runtime.BindStyledParameterWithLocation("simple", false, "participation-id", runtime.ParamLocationPath, ctx.Param("participation-id"), &participationId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter participation-id: %s", err))
	}

	ctx.Set(Api_keyScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeleteParticipationKeyByID(ctx, participationId)
	return err
}

// GetParticipationKeyByID converts echo context to params.
func (w *ServerInterfaceWrapper) GetParticipationKeyByID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "participation-id" -------------
	var participationId string

	err = runtime.BindStyledParameterWithLocation("simple", false, "participation-id", runtime.ParamLocationPath, ctx.Param("participation-id"), &participationId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter participation-id: %s", err))
	}

	ctx.Set(Api_keyScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetParticipationKeyByID(ctx, participationId)
	return err
}

// AppendKeys converts echo context to params.
func (w *ServerInterfaceWrapper) AppendKeys(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "participation-id" -------------
	var participationId string

	err = runtime.BindStyledParameterWithLocation("simple", false, "participation-id", runtime.ParamLocationPath, ctx.Param("participation-id"), &participationId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter participation-id: %s", err))
	}

	ctx.Set(Api_keyScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AppendKeys(ctx, participationId)
	return err
}

// ShutdownNode converts echo context to params.
func (w *ServerInterfaceWrapper) ShutdownNode(ctx echo.Context) error {
	var err error

	ctx.Set(Api_keyScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params ShutdownNodeParams
	// ------------- Optional query parameter "timeout" -------------

	err = runtime.BindQueryParameter("form", true, false, "timeout", ctx.QueryParams(), &params.Timeout)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter timeout: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ShutdownNode(ctx, params)
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

	router.DELETE(baseURL+"/v2/catchup/:catchpoint", wrapper.AbortCatchup, m...)
	router.POST(baseURL+"/v2/catchup/:catchpoint", wrapper.StartCatchup, m...)
	router.GET(baseURL+"/v2/participation", wrapper.GetParticipationKeys, m...)
	router.POST(baseURL+"/v2/participation", wrapper.AddParticipationKey, m...)
	router.DELETE(baseURL+"/v2/participation/:participation-id", wrapper.DeleteParticipationKeyByID, m...)
	router.GET(baseURL+"/v2/participation/:participation-id", wrapper.GetParticipationKeyByID, m...)
	router.POST(baseURL+"/v2/participation/:participation-id", wrapper.AppendKeys, m...)
	router.POST(baseURL+"/v2/shutdown", wrapper.ShutdownNode, m...)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
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
