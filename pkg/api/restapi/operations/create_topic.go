// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"

	models "github.com/dairlair/tweetwatch/pkg/api/models"
)

// CreateTopicHandlerFunc turns a function with the right signature into a create topic handler
type CreateTopicHandlerFunc func(CreateTopicParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateTopicHandlerFunc) Handle(params CreateTopicParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// CreateTopicHandler interface for that can handle valid create topic params
type CreateTopicHandler interface {
	Handle(CreateTopicParams, *models.User) middleware.Responder
}

// NewCreateTopic creates a new http.Handler for the create topic operation
func NewCreateTopic(ctx *middleware.Context, handler CreateTopicHandler) *CreateTopic {
	return &CreateTopic{Context: ctx, Handler: handler}
}

/*CreateTopic swagger:route POST /topics createTopic

CreateTopic create topic API

*/
type CreateTopic struct {
	Context *middleware.Context
	Handler CreateTopicHandler
}

func (o *CreateTopic) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateTopicParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal *models.User
	if uprinc != nil {
		principal = uprinc.(*models.User) // this is really a models.User, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
