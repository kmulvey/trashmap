package config

import "database/sql"

type Config struct {
	DBUsername string
	DBPassword string
	DBDatabase string
	DBHostname string
	DBPort     int
	DBConn     *sql.DB

	HTTPAddr string
}
