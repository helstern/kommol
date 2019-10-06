package routes

import (
	"bytes"
	"github.com/gorilla/mux"
	"github.com/helstern/kommol/internal/presentation/api/gcp/object"
	"net/http"
)

type website struct{}

func (ww website) Provide(r *mux.Router, h object.GetHandler) {
	r.
		HandleFunc("/{_:.+}", func(w http.ResponseWriter, req *http.Request) {
			ww.Handle(w, req, h)
		}).
		Headers("X-KOMMOL-STRATEGY", "GCP_WEBSITE").
		Methods("GET")
}

func (_ website) ExtractPath(req *http.Request) string {
	var b bytes.Buffer

	b.WriteString("gs://")
	b.WriteString(req.URL.Host)
	b.WriteString(req.URL.Path)

	return b.String()
}

func (ww website) Handle(w http.ResponseWriter, req *http.Request, h object.GetHandler) {
	path := ww.ExtractPath(req)
	h(w, path)
}

var (
	routeWebsite *website
)

func Website() *website {
	if routeWebsite == nil {
		routeWebsite = &website{}
	}
	return routeWebsite
}
