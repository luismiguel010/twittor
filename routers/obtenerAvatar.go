package routers

import (
	"github/luismiguel010/twittor/bd"
	"io"
	"net/http"
	"os"
)

// ObtenerAvatar envía el Avatar al HTTP
func ObtenerAvatar(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro ID", http.StatusBadRequest)
		return
	}

	perfil, err := bd.BuscoPerfil(ID)
	if err != nil {
		http.Error(w, "Usuario no encontrado", http.StatusBadRequest)
		return
	}

	openFile, err := os.Open("uploads/avatars/" + perfil.Avatar)
	if err != nil {
		http.Error(w, "Imagen no encontrado", http.StatusBadRequest)
		return
	}

	_, err = io.Copy(w, openFile)
	if err != nil {
		http.Error(w, "Error al copia la imagen", http.StatusBadRequest)
	}

}
