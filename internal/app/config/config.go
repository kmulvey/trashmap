package config

import (
	"database/sql"

	"github.com/kmulvey/trashmap/internal/app/db"
)

type Config struct {
	DBUsername string
	DBPassword string
	DBName     string
	DBHostname string
	DBPort     int
	DBConn     *sql.DB

	HTTPAddr string
}

func NewConfig(DBUsername, DBPassword, DBName, DBHostname, HTTPAddr string, DBPort int) (*Config, error) {
	var c = Config{
		DBUsername: DBUsername,
		DBPassword: DBPassword,
		DBName:     DBName,
		DBHostname: DBHostname,
		DBPort:     DBPort,
		HTTPAddr:   HTTPAddr,
	}
	var err error
	c.DBConn, err = db.DBConnect(DBHostname, DBUsername, DBPassword, DBName, DBPort)

	return &c, err
}
