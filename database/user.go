package database

import (
	"database/sql"
	"github.com/CardinalDevLab/Schwi-Backend/def"
	"github.com/CardinalDevLab/Schwi-Backend/utils"
)

func CreateUser(name string, password string, email string, experience int) error {
	//Insert New User
	password = utils.PasswordCrypto(password)
	statement, error := MainDatabase.Prepare("INSERT INTO users (name,password,email,experience) VALUES (?,?,?,?)")
	if error != nil {
		return error
	}

	_, error = statement.Exec(name, password, email, experience)
	if error != nil {
		return error
	}
	defer statement.Close()

	return nil
}

func GetUser(uid int, email string) (*def.User, error) {
	var query string
	if uid != 0 {
		query += `SELECT uid,name,password,email,experience FROM users WHERE uid = ?`
	} else if email != "" {
		query += `SELECT uid,name,password,email,experience FROM users WHERE email = ?`
	}
	statement, _ := MainDatabase.Prepare(query)
	var experience int
	var password, username string
	if uid != 0 {
		err = statement.QueryRow(uid).Scan(&uid, &username, &password, &email, &experience)
	} else if email != "" {
		err = statement.QueryRow(email).Scan(&uid, &username, &password, &email, &experience)
	}
	defer statement.Close()
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if err == sql.ErrNoRows {
		return nil, nil
	}
	res := &def.User{Uid: uid, Username: username, Password: password, Email: email, Experience: experience}

	return res, nil
}
