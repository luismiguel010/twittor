package routers

import (
	"encoding/json"
	"github/luismiguel010/twittor/bd"
	"github/luismiguel010/twittor/jwt"
	"github/luismiguel010/twittor/models"
	"net/http"
	"time"
)

// Login realiza el login recibiendo un writter y un reques del paquete http
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Usuario y/o contraseña invalidos"+err.Error(), http.StatusBadRequest)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email del usuario es requerido", http.StatusBadRequest)
		return
	}

	documento, existe := bd.IntentoLogin(t.Email, t.Password)
	if !existe {
		http.Error(w, "Usuario y/o contraseña invalidos", http.StatusBadRequest)
		return
	}

	jwtKey, err := jwt.GeneroJWT(documento)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar generar el Token correspondiente "+err.Error(), http.StatusInternalServerError)
		return
	}

	resp := models.RespuestaLogin{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}
