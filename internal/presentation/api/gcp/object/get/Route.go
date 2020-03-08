package get

import (
	"github.com/gorilla/mux"
	coreObject "github.com/helstern/kommol/internal/core/object"
	"net/http"
)

type Route interface {
	Provide(r *mux.Router, h Operation)
	ExtractObject(req *http.Request) (coreObject.Object, error)
}

func Params(req *http.Request, route Route) coreObject.Object {
	obj, _ := route.ExtractObject(req)
	return obj
}

func Handle(w http.ResponseWriter, req *http.Request, route Route, h Operation) {
	h(w, req, Params(req, route))
}
