package seguranca

import "golang.org/x/crypto/bcrypt"

// Hash receb uma string e coloca um hash nela
func Hash(senha string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
	// bcrypt.DefaultCost é o tamanho do hash a ser criado
}

// Compara uma senha e um hash e retorna se elas são iguais
func VerificarSenha(senhaComHash, senhaString string) error {
	return bcrypt.CompareHashAndPassword([]byte(senhaComHash), []byte(senhaString))
}
