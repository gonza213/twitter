package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gonza213/twitter/middlew"
	"github.com/gonza213/twitter/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

//Manejadores
func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
