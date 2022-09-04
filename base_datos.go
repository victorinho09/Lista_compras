package main

import (
	"database/sql" // Interactuar con bases de datos
	"fmt"          // Imprimir mensajes y esas cosas
	"io/ioutil"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql" // La librería que nos permite conectar a MySQL
)

var MySql *sql.DB

func init() {
	MySql, _ = obtenerBaseDeDatos()
}

func obtenerBaseDeDatos() (db *sql.DB, e error) {

	//leemos el fichero para obtener la contraseña
	file, errf := ioutil.ReadFile("./fichero_contraseña.txt")
	if errf != nil {
		log.Fatal(errf)
	}
	textLine := string(file)
	aParts := strings.Split(textLine, ":")

	if len(aParts) != 2 {
		log.Fatal("Mal fichero contraseñas")
	}

	pass := strings.TrimSpace(aParts[1])
	usuario := strings.TrimSpace(aParts[0])

	host := "tcp(192.168.64.4:3306)"
	nombreBaseDeDatos := "lista_compra"
	// Debe tener la forma usuario:contraseña@host/nombreBaseDeDatos
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s/%s?maxAllowedPacket=0&allowOldPasswords=1", usuario, pass, host, nombreBaseDeDatos))
	if err != nil {
		debugErr("Error : ********* %s", err.Error())
		return nil, err
	}
	return db, nil
}
