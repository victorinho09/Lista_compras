package main

import ()

type user struct{
	id string
	email string
	contraseña string
}

func userNew() *user{
	return &user{}
}

func createUser(id string,contraseña string,email string) *user{
	user := userNew()
	user.id = id
	user.contraseña = contraseña
	user.email = email
	return user
}

func (u *user) saveToDBCookie(cookieValue string) error{

	//me creo la estructura para meterla en la base de datos
	cookieDB := DBCookieNew()
	cookieDB.cookie = cookieValue
	cookieDB.usuarioId = u.id

	//meto la estructura en la tabla de cookies de la base de datos
	err := cookieDB.insert()
	if err != nil{
		debugErr("Error en el insert de DBCookie")
	}
	return err

}

func (u *user) saveToDBUser() error{

	//creo el usuario nuevo
	userDB := createNewUserDB(u)

	//lo meto en la base de datos de usuarios
	err := userDB.insert()

	if err != nil{
		return err
	}
	return nil
}

// crear el jugador nuevo con su email y contraseña
//y meterlo en la base de datos de usuarios y cookies
func (u *user) saveNewUserToDB(cookieValue string) error{

	//establezco un id unico al user
	u.id = UuidSorted()
	//Tabla cookies

	err := u.saveToDBCookie(cookieValue)

	if err != nil{
		return err
	}


	//Tabla usuarios

	err = u.saveToDBUser()

	if err != nil{
		return err
	}

	return nil
}