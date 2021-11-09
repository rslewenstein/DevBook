package rotas

import "net/http"

// Representa todas as rotas da API
type Rotas struct {
	URI                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}
