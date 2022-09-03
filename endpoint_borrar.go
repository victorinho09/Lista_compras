package main

import (
	"net/http"
)

//comprobara si debe borrar una lista o un elemento y lo hara
func borrar(w http.ResponseWriter, peticion *http.Request) {
	if peticion.URL.Path == "/favicon.ico" {
		return
	}
	cookieValue, _ := getCookie(peticion)
	debug("/borrar entrando : cookie " + cookieValue)

	if cookieValue == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Must register first!"))
		return
	}

	//debo coger por los query args el id del elemento
	elementos, ok := peticion.URL.Query()["elemento"]

	if !ok {

		debug("doy fallo")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Wrong element!"))
		return
	}

	if len(elementos) != 1 {

		if len(elementos) > 1 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 - Wrong element len >1!"))
			return
		}
	} else {

		elemento := elementos[0]
		debug(elemento)

		//pongo list para que quede claro que debe hacer la borrada de la lista
		if elemento == "list" {
			//significa que nos han pasado una lista

			//debo coger por los query args el id de la lista
			ids, ok := peticion.URL.Query()["idLista"]

			if !ok || len(ids) != 1 {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("400 - Wrong idLista!"))
				return
			}

			idLista := ids[0]

			debug("Tengo query args," + idLista)

			//debo borrar la lista de la base de datos
			err := borrarListaInDB(idLista)

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("400 - Error al borrar lista!"))
				return
			}

			w.WriteHeader(http.StatusOK)
			w.Write([]byte("200 - Lista borrada con éxito!"))
			return
		} else {

			//cogemos el id del elemento ya que nos han pasado a borrar un elemento

			elemento := elementos[0]

			debug("Tengo query args, " + elemento)

			//debo borrar el elemento de la base de datos
			err := borrarElementoInDB(elemento)

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("400 - Error al borrar elemento!"))
				return
			}

			w.WriteHeader(http.StatusOK)
			w.Write([]byte("200 - Elemento borrado de la lista!"))
			return
		}
	}

}
