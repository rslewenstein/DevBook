package controllers

import "net/http"

func CriarUsuario(w http.ResponseWriter, r *http.Response) {
	w.Write([]byte("Criando usuários!"))
}
func BuscarUsuarios(w http.ResponseWriter, r *http.Response) {
	w.Write([]byte("Buscando todos os usuários!"))
}
func BuscarUsuario(w http.ResponseWriter, r *http.Response) {
	w.Write([]byte("Buscando um usuário!"))
}
func AtualizarUsuario(w http.ResponseWriter, r *http.Response) {
	w.Write([]byte("Atualizando usuário!"))
}
func DeletarUsuario(w http.ResponseWriter, r *http.Response) {
	w.Write([]byte("Deletando usuário!"))
}
