package main

type DBElement struct {
	Id      string
	Nombre  string
	ListaId string
	Marcado string
}

func DBElementNew() *DBElement {

	return &DBElement{}
}

func (e *DBElement) insert() error {
	//almaceno la estructura DBElement en la base de datos

	//comprobamos la injection
	sql := DBSprintf("INSERT INTO lista_compra.elementos (id,nombre,listaId,marcado) VALUES('%s','%s','%s','%s')",
		e.Id, e.Nombre, e.ListaId, e.Marcado)
	_, err := MySql.Exec(sql)

	if err != nil {
		debugErr("Error en insert: %s", err.Error())
		return err
	}

	return nil
}

func createDBElement(nombre string, idLista string) *DBElement {

	elemento := DBElementNew()
	elemento.ListaId = idLista
	elemento.Nombre = nombre
	elemento.Id = UuidSorted()
	elemento.Marcado = "no"
	return elemento
}

func findElements(l *list) []DBElement {

	elementos := []DBElement{}

	query := DBSprintf("SELECT nombre,id,listaId,marcado FROM lista_compra.elementos WHERE listaId='%s'", l.id)
	filas, err := MySql.Query(query)
	if err != nil {
		debugErr("Error en la query -> %s", err.Error())
		return nil
	}
	defer filas.Close()

	for filas.Next() {

		nombreElemento := DBElement{}

		err = filas.Scan(&nombreElemento.Nombre, &nombreElemento.Id, &nombreElemento.ListaId, &nombreElemento.Marcado)

		if err != nil {
			debugErr("Error en scan %s", err.Error())
			return nil
		}

		elementos = append(elementos, nombreElemento)
	}
	return elementos
}

func saveToDBElement(idLista string, nombre string) error {

	//crear el objeto DBElement
	elemento := createDBElement(nombre, idLista)

	//llamar a su metodo insert
	err := elemento.insert()

	if err != nil {
		return err
	}

	return nil
}

func borrarElementoInDB(idElemento string) error {

	debug("elemento" + idElemento)
	//debo hacer la query para borrar el elemento de la tabla de elementos
	sql := DBSprintf("DELETE FROM lista_compra.elementos WHERE id='%s'", idElemento)
	_, err := MySql.Exec(sql)

	if err != nil {
		debugErr("Error en delete: %s", err.Error())
		return err
	}

	return nil

}

func borrarTodosElementosLista(idLista string) error {

	//debo hacer la query para borrar los elementos de la tabla de elementos
	sql := DBSprintf("DELETE FROM lista_compra.elementos WHERE listaId='%s'", idLista)
	_, err := MySql.Exec(sql)

	if err != nil {
		debugErr("Error en delete: %s", err.Error())
		return err
	}

	return nil
}
