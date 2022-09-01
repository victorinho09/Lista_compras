package main

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
