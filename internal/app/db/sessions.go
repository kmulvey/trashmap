package db

import "database/sql"

func ExpireSessions(db *sql.DB) error {
	var _, err = db.Exec(`DELETE from auth WHERE created < NOW() - INTERVAL '3 hours'`)
	return err
}

func GetSessions(db *sql.DB) ([]string, error) {
	var rows, err = db.Query(`SELECT uuid from auth WHERE created < NOW() - INTERVAL '3 hours'`)
	if err != nil {
		return nil, err
	}

	var sessions []string
	for rows.Next() {
		var uuid string
		err = rows.Scan(&uuid)
		if err != nil {
			return nil, err
		}
		sessions = append(sessions, uuid)
	}

	return sessions, err
}
