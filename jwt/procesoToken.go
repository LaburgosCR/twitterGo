package jwt

import (
	"errors"
	"strings"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/twitterGo/models"
)

var Email string
var IDUsuario string

func ProcesoToken(tk string, JWTSigm string) (*models.Claim, bool, string, error) {
	miClave := []byte(JWTSigm)
	var claims models.Claim

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return &claims, false, string(""), errors.New("Formato de token invalido")
	}
	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, &claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})
	if err == nil {
		//rutina que chequea contra la bd
	}

	if !tkn.Valid {
		return &claims, false, string(""), errors.New("Token Inválido")
	}

	return &claims, false, string(""), err
}
