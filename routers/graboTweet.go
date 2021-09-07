package routers

import (
	"encoding/json"
	"github/luismiguel010/twittor/bd"
	"github/luismiguel010/twittor/models"
	"net/http"
	"time"
)

// GraboTweet permite grabar el tweet en la base de datos
func GraboTweet(w http.ResponseWriter, r *http.Request) {
	var mensaje models.Tweet
	err := json.NewDecoder(r.Body).Decode(&mensaje)
	if err != nil {
		http.Error(w, "Datos incorrectos "+err.Error(), http.StatusBadRequest)
		return
	}

	registro := models.GraboTweet{
		UserID:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	var status bool
	_, status, err = bd.InsertoTweet(registro)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar insertar el registro, reintente nuevamente"+err.Error(), http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(w, "No se ha logrado insertar el tweet", http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusCreated)

}
