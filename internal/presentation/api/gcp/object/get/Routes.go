package get

import (
	"github.com/gorilla/mux"
	logging "github.com/helstern/kommol/internal/core/logging/app"
	"github.com/helstern/kommol/internal/core/object/app"
)

var (
	routeWebsite *RouteWebsite
	routeBucket  *RouteBucket
)

func Routes(r *mux.Router, s app.ObjectProxy, l logging.LoggerFactory) {

	h := NewOperation(s, l)

	if routeWebsite == nil {
		routeWebsite = &RouteWebsite{}
	}
	routeWebsite.Provide(r, h)

	if routeBucket == nil {
		routeBucket = &RouteBucket{}
	}
	routeBucket.Provide(r, h)
}
