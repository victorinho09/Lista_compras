package main

import "net/http"


func registrarUsuario(w http.ResponseWriter, peticion *http.Request) {
	if peticion.URL.Path == "/favicon.ico" {
		return
	}
	cookieValue, _ := getCookie(peticion)

	if cookieValue == "" {
		cookieValue = setNewCookie(w, peticion)
	}
	debug("-----------------------------: %s", cookieValue)
	debug("Entrada de cliente, usuario: %s\n", cookieValue)

	//comprobamos si el usuario ya esta metido en la base de datos

	//Para ello, comprobamos si la cookie existente esta en la base de datos
	//en la tabla de usuarios y cookies

	user := existsCookieDB(cookieValue)

	if user == "" {
		//no esta la cookie

		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Must autenticate"))
		return

	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("200 - Registered correctly"))
		return
		//si esta la cookie, mostraremos las listas que tiene el usuario
	}

}
