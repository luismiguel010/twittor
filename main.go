package main

import (
	"github/luismiguel010/twittor/bd"
	"github/luismiguel010/twittor/handlers"
	"log"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexión a la base de datos")
		return
	}
	handlers.Manejadores()
}