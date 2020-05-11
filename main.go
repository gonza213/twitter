package main

import (
	"log"

	"github.com/gonza213/twitter/bd"
	"github.com/gonza213/twitter/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}
	handlers.Manejadores()
}
