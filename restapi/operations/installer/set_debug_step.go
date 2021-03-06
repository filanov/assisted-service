// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// SetDebugStepHandlerFunc turns a function with the right signature into a set debug step handler
type SetDebugStepHandlerFunc func(SetDebugStepParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn SetDebugStepHandlerFunc) Handle(params SetDebugStepParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// SetDebugStepHandler interface for that can handle valid set debug step params
type SetDebugStepHandler interface {
	Handle(SetDebugStepParams, interface{}) middleware.Responder
}

// NewSetDebugStep creates a new http.Handler for the set debug step operation
func NewSetDebugStep(ctx *middleware.Context, handler SetDebugStepHandler) *SetDebugStep {
	return &SetDebugStep{Context: ctx, Handler: handler}
}

/*SetDebugStep swagger:route POST /clusters/{cluster_id}/hosts/{host_id}/actions/debug installer setDebugStep

Sets a single shot debug step that will be sent next time the host agent will ask for a command.

*/
type SetDebugStep struct {
	Context *middleware.Context
	Handler SetDebugStepHandler
}

func (o *SetDebugStep) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewSetDebugStepParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
