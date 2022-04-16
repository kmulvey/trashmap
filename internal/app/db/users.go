package db

import "database/sql"

func GetUserIDByEmail(db *sql.DB, email string) (int, error) {
	var id int
	var err = db.QueryRow(`SELECT id from users where email="$1"`).Scan(&id)
	return id, err
}

func InsertUser(db *sql.DB, email string, contactAllowed bool) error {
	updateStmt := `INSERT INTO users(email, contact_allowed) VALUES($1, $2)`
	var _, err = db.Exec(updateStmt, email, contactAllowed)
	return err
}

func DeleteUser(db *sql.DB, email string) error {
	updateStmt := `DELETE FROM users where email="$1"`
	var _, err = db.Exec(updateStmt, email)
	return err
}

func Login(db *sql.DB, id int, uuid string) error {
	updateStmt := `INSERT INTO auth(auth_token, user_id) VALUES($1, $2)`
	var _, err = db.Exec(updateStmt, uuid, id)
	return err
}
