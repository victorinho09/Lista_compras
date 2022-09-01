package main

import (
	"net/http"
)

func crearLista(w http.ResponseWriter, peticion *http.Request) {
	if peticion.URL.Path == "/favicon.ico" {
		return
	}
	cookieValue, _ := getCookie(peticion)
	debug("/crear lista entrando : cookie " + cookieValue)

	if cookieValue == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Must register first!"))
		return
	}

	//debo conseguir el usuario
	userId := existsCookieDB(cookieValue)

	//debo coger por los query args el nombre de la lista
	nombres, ok := peticion.URL.Query()["nombre"]

	if !ok || len(nombres) != 1 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Wrong name!"))
		return
	}

	nombre := nombres[0]

	debug("Tengo query args," + nombre)

	//debo crear la nueva lista con el nombre de usuario
	//establecerle un id unico

	list := createList(UuidSorted(), nombre, userId)
	err := list.saveToDBLists()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Error al crear y guardar lista!"))
		return
	}

}
