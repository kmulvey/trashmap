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
		HTTPReadSigningSecret:  HTTPReadSigningSecret,  // not currently used, still early https://github.com/gin-contrib/httpsign
		HTTPWriteSigningSecret: HTTPWriteSigningSecret, // not currently used, still early https://github.com/gin-contrib/httpsign
	}
	var err error
	c.DBConn, err = db.DBConnect(DBHostname, DBUsername, DBPassword, DBName, DBPort)

	return &c, err
}

var globalTestConfig *Config // singleton for tests

func NewTestConfig() (*Config, error) {
	var err error
	if globalTestConfig == nil {
		globalTestConfig, err = NewConfig("postgres", "postgres", "postgres", "localhost", 5432, "http://localhost", "salt", "salt", "salt")
	}
	return globalTestConfig, err
}
