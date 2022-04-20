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

	PasswordSalt string

	HTTPAddr               string
	HTTPReadSigningSecret  string
	HTTPWriteSigningSecret string
}

func NewConfig(DBUsername, DBPassword, DBName, DBHostname string, DBPort int, HTTPAddr, PasswordSalt string) (*Config, error) {
	var c = Config{
		DBUsername:   DBUsername,
		DBPassword:   DBPassword,
		DBName:       DBName,
		DBHostname:   DBHostname,
		DBPort:       DBPort,
		PasswordSalt: PasswordSalt,
		HTTPAddr:     HTTPAddr,
	}
	var err error
	c.DBConn, err = db.DBConnect(DBHostname, DBUsername, DBPassword, DBName, DBPort)

	return &c, err
}
