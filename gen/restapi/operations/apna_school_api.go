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
)

// NewApnaSchoolAPI creates a new ApnaSchool instance
func NewApnaSchoolAPI(spec *loads.Document) *ApnaSchoolAPI {
	return &ApnaSchoolAPI{
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
		SignUpStudentHandler: SignUpStudentHandlerFunc(func(params SignUpStudentParams) middleware.Responder {
			return middleware.NotImplemented("operation SignUpStudent has not yet been implemented")
		}),
		SignUpTeacherHandler: SignUpTeacherHandlerFunc(func(params SignUpTeacherParams) middleware.Responder {
			return middleware.NotImplemented("operation SignUpTeacher has not yet been implemented")
		}),
		AddStudentHandler: AddStudentHandlerFunc(func(params AddStudentParams) middleware.Responder {
			return middleware.NotImplemented("operation AddStudent has not yet been implemented")
		}),
		AddTeacherHandler: AddTeacherHandlerFunc(func(params AddTeacherParams) middleware.Responder {
			return middleware.NotImplemented("operation AddTeacher has not yet been implemented")
		}),
		DeleteStudentHandler: DeleteStudentHandlerFunc(func(params DeleteStudentParams) middleware.Responder {
			return middleware.NotImplemented("operation DeleteStudent has not yet been implemented")
		}),
		DeleteTeacherHandler: DeleteTeacherHandlerFunc(func(params DeleteTeacherParams) middleware.Responder {
			return middleware.NotImplemented("operation DeleteTeacher has not yet been implemented")
		}),
		EditStudentHandler: EditStudentHandlerFunc(func(params EditStudentParams) middleware.Responder {
			return middleware.NotImplemented("operation EditStudent has not yet been implemented")
		}),
		EditTeacherHandler: EditTeacherHandlerFunc(func(params EditTeacherParams) middleware.Responder {
			return middleware.NotImplemented("operation EditTeacher has not yet been implemented")
		}),
		GetStudentHandler: GetStudentHandlerFunc(func(params GetStudentParams) middleware.Responder {
			return middleware.NotImplemented("operation GetStudent has not yet been implemented")
		}),
		GetTeacherHandler: GetTeacherHandlerFunc(func(params GetTeacherParams) middleware.Responder {
			return middleware.NotImplemented("operation GetTeacher has not yet been implemented")
		}),
	}
}

/*ApnaSchoolAPI A School Management System */
type ApnaSchoolAPI struct {
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
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for a "application/json" mime type
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for a "application/json" mime type
	JSONProducer runtime.Producer

	// SignUpStudentHandler sets the operation handler for the sign up student operation
	SignUpStudentHandler SignUpStudentHandler
	// SignUpTeacherHandler sets the operation handler for the sign up teacher operation
	SignUpTeacherHandler SignUpTeacherHandler
	// AddStudentHandler sets the operation handler for the add student operation
	AddStudentHandler AddStudentHandler
	// AddTeacherHandler sets the operation handler for the add teacher operation
	AddTeacherHandler AddTeacherHandler
	// DeleteStudentHandler sets the operation handler for the delete student operation
	DeleteStudentHandler DeleteStudentHandler
	// DeleteTeacherHandler sets the operation handler for the delete teacher operation
	DeleteTeacherHandler DeleteTeacherHandler
	// EditStudentHandler sets the operation handler for the edit student operation
	EditStudentHandler EditStudentHandler
	// EditTeacherHandler sets the operation handler for the edit teacher operation
	EditTeacherHandler EditTeacherHandler
	// GetStudentHandler sets the operation handler for the get student operation
	GetStudentHandler GetStudentHandler
	// GetTeacherHandler sets the operation handler for the get teacher operation
	GetTeacherHandler GetTeacherHandler

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
func (o *ApnaSchoolAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *ApnaSchoolAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *ApnaSchoolAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *ApnaSchoolAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *ApnaSchoolAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *ApnaSchoolAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *ApnaSchoolAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the ApnaSchoolAPI
func (o *ApnaSchoolAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.SignUpStudentHandler == nil {
		unregistered = append(unregistered, "SignUpStudentHandler")
	}

	if o.SignUpTeacherHandler == nil {
		unregistered = append(unregistered, "SignUpTeacherHandler")
	}

	if o.AddStudentHandler == nil {
		unregistered = append(unregistered, "AddStudentHandler")
	}

	if o.AddTeacherHandler == nil {
		unregistered = append(unregistered, "AddTeacherHandler")
	}

	if o.DeleteStudentHandler == nil {
		unregistered = append(unregistered, "DeleteStudentHandler")
	}

	if o.DeleteTeacherHandler == nil {
		unregistered = append(unregistered, "DeleteTeacherHandler")
	}

	if o.EditStudentHandler == nil {
		unregistered = append(unregistered, "EditStudentHandler")
	}

	if o.EditTeacherHandler == nil {
		unregistered = append(unregistered, "EditTeacherHandler")
	}

	if o.GetStudentHandler == nil {
		unregistered = append(unregistered, "GetStudentHandler")
	}

	if o.GetTeacherHandler == nil {
		unregistered = append(unregistered, "GetTeacherHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *ApnaSchoolAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *ApnaSchoolAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {

	return nil

}

// Authorizer returns the registered authorizer
func (o *ApnaSchoolAPI) Authorizer() runtime.Authorizer {

	return nil

}

// ConsumersFor gets the consumers for the specified media types
func (o *ApnaSchoolAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {

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
func (o *ApnaSchoolAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {

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
func (o *ApnaSchoolAPI) HandlerFor(method, path string) (http.Handler, bool) {
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

// Context returns the middleware context for the apna school API
func (o *ApnaSchoolAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *ApnaSchoolAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/student/signup"] = NewSignUpStudent(o.context, o.SignUpStudentHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/teacher/signup"] = NewSignUpTeacher(o.context, o.SignUpTeacherHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/student"] = NewAddStudent(o.context, o.AddStudentHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/teacher"] = NewAddTeacher(o.context, o.AddTeacherHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/student/{ID}"] = NewDeleteStudent(o.context, o.DeleteStudentHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/teacher/{ID}"] = NewDeleteTeacher(o.context, o.DeleteTeacherHandler)

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/student/{ID}"] = NewEditStudent(o.context, o.EditStudentHandler)

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/teacher/{ID}"] = NewEditTeacher(o.context, o.EditTeacherHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/student/{ID}"] = NewGetStudent(o.context, o.GetStudentHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/teacher/{ID}"] = NewGetTeacher(o.context, o.GetTeacherHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *ApnaSchoolAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *ApnaSchoolAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *ApnaSchoolAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *ApnaSchoolAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}
