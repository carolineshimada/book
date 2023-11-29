package router

import (
	"github.com/gorilla/mux"
)

// retorna router com rotas configuradas
func Gerar() *mux.Router {
	return mux.NewRouter()
}
