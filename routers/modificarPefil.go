package routers

import (
	"encoding/json"
	"github/luismiguel010/twittor/bd"
	"github/luismiguel010/twittor/models"
	"net/http"
)

// ModificarPefil modifica el perfil de usuario
func ModificarPefil(w http.ResponseWriter, r *http.Request) {

	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Datos incorrectos "+err.Error(), http.StatusBadRequest)
		return
	}

	var status bool
	status, err = bd.ModificoRegistro(t, IDUsuario)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar modificar el registro. Reintente nuevamente "+err.Error(), http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(w, "No se ha logrado modificar el registro del usuario "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
