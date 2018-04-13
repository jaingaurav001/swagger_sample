// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetSearchHandlerFunc turns a function with the right signature into a get search handler
type GetSearchHandlerFunc func(GetSearchParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetSearchHandlerFunc) Handle(params GetSearchParams) middleware.Responder {
	return fn(params)
}

// GetSearchHandler interface for that can handle valid get search params
type GetSearchHandler interface {
	Handle(GetSearchParams) middleware.Responder
}

// NewGetSearch creates a new http.Handler for the get search operation
func NewGetSearch(ctx *middleware.Context, handler GetSearchHandler) *GetSearch {
	return &GetSearch{Context: ctx, Handler: handler}
}

/*GetSearch swagger:route GET /search user getSearch

get user information

*/
type GetSearch struct {
	Context *middleware.Context
	Handler GetSearchHandler
}

func (o *GetSearch) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetSearchParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
