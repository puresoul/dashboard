package userdb

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// Connection is an interface for making queries.
type Connection interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Get(dest interface{}, query string, args ...interface{}) error
	Select(dest interface{}, query string, args ...interface{}) error
}

func InsertUser(db Connection, em, pswd string) (string, error) {
	res, err := db.Exec(fmt.Sprint(`INSERT INTO Users VALUES (null,?, ?, 1)`), em, pswd)
	if err != nil {
		return "", err
	}
	uid, err := res.LastInsertId()
	return fmt.Sprint(uid), err
}

func UserId(db Connection, em string) (string, error) {
	var usr string
	err := db.Get(&usr, fmt.Sprint(`select UserId from Users where Email=?`), em)
	if err != nil {
		return usr, err
	}
	return usr, err
}

func UserMail(db Connection, id string) (string, error) {
	var usr string
	err := db.Get(&usr, fmt.Sprint(`select Email from Users where UserId=?`), id)
	if err != nil {
		return usr, err
	}
	return usr, err
}
