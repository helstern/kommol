package get

import (
	"github.com/gorilla/mux"
	"github.com/helstern/kommol/internal/core/object/app"
)

var (
	routeWebsite *RouteWebsite
	routeBucket  *RouteBucket
)

func Routes(r *mux.Router, s app.ObjectProxy) {

	h := NewOperation(s)

	if routeWebsite == nil {
		routeWebsite = &RouteWebsite{}
	}
	routeWebsite.Provide(r, h)

	if routeBucket == nil {
		routeBucket = &RouteBucket{}
	}
	routeBucket.Provide(r, h)
}
