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
	http.HandleFunc("/", downloads)
	http.HandleFunc("/autenticate",autenticate)
	http.HandleFunc("/next-step",nextStep)
	http.HandleFunc("/listas",listas)
	http.HandleFunc("/crear-lista",crearLista)
	http.HandleFunc("/elementos-lista",elementosLista)
	http.HandleFunc("/nuevo-elemento",nuevoElemento)
	http.HandleFunc("/borrar",borrar)
	http.HandleFunc("/marcar-elemento",marcarElemento)
	http.HandleFunc("/reset-elementos",resetElementos)
	http.HandleFunc("/nombre-lista",nombreLista)



	direccion := ":8080" // Como cadena, no como entero; porque representa una dirección
	debug("Servidor listo escuchando en " + direccion)
	log.Fatal(http.ListenAndServe(direccion, nil))

}
