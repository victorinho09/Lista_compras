package main

import (
	"encoding/json"
	"net/http"
)

func elementosLista(w http.ResponseWriter, peticion *http.Request) {
	if peticion.URL.Path == "/favicon.ico" {
		return
	}
	cookieValue, _ := getCookie(peticion)
	debug("/elementos lista entrando : cookie " + cookieValue)

	if cookieValue == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Must register first!"))
		return
	}

	//debo conseguir el usuario
	userId := existsCookieDB(cookieValue)
	//debo coger por los query args el id de la lista
	ids, ok := peticion.URL.Query()["id"]

	if !ok || len(ids) != 1 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Wrong id!"))
		return
	}

	id := ids[0]

	debug("Tengo query args," + id)

	//debo buscar los elementos que pertenecen a la lista en la base de datos
	//debo conseguir la lista
	elementos := getElementsFromList(id,userId)

	if elementos == nil{
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Wrong dbqueries!"))
		return
	}

	jsonEncoded, err := json.Marshal(elementos)

	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Error al codificar json"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonEncoded)

	

}
