package operations

import (
	"context"
	"github.com/apex/log"
	"github.com/helstern/kommol/internal/core/object/app"
	"github.com/helstern/kommol/internal/presentation/api/gcp/object"
	"net/http"
)

func objectHandler(w http.ResponseWriter, path string, s app.ObjectRetriever) {
	logCtx := log.WithFields(log.Fields{
		"path": path,
	})
	logCtx.Info("retrieving path")

	o, err := s.Retrieve(path)

	if err != nil {
		logCtx.Info(err.Error())
		http.Error(w, "failed to retrieve", http.StatusInternalServerError)
		return
	}

	ctx := context.Background()

	err = o.ModifyHeaders(ctx, w.Header())
	if err != nil {
		logCtx.Info(err.Error())
		http.Error(w, "failed to retrieve", http.StatusInternalServerError)
		return
	}
	_, err = o.WriteContent(ctx, w)
	if err != nil {
		logCtx.Info(err.Error())
		http.Error(w, "failed to retrieve", http.StatusInternalServerError)
	}
}

func NewGetHandler(s app.ObjectRetriever) object.GetHandler {
	return func(w http.ResponseWriter, path string) {
		objectHandler(w, path, s)
	}
}
