package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

// Usuarios representa um repositório de usuarios
type Usuarios struct {
	db *sql.DB
}

// NovoRepositorioDeUsuarios cria um repositório de usuarios
func NovoRepositorioDeUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

// Criar insere um usuaŕio no banco de dados
func (u Usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	return 0, nil
}
