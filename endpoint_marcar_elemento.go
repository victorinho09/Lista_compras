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

	//debo coger el valor de marcado
	valores, ok := peticion.URL.Query()["valor"]

	if !ok || len(valores) != 1 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Wrong valor!"))
		return
	}

	valor := valores[0]

	debug("Tengo query args," + idElemento + valor)

	//debo modificar el marcado del elemento

	if valor == "si" {

		err := updateMarcado(idElemento, valor)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("400 - Error al modificar marcado elemento!"))
			return
		}

		//devolvemos marcado
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("200 - Elemento marcado!"))
		return
	} else {
		err := updateMarcado(idElemento, valor)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("400 - Error al modificar marcado elemento!"))
			return
		}

		//devolvemos desmarcado
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("200 - Elemento desmarcado!"))
		return
	}

}
