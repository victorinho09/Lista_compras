package main

import "database/sql"

type DBList struct {
	Id        string
	Nombre    string
	UsuarioId string
}

func DBListNew() *DBList {
	return &DBList{}
}

func todasListasDeUsuario(usuarioId string) []DBList {

	listas := []DBList{}

	query := DBSprintf("SELECT nombre,id,usuarioId FROM lista_compra.listas WHERE usuarioId='%s'", usuarioId)
	filas, err := MySql.Query(query)
	if err != nil {
		debugErr("Error en la query -> %s", err.Error())
		return nil
	}
	defer filas.Close()

	for filas.Next() {

		lista := DBList{}

		err = filas.Scan(&lista.Nombre, &lista.Id, &lista.UsuarioId)

		if err != nil {
			debugErr("Error en scan" + err.Error())
			return nil
		}
		listas = append(listas, lista)

	}
	return listas
}

func findNameList(idLista string) string {

	nombre := ""

	query := DBSprintf("SELECT nombre FROM lista_compra.listas WHERE id='%s'", idLista)
	fila := MySql.QueryRow(query)
	err := fila.Scan(&nombre)

	if err == sql.ErrNoRows {

		return ""

	}
	if err != nil {
		debugErr("Error de query -> %s", err.Error())
		return ""
	}
	return nombre
}

func (l *DBList) insert() error {

	//almaceno la estructura DBList en la base de datos

	//comprobamos la injection
	sql := DBSprintf("INSERT INTO lista_compra.listas (id,nombre,usuarioId) VALUES('%s','%s','%s')",
		l.Id, l.Nombre, l.UsuarioId)
	_, err := MySql.Exec(sql)

	if err != nil {
		debugErr("Error en insert: %s", err.Error())
		return err
	}

	return nil
}

func borrarListaInDBListas(idLista string) error{

	//debo hacer la query para borrar los elementos de la tabla de elementos
	sql := DBSprintf("DELETE FROM lista_compra.listas WHERE id='%s'", idLista)
	_, err := MySql.Exec(sql)

	if err != nil {
		debugErr("Error en delete: %s", err.Error())
		return err
	}

	return nil
}

func borrarListaInDB(idLista string) error{

	//primero debo borrar todos los elementos que tengan el idLista
	err := borrarTodosElementosLista(idLista)

	if err != nil{
		return err
	}

	//borro la lista de la lista de listas
	err = borrarListaInDBListas(idLista)
	if err != nil{
		return err
	}

	return err
}
