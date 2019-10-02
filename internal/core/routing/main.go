package routing

import "github.com/gorilla/mux"

type RouteProvider func(r *mux.Router)
