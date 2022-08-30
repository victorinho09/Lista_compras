package main

import (
	_ "embed"
	"net/http"
	"os"
)

func downloads(w http.ResponseWriter, peticion *http.Request) {
	if peticion.URL.Path == "/favicon.ico" {
		return
	}

	debug("downloads: " + peticion.URL.Path)

	switch peticion.URL.Path {
	case "/funciones.js":
		contentFile, err := os.ReadFile("./html/funciones.js")
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404 - Not found"))
			return
		}

		w.Header().Set("Content-Type", "text/javascript; charset=UTF-8")
		w.Write(contentFile)

	case "/estilos.css":
		contentFile, err := os.ReadFile("./html/estilos.css")
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404 - Not found"))
			return
		}

		w.Header().Set("Content-Type", "text/css; charset=UTF-8")
		w.Write(contentFile)

	default:
		contentFile, err := os.ReadFile("./html/pagina_principal.html")
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404 - Not found"))
			return
		}

		w.Header().Set("Content-Type", "text/html; charset=UTF-8")
		w.Write(contentFile)

	}

}
