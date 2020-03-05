package get

import (
	"github.com/helstern/kommol/internal/core/object"
	"github.com/helstern/kommol/internal/presentation/api/gcp/object/get"
	"net/http"
	"strings"
)

func createGetObjectOperation() get.Operation {
	op := func(w http.ResponseWriter, req *http.Request, obj object.Object) {

		var sbuilder strings.Builder
		sbuilder.WriteString("gs:")
		sbuilder.WriteString("//")
		sbuilder.WriteString(strings.Join(obj.Path, "/"))

		_, err := w.Write([]byte(sbuilder.String()))
		if err != nil {
			panic(err)
		}
	}
	return op
}
