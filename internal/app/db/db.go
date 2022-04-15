package trashapp

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// DBConnect connects to postgres and returns the handle
func DBConnect(host, user, password, dbName string, port int) (*sql.DB, error) {
	var psqlconn = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

	// open database
	var db, err = sql.Open("postgres", psqlconn)
	if err != nil {
		return nil, err
	}

	// check db
	err = db.Ping()

	return db, err
}

func InsertUser(db *sql.DB, email string, contactAllowed bool) error {
	updateStmt := `INSERT INTO auth(email, contact_allowed) VALUES($1, $2)`
	var _, err = db.Exec(updateStmt, email, contactAllowed)
	return err
}
