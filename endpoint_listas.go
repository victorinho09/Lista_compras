package main

import (
	"encoding/json"
	"net/http"
)

func listas(w http.ResponseWriter, peticion *http.Request) {
	if peticion.URL.Path == "/favicon.ico" {
		return
	}
	cookieValue, _ := getCookie(peticion)
	debug("/listas entrando : cookie " + cookieValue)

	if cookieValue == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Must register first!"))
		return
	}

	//conseguirmos las listas del usuario

	jsonEncoded, err := json.Marshal(todasListasDeUsuario((existsCookieDB((cookieValue)))))

	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Error al codificar json"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonEncoded)

}
