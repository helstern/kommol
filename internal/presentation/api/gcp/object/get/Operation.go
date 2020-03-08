package get

import (
	logging "github.com/helstern/kommol/internal/core/logging/app"
	"github.com/helstern/kommol/internal/core/object"
	"github.com/helstern/kommol/internal/core/object/app"
	"github.com/helstern/kommol/internal/presentation/api"
	"io"
	"net/http"
	"strings"
)

type Operation func(w http.ResponseWriter, req *http.Request, obj object.Object)

func NewOperation(s app.ObjectProxy, l logging.LoggerFactory) Operation {
	operation := func(w http.ResponseWriter, req *http.Request, obj object.Object) {
		ctx := logging.WithLogContext(req.Context(), logging.Fields{
			"requestId": api.RequestId(req),
		})
		logger := logging.ContextLogger(ctx, l).WithFields(logging.Fields{
			"path": strings.Join(obj.Path, "/"),
		})
		logger.Info("retrieving path")
		httpResponse, err := s.Http(ctx, obj)

		if err != nil {
			logger.WithError(err).Info("failed to retrieve path")
			http.Error(w, "failed to retrieve", http.StatusInternalServerError)
			return
		}

		for _, v := range httpResponse.Headers {
			w.Header().Add(v.Name, v.Value)
		}

		size, err := io.Copy(w, httpResponse.Body)
		if err != nil {
			logger.WithError(err).Info("failed to write object")
			http.Error(w, "failed to retrieve", http.StatusInternalServerError)
		}
		logger.WithFields(logging.Fields{"bytes": size}).Info("wrote object")
	}

	return operation
}
