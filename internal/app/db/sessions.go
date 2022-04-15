package db

import "database/sql"

func ExpireSessions(db *sql.DB) error {
	var _, err = db.Exec(`DELETE from auth WHERE created < NOW() - INTERVAL '30 days'`)
	return err
}
