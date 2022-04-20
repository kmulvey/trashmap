package db

import "database/sql"

func GetUserIDByEmail(db *sql.DB, email string) (int, error) {
	var id int
	var err = db.QueryRow(`SELECT id from users where email='$1'`).Scan(&id)
	return id, err
}

func InsertUser(db *sql.DB, email, password string, contactAllowed bool) error {
	updateStmt := `INSERT INTO users(email, passwd, contact_allowed) VALUES($1, $2, $3)`
	var _, err = db.Exec(updateStmt, email, password, contactAllowed)
	return err
}

func DeleteUser(db *sql.DB, email string) error {
	updateStmt := `DELETE FROM users where email='$1'`
	var _, err = db.Exec(updateStmt, email)
	return err
}

func Login(db *sql.DB, email string) (int, string, bool, error) {
	var updateStmt = `SELECT id, password_hash, contact_allowed from users where email = '$1'`
	var id int
	var passwordHash string
	var contactAllowed bool
	var err = db.QueryRow(updateStmt, email).Scan(id, passwordHash, contactAllowed)
	return id, passwordHash, contactAllowed, err
}
