package Routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Rout struct {
	Uri                   string
	Method                string
	Function              func(http.ResponseWriter, *http.Request)
	RequireAuthentication bool
}

func Config(r *mux.Router) *mux.Router {
	rout := routeUser

	for _, rout := range rout {
		r.HandleFunc(rout.Uri, rout.Function).Methods(rout.Method)
	}

	return r
}
