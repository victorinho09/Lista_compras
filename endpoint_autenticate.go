package main

import "net/http"

func autenticate(w http.ResponseWriter, peticion *http.Request) {
	if peticion.URL.Path == "/favicon.ico" {
		return
	}
	cookieValue, _ := getCookie(peticion)
	debug("/autenticate entrando : cookie " + cookieValue)

	if cookieValue == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Must register first!"))
		return
	}
	debug("-----------------------------: %s", cookieValue)
	debug("Cliente autenticate, usuario: %s\n", cookieValue)

	//debo coger por los query args el email
	emails, ok := peticion.URL.Query()["email"]

	if !ok || len(emails) != 1 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Wrong email!"))
		return
	}

	email := emails[0]
	//debo coger por los query args la contraseña

	contraseñas, ok := peticion.URL.Query()["contraseña"]

	if !ok || len(contraseñas) != 1 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Wrong password!"))
		return
	}
	contraseña := contraseñas[0]

	debug("Tengo query args," + contraseña + " " + email)

	//ya tengo los query args
	//debo comprobar si en la base de datos coinciden con algun usuario ya registrado
	userId := existsUserDB(contraseña, email)

	//creamos la estructura de tipo user
	user := createUser(userId, contraseña, email)

	if user.id == "" {
		//no existe,debo crear el jugador nuevo con su email y contraseña
		//y meterlo en la base de datos de usuarios y cookies

		err := user.saveNewUserToDB(cookieValue)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - Error al guardar cookie y usuario nuevo!"))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("200 - Jugador nuevo correcto!"))
		return

	} else {

		//debo añadir esta cookie a la base de datos de cookies
		err := user.saveToDBCookie(cookieValue)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - Error al guardar cookie!"))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("200 - Cookie añadida correctamente!"))
		return

	}

	//debo ir a otro endpoint donde se muestren las listas disponibles para mostrarle

}
