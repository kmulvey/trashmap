package main

import (
	"os"

	"github.com/kmulvey/trashmap/internal/app/config"
	"github.com/kmulvey/trashmap/internal/app/webserver"
	log "github.com/sirupsen/logrus"
	cli "github.com/urfave/cli/v2"
)

func main() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	var config, err = configFromUserOps()
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(webserver.StartWebServer(config))
}

func configFromUserOps() (*config.Config, error) {

	var postgresHost string
	var postgresUser string
	var postgresPassword string
	var postgresDBName string
	var postgresSchemaName string
	var postgresPort int
	var httpHost string
	var httpReadSignSecret string
	var httpWriteSignSecret string
	var passwordHashSalt string

	var app = &cli.App{
		Name:  "TrashMap",
		Usage: "a website that allows people to commit to cleaning up trash in their selected area.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "postgres-host",
				Value:       "localhost",
				Usage:       "db host",
				EnvVars:     []string{"POSTGRES_HOST"},
				Destination: &postgresHost,
			},
			&cli.StringFlag{
				Name:        "postgres-user",
				Value:       "postgres",
				Usage:       "db user",
				EnvVars:     []string{"POSTGRES_USER"},
				Destination: &postgresUser,
			},
			&cli.StringFlag{
				Name:        "postgres-password",
				Value:       "postgres",
				Usage:       "db password",
				EnvVars:     []string{"POSTGRES_PASSWORD"},
				Destination: &postgresPassword,
			},
			&cli.StringFlag{
				Name:        "postgres-db-name",
				Value:       "postgres",
				Usage:       "db name",
				EnvVars:     []string{"POSTGRES_DB_NAME"},
				Destination: &postgresDBName,
			},
			&cli.StringFlag{
				Name:        "postgres-schema-name",
				Value:       "public",
				Usage:       "schema name",
				EnvVars:     []string{"POSTGRES_SCHEMA_NAME"},
				Destination: &postgresSchemaName,
			},
			&cli.IntFlag{
				Name:        "postgres-port",
				Value:       5432,
				Usage:       "db name",
				EnvVars:     []string{"POSTGRES_PORT"},
				Destination: &postgresPort,
			},
			&cli.StringFlag{
				Name:        "http-host",
				Value:       "http://localhost",
				Usage:       "http host",
				EnvVars:     []string{"HTTP_HOST"},
				Destination: &httpHost,
			},
			&cli.StringFlag{
				Name:        "http-read-sign-secret",
				Value:       "devonly",
				Usage:       "http read signing secret",
				EnvVars:     []string{"HTTP_READ_SIGN_SECRET"},
				Destination: &httpReadSignSecret,
			},
			&cli.StringFlag{
				Name:        "http-write-sign-secret",
				Value:       "devonly",
				Usage:       "http write signing secret",
				EnvVars:     []string{"HTTP_WRITE_SIGN_SECRET"},
				Destination: &httpWriteSignSecret,
			},
			&cli.StringFlag{
				Name:        "password-hash-salt",
				Value:       "devonly",
				Usage:       "salt for password hashing",
				EnvVars:     []string{"PASSWORD_HASH_SALT"},
				Destination: &httpWriteSignSecret,
			},
		},
	}
	var err = app.Run(os.Args)
	if err != nil {
		return nil, err
	}

	return config.NewConfig(postgresUser, postgresPassword, postgresDBName, postgresSchemaName, postgresHost, postgresPort, httpHost, passwordHashSalt, httpReadSignSecret, httpWriteSignSecret)
}
