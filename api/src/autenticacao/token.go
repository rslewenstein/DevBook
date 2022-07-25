package autenticacao

import (
	"api/src/config"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// retorna um token assinado
func CriarToken(usuarioID uint64) (string, error) {
	permissoes := jwt.MapClaims{}
	permissoes["authorized"] = true
	// expiração do token (6h)
	permissoes["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissoes["usuarioId"] = usuarioID
	// metodo de assinatura
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)
	// com o secret estamos assinando
	return token.SignedString([]byte(config.SecretKey))
}
