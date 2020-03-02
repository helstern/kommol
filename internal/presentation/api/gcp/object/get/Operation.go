package get

import (
	"github.com/apex/log"
	"github.com/helstern/kommol/internal/core/object"
	"github.com/helstern/kommol/internal/core/object/app"
	"io"
	"net/http"
)

type Operation func(w http.ResponseWriter, req *http.Request, obj object.Object)

func NewOperation(s app.ObjectProxy) Operation {
	operation := func(w http.ResponseWriter, req *http.Request, obj object.Object) {
		path := ""

		logCtx := log.WithFields(log.Fields{
			"path": path,
		})
		logCtx.Info("retrieving path")

		ctx := req.Context()
		httpResponse, err := s.Http(ctx, obj)

		if err != nil {
			logCtx.Info(err.Error())
			http.Error(w, "failed to retrieve", http.StatusInternalServerError)
			return
		}

		for _, v := range httpResponse.Headers {
			w.Header().Add(v.Name, v.Value)
		}

		size, err := io.Copy(w, httpResponse.Body)
		if err != nil {
			logCtx.Info(err.Error())
			http.Error(w, "failed to retrieve", http.StatusInternalServerError)
		}
		logCtx.Infof("wrote %d bytes", size)
	}

	return operation
}
