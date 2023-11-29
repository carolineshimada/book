package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

// retorna router com rotas configuradas
func Gerar() *mux.Router {
	r := mux.NewRouter()
	return routes.Config(r)
}
