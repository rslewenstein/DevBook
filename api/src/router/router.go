package router

import "github.com/gorilla/mux"

// vai retornar um router com as rotas configuradas
func Gerar() *mux.Router {
	return mux.NewRouter()
}
