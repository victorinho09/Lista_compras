package main

import()

func todasListasDeUsuario(usuarioId string) []string{

	listas := []string{}

	query := DBSprintf("SELECT nombre FROM lista_compra.listas WHERE usuarioId='%s'", usuarioId)
	filas,err := MySql.Query(query)
	if err != nil{
		debugErr("Error en la query -> %s",err.Error())
		return nil
	}
	defer filas.Close()

	for filas.Next(){

		nombre := ""
		
		err = filas.Scan(&nombre)

		if err != nil{
			debugErr("Error en scan",err.Error())
			return nil
		}

		listas = append(listas,nombre)
	}
	return listas
}