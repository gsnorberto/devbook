package router

import "github.com/gorilla/mux"

// Retorna um router com as rotas configuradas
func Generate() *mux.Router {
	return mux.NewRouter()
}
