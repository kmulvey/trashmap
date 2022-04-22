package config

import (
	"github.com/jackc/pgx/v4"
	"github.com/kmulvey/trashmap/internal/app/db"
)

// Config is the main app config
type Config struct {
	DBUsername string
	DBPassword string
	DBName     string
	DBSchema   string
	DBHostname string
	DBPort     int
	DBConn     *pgx.Conn

	PasswordSalt string

	HTTPS                  bool
	HTTPBindAddr           string // e.g. ":8000"
	HTTPReadSigningSecret  string
	HTTPWriteSigningSecret string
	Development            bool // true=local false=prod
}

// NewConfig is a Config constructor
func NewConfig(dbUsername, dbPassword, dbName, dbSchema, dbHostname string, dbPort int, httpBindAddr, passwordSalt, httpReadSigningSecret, httpWriteSigningSecret string, https, development bool) (*Config, error) {
	var c = Config{
		DBUsername:             dbUsername,
		DBPassword:             dbPassword,
		DBName:                 dbName,
		DBSchema:               dbSchema,
		DBHostname:             dbHostname,
		DBPort:                 dbPort,
		PasswordSalt:           passwordSalt,
		HTTPS:                  https,
		HTTPBindAddr:           httpBindAddr,
		HTTPReadSigningSecret:  httpReadSigningSecret,  // not currently used, still early https://github.com/gin-contrib/httpsign
		HTTPWriteSigningSecret: httpWriteSigningSecret, // not currently used, still early https://github.com/gin-contrib/httpsign
		Development:            development,
	}
	var err error
	c.DBConn, err = db.DBConnect(dbHostname, dbUsername, dbPassword, dbName, dbSchema, dbPort)

	return &c, err
}

func NewTestConfig(schemaName string) (*Config, error) {
	return NewConfig("postgres", "postgres", "postgres", schemaName, "localhost", 5432, ":8000", "salt", "salt", "salt", false, true)
}
