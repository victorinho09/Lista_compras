package main

import()

type DBElement struct{
	Id string
	Nombre string
	ListaId string
}


func findElements(l *list) []DBElement{

	elementos := []DBElement{}

	query := DBSprintf("SELECT (nombre,id,listaId) FROM lista_compra.elementos WHERE listaId='%s'", l.id)
	filas,err := MySql.Query(query)
	if err != nil{
		debugErr("Error en la query -> %s",err.Error())
		return nil
	}
	defer filas.Close()

	for filas.Next(){

		nombreElemento := DBElement{}
		
		err = filas.Scan(&nombreElemento.Nombre,&nombreElemento.Id,&nombreElemento.ListaId)

		if err != nil{
			debugErr("Error en scan",err.Error())
			return nil
		}

		elementos = append(elementos,nombreElemento)
	}
	return elementos
}