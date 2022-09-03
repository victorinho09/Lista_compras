package main

import "database/sql"

type DBUser struct{
	id string
	email string
	contraseña string
}

func DBUserNew() *DBUser{

	return &DBUser{}
}

func createNewUserDB(u *user) *DBUser{
	
	user := DBUserNew()

	user.contraseña = u.contraseña
	user.email = u.email
	user.id = u.id
	return user
}

func (u *DBUser) insert() error {

	//almaceno la estructura DBUser en la base de datos

	//comprobamos la injection
	sql := DBSprintf("INSERT INTO lista_compra.usuarios (id,email,contraseña) VALUES('%s','%s','%s')",
		u.id,u.email,u.contraseña)
	_, err := MySql.Exec(sql)

	if err != nil {
		debugErr("Error en insert: %s", err.Error())
		return err
	}

	return nil
}

//devuelve el id del usuario 
func existsUserDB(contraseña string,email string) string{

	//busca en la base de datos en la lista de usuarios

	user := ""

	query := DBSprintf("SELECT id FROM lista_compra.usuarios WHERE email='%s' AND contraseña='%s' ", email,contraseña)
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



