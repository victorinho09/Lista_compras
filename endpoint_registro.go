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
	//si lo esta, no hacemos nada
	
	//si no lo esta, debemos meterle en la lista de usuarios de la base de datos
	

}
