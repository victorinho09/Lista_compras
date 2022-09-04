package main

import (
	"net/http"
)

func resetElementos(w http.ResponseWriter, peticion *http.Request) {
	if peticion.URL.Path == "/favicon.ico" {
		return
	}
	cookieValue, _ := getCookie(peticion)
	debug("/reset elementos entrando : cookie " + cookieValue)

	if cookieValue == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Must register first!"))
		return
	}

	//debo modificar todos los elementos a no marcados
	err := updateResetElementos()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("400 - Error al reset de elementos!"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("200 - Elementos reseteados!"))
	return

}
