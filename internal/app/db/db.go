package db

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
