package db

import (
	"database/sql"
	"fmt"
)

func GetUserIDByEmail(db *sql.DB, schema, email string) (int64, error) {
	var id int64
	var err = db.QueryRow(fmt.Sprintf(`SELECT id from %s.users where email=$1`, schema), email).Scan(&id)
	return id, err
}

// we cannot use rs.LastInsertId() because its not supported by pq
// https://github.com/lib/pq/issues/24
func InsertUser(db *sql.DB, schema, email, password string, contactAllowed bool) error {
	var stmt = fmt.Sprintf(`INSERT INTO %s.users(email, password_hash, contact_allowed) VALUES($1, $2, $3)`, schema)
	var _, err = db.Exec(stmt, email, password, contactAllowed)
	return err
}

func DeleteUser(db *sql.DB, schema, email string) error {
	var stmt = fmt.Sprintf(`DELETE FROM %s.users where email=$1`, schema)
	var _, err = db.Exec(stmt, &email)
	return err
}

func Login(db *sql.DB, schema, email string) (int64, string, bool, error) {
	var stmt = fmt.Sprintf(`SELECT id, password_hash, contact_allowed from %s.users where email=$1`, schema)
	var id int64
	var passwordHash string
	var contactAllowed bool
	var err = db.QueryRow(stmt, email).Scan(&id, &passwordHash, &contactAllowed)
	return id, passwordHash, contactAllowed, err
}
