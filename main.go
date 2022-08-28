package main

import (
	"log"
	"net/http"

	"github.com/google/uuid"
)

var numberCookies = 0

func generateRandomValueCookie() string {
	return uuid.New().String()
}

func main() {

	// se toma la peticion y se analiza el usuario
	http.HandleFunc("/register", registrarUsuario)

	direccion := ":8080" // Como cadena, no como entero; porque representa una dirección
	debug("Servidor listo escuchando en " + direccion)
	log.Fatal(http.ListenAndServe(direccion, nil))

}
