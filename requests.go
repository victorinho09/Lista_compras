package main

import (
	"fmt"
	"net/http"
)

func getCookie(r *http.Request) (string, error) {

	//establecemos una cookie para el cliente
	cookie, err := r.Cookie("Usuario")

	if err != nil && err != http.ErrNoCookie {
		return "", err
	}
	if cookie == nil {
		return "", nil
	}
	return cookie.Value, nil
}

func setNewCookie(w http.ResponseWriter, r *http.Request) string {

	cookieValue := generateRandomValueCookie()
	cookieNew := http.Cookie{Name: "Usuario", Value: cookieValue, Domain: r.URL.Host}
	debug("Se creo una cookie: " + fmt.Sprintf("%s", cookieValue))
	http.SetCookie(w, &cookieNew)
	return cookieValue

}
