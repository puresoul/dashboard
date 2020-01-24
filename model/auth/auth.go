package auth

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Access struct {
	UserId   string `db:"UserId"`
	Password string `db:"Password"`
}

// Connection is an interface for making queries.
type Connection interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Get(dest interface{}, query string, args ...interface{}) error
	Select(dest interface{}, query string, args ...interface{}) error
}

func LogIn(db Connection, email, password string) (bool, string, error) {
	var result Access
	err := db.Get(&result, fmt.Sprint(`select UserId,Password from Users where Email=? and statusId=1`), email)
	if result.Password == password {
		return true, result.UserId, err
	}
	return false, "0", err
}
