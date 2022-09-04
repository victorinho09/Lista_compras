package main

import (
	"net/http"
)

func nombreLista(w http.ResponseWriter, peticion *http.Request) {
	if peticion.URL.Path == "/favicon.ico" {
		return
	}
	cookieValue, _ := getCookie(peticion)
	debug("/nombre lista entrando : cookie " + cookieValue)

	if cookieValue == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Must register first!"))
		return
	}

	//obtenemos el id de la lista
	ids, ok := peticion.URL.Query()["id"]

	if !ok || len(ids) != 1 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Wrong id!"))
		return
	}

	id := ids[0]

	debug("Tengo query args," + id)

	//conseguirmos el nombre de la lista dado el id

	nombreLista := findNameList(id)

	if nombreLista == "" {

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Error al buscar el nombre de lista"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(nombreLista))
	return

}
