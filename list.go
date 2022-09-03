package main

type list struct {
	id        string
	nombre    string
	usuarioId string
}

func listNew() *list {
	return &list{}
}

func createList(id string, nombre string, usuarioId string) *list {
	list := listNew()
	list.id = id
	list.nombre = nombre
	list.usuarioId = usuarioId
	return list
}

func getElementsFromList(idLista string, userId string) []DBElement {

	//busco el nombre de mi lista
	nombreLista := findNameList(idLista)
	//creo mi variable lista
	list := createList(idLista, nombreLista, userId)

	return findElements(list)
}

func (l *list) saveToDBLists() error {

	//me creo la estructura para meterla en la base de datos, por si acaso el programa crece
	//ya que ahora mismo la estructura de list valdría
	DBList := DBListNew()
	DBList.Nombre = l.nombre
	DBList.UsuarioId = l.usuarioId
	DBList.Id = l.id

	//meto la estructura en la tabla de listas de la base de datos
	err := DBList.insert()
	if err != nil {
		debugErr("Error en el insert de DBList")
	}
	return err

}
