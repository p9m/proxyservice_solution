// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	errors "github.com/go-openapi/errors"
	loads "github.com/go-openapi/loads"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	security "github.com/go-openapi/runtime/security"
	spec "github.com/go-openapi/spec"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"ProxyService/proxyservice/openapi/autogen/proxyservice/server/operations/rpc"
)

// NewProxyServiceAPI creates a new ProxyService instance
func NewProxyServiceAPI(spec *loads.Document) *ProxyServiceAPI {
	return &ProxyServiceAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		ServerShutdown:      func() {},
		spec:                spec,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,
		JSONConsumer:        runtime.JSONConsumer(),
		JSONProducer:        runtime.JSONProducer(),
		RPCGetClientsHandler: rpc.GetClientsHandlerFunc(func(params rpc.GetClientsParams) middleware.Responder {
			return middleware.NotImplemented("operation RPCGetClients has not yet been implemented")
		}),
		RPCGetInvoicesHandler: rpc.GetInvoicesHandlerFunc(func(params rpc.GetInvoicesParams) middleware.Responder {
			return middleware.NotImplemented("operation RPCGetInvoices has not yet been implemented")
		}),
	}
}

/*ProxyServiceAPI A proxy service acquiring information from RabbitMQ RPC Queue */
type ProxyServiceAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for a "application/json" mime type
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for a "application/json" mime type
	JSONProducer runtime.Producer

	// RPCGetClientsHandler sets the operation handler for the get clients operation
	RPCGetClientsHandler rpc.GetClientsHandler
	// RPCGetInvoicesHandler sets the operation handler for the get invoices operation
	RPCGetInvoicesHandler rpc.GetInvoicesHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *ProxyServiceAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *ProxyServiceAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *ProxyServiceAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *ProxyServiceAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *ProxyServiceAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *ProxyServiceAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *ProxyServiceAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the ProxyServiceAPI
func (o *ProxyServiceAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.RPCGetClientsHandler == nil {
		unregistered = append(unregistered, "rpc.GetClientsHandler")
	}

	if o.RPCGetInvoicesHandler == nil {
		unregistered = append(unregistered, "rpc.GetInvoicesHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *ProxyServiceAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *ProxyServiceAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {

	return nil

}

// Authorizer returns the registered authorizer
func (o *ProxyServiceAPI) Authorizer() runtime.Authorizer {

	return nil

}

// ConsumersFor gets the consumers for the specified media types
func (o *ProxyServiceAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {

	result := make(map[string]runtime.Consumer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONConsumer

		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result

}

// ProducersFor gets the producers for the specified media types
func (o *ProxyServiceAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {

	result := make(map[string]runtime.Producer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONProducer

		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result

}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *ProxyServiceAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the proxy service API
func (o *ProxyServiceAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *ProxyServiceAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/clients"] = rpc.NewGetClients(o.context, o.RPCGetClientsHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/invoices/{client_id}"] = rpc.NewGetInvoices(o.context, o.RPCGetInvoicesHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *ProxyServiceAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *ProxyServiceAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *ProxyServiceAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *ProxyServiceAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}
