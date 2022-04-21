package db

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/lib/pq"
)

var createSql = `
DO $$
BEGIN
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(128) NOT NULL UNIQUE,
    password_hash VARCHAR(44) NOT NULL,
    contact_allowed boolean DEFAULT false
);
CREATE TABLE IF NOT EXISTS areas (
    id SERIAL PRIMARY KEY,
    user_id SERIAL,
    polygon GEOMETRY,
    FOREIGN KEY (user_id) REFERENCES users(id)
);
CREATE INDEX IF NOT EXISTS areas_polygon_idx ON areas USING GIST (polygon);
END;
$$;`

var lock sync.RWMutex // this is for tests

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
	if err != nil {
		return nil, err
	}

	// init tables
	lock.Lock()
	defer lock.Unlock()
	_, err = db.Exec(createSql)

	return db, err
}
