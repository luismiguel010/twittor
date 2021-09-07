package routers

import (
	"encoding/json"
	"github/luismiguel010/twittor/bd"
	"github/luismiguel010/twittor/models"
	"net/http"
)

// Registro es la función para crear en la BD el registro de usuario
func Registro(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), http.StatusNotFound)
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email de usuario es requerido", http.StatusBadRequest)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "Debe especificar una contraseña de al menos 6 caracteres", http.StatusBadRequest)
		return
	}

	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado {
		http.Error(w, "Ya existe un usuario registrado con ese email", http.StatusBadRequest)
		return
	}

	_, status, err := bd.InsertoRegistro(t)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar realizar el registro de usuario "+err.Error(), http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(w, "No se ha logrado insertar el registro de usuario", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
