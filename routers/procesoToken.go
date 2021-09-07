package routers

import (
	"errors"
	"github/luismiguel010/twittor/bd"
	"github/luismiguel010/twittor/models"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
)

// Email valor usado en todos los EndPoints
var Email string

// IDUsuario es el ID retornado del modelo, que se usará en todos los Endpoint
var IDUsuario string

// ProcesoToken para extraer sus valores
func ProcesoToken(token string) (*models.Claim, bool, string, error) {
	miClave := []byte("MastersdelDesarrollo_grupodeFacebook")
	claims := &models.Claim{}

	splitToken := strings.Split(token, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido")
	}

	token = strings.TrimSpace(splitToken[1])
	tkn, err := jwt.ParseWithClaims(token, claims, func(jwttoken *jwt.Token) (interface{}, error) {
		return miClave, nil
	})

	if err == nil {
		_, encontrado, ID := bd.ChequeoYaExisteUsuario(claims.Email)
		if encontrado {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}
		return claims, encontrado, ID, nil
	}

	if !tkn.Valid {
		return claims, false, string(""), errors.New("token invalido")
	}

	return claims, false, string(""), err
}
