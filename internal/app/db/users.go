package db

import (
	"database/sql"
)

func GetUserIDByEmail(db *sql.DB, email string) (int64, error) {
	var id int64
	var err = db.QueryRow(`SELECT id from users where email=$1`, email).Scan(&id)
	return id, err
}

// we cannot use rs.LastInsertId() because its not supported by pq
// https://github.com/lib/pq/issues/24
func InsertUser(db *sql.DB, email, password string, contactAllowed bool) error {
	var stmt = `INSERT INTO users(email, password_hash, contact_allowed) VALUES($1, $2, $3)`
	var _, err = db.Exec(stmt, email, password, contactAllowed)
	return err
}

func DeleteUser(db *sql.DB, email string) error {
	var stmt = `DELETE FROM users where email=$1`
	var _, err = db.Exec(stmt, &email)
	return err
}

func Login(db *sql.DB, email string) (int64, string, bool, error) {
	var stmt = `SELECT id, password_hash, contact_allowed from users where email=$1`
	var id int64
	var passwordHash string
	var contactAllowed bool
	var err = db.QueryRow(stmt, email).Scan(&id, &passwordHash, &contactAllowed)
	return id, passwordHash, contactAllowed, err
}
