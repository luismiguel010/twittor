package bd

import (
	"github/luismiguel010/twittor/models"

	"golang.org/x/crypto/bcrypt"
)

// IntentoLogin realiza el chequeo de login a la BD
func IntentoLogin(email string, password string) (models.Usuario, bool) {
	usuario, encontrado, _ := ChequeoYaExisteUsuario(email)
	if !encontrado {
		return usuario, false
	}
	passwordByte := []byte(password)
	passwordBD := []byte(usuario.Password)
	err := bcrypt.CompareHashAndPassword(passwordBD, passwordByte)
	if err != nil {
		return usuario, false
	}
	return usuario, true
}
