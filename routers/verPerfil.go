package routers

import (
	"encoding/json"
	"github/luismiguel010/twittor/bd"
	"net/http"
)

// VerPerfil permite extraer los valores del perfil
func VerPerfil(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetros ID", http.StatusBadRequest)
		return
	}

	perfil, err := bd.BuscoPerfil(ID)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar buscar el registro"+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(perfil)
}
