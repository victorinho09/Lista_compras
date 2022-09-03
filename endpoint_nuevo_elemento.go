package main

import (
	"net/http"
)

func nuevoElemento(w http.ResponseWriter, peticion *http.Request) {
	if peticion.URL.Path == "/favicon.ico" {
		return
	}
	cookieValue, _ := getCookie(peticion)
	debug("/nuevo elemento entrando : cookie " + cookieValue)

	if cookieValue == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Must register first!"))
		return
	}

	//debo coger por los query args el nombre del elemento
	nombres, ok := peticion.URL.Query()["nombre"]

	if !ok || len(nombres) != 1 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Wrong name!"))
		return
	}

	nombre := nombres[0]

	//debo coger por los query args el id de la lista
	ids, ok := peticion.URL.Query()["idLista"]

	if !ok || len(ids) != 1 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Wrong id!"))
		return
	}

	idLista := ids[0]

	debug("Tengo query args," + nombre + " " + idLista)

	//debo crear un nuevo DBElement y meterlo en la base de datos de elementos
	err := saveToDBElement(idLista, nombre)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("400 - Error al guardar elemento!"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("200 - Elemento añadido a lista!"))
	return

}
