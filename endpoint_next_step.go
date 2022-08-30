package main

import (
	"net/http"
)

var CmdRegister = "register"
var CmdAutenticate = "autenticate"
var CmdCorrecto = "continue"

func nextStep(w http.ResponseWriter, peticion *http.Request) {
	if peticion.URL.Path == "/favicon.ico" {
		return
	}

	//establecemos una cookie para el usuario
	cookieValue, _ := getCookie(peticion)
	debug("-----------------------------: %s", cookieValue)
	debug("Next-Step, usuario: %s\n", cookieValue)

	if cookieValue == "" {

		debug("Next-step: register")
		w.Header().Set("Content-Type", "text/html; charset=UTF-8")
		w.Write([]byte(CmdRegister))
		return
	}

	user := existsCookieDB(cookieValue)

	if user == "" {
		debug("Next-step: autenticate")
		w.Header().Set("Content-Type", "text/html; charset=UTF-8")
		w.Write([]byte(CmdAutenticate))
		return
	}

	debug("Next-step: continue")
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.Write([]byte(CmdCorrecto))
	return

}
