package main

import "database/sql"

type DBCookie struct {
	cookie    string
	usuarioId string
}

func DBCookieNew() *DBCookie {
	return &DBCookie{}
}

func existsCookieDB(cookie string) string {

	//busca en la base de datos en la lista de cookies

	user := ""

	query := DBSprintf("SELECT usuarioId FROM lista_compra.cookies WHERE cookie='%s' LIMIT 1", cookie)
	fila := MySql.QueryRow(query)
	err := fila.Scan(&user)

	if err == sql.ErrNoRows {

		return ""

	}
	if err != nil {
		debugErr("Error de query -> %s", err.Error())
	}
	return user
}

func (c *DBCookie) insert() error {

	//almaceno la estructura DBCookie en la base de datos

	//comprobamos la injection
	sql := DBSprintf("INSERT INTO lista_compra.cookies (cookie,usuarioId) VALUES('%s','%s')",
		c.cookie, c.usuarioId)
	_, err := MySql.Exec(sql)

	if err != nil {
		debugErr("Error en insert: %s", err.Error())
		return err
	}

	return nil
}




