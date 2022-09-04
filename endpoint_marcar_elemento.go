package main

import (
	"net/http"
)

func marcarElemento(w http.ResponseWriter, peticion *http.Request) {
	if peticion.URL.Path == "/favicon.ico" {
		return
	}
	cookieValue, _ := getCookie(peticion)
	debug("/marcar elemento entrando : cookie " + cookieValue)

	if cookieValue == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Must register first!"))
		return
	}

	//debo coger por los query args el nombre del elemento
	idElementos, ok := peticion.URL.Query()["elemento"]

	if !ok || len(idElementos) != 1 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Wrong idElemento!"))
		return
	}

	idElemento := idElementos[0]
	debug("Tengo query args," + idElemento )

	//debo modificar el marcado del elemento
	err := updateMarcado(idElemento)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("400 - Error al modificar marcado elemento!"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("200 - Elemento marcado!"))
	return

}
