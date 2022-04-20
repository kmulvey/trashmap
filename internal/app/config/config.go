package config

import (
	"database/sql"

	"github.com/kmulvey/trashmap/internal/app/db"
)

// Config is the main app config
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

// NewConfig is a Config constructor
func NewConfig(DBUsername, DBPassword, DBName, DBHostname string, DBPort int, HTTPAddr, PasswordSalt, HTTPReadSigningSecret, HTTPWriteSigningSecret string) (*Config, error) {
	var c = Config{
		DBUsername:             DBUsername,
		DBPassword:             DBPassword,
		DBName:                 DBName,
		DBHostname:             DBHostname,
		DBPort:                 DBPort,
		PasswordSalt:           PasswordSalt,
		HTTPAddr:               HTTPAddr,
		HTTPReadSigningSecret:  HTTPReadSigningSecret,
		HTTPWriteSigningSecret: HTTPWriteSigningSecret,
	}
	var err error
	c.DBConn, err = db.DBConnect(DBHostname, DBUsername, DBPassword, DBName, DBPort)

	return &c, err
}

func NewTestConfig() (*Config, error) {
	return NewConfig("postgres", "postgres", "postgres", "localhost", 5432, "http://localhost", "salt", "salt", "salt")
}
