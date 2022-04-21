package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
)

func getCreateSql(schema string) string {
	return fmt.Sprintf(`
DO $$
BEGIN
CREATE TABLE IF NOT EXISTS %s.users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(128) NOT NULL UNIQUE,
    password_hash VARCHAR(44) NOT NULL,
    contact_allowed boolean DEFAULT false
);
CREATE TABLE IF NOT EXISTS %s.areas (
    id SERIAL PRIMARY KEY,
    user_id SERIAL,
    polygon GEOMETRY,
    FOREIGN KEY (user_id) REFERENCES %s.users(id)
);
CREATE INDEX IF NOT EXISTS areas_polygon_idx ON %s.areas USING GIST (polygon);
END;
$$;`, schema, schema, schema, schema)
}

// DBConnect connects to postgres and returns the handle
func DBConnect(host, user, password, dbName, schemaName string, port int) (*pgx.Conn, error) {
	var psqlconn = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

	// open database
	var ctx = context.Background()
	var db, err = pgx.Connect(ctx, psqlconn)
	if err != nil {
		return nil, err
	}

	// check db
	err = db.Ping(ctx)
	if err != nil {
		return nil, err
	}

	// init schema & tables
	_, err = db.Exec(ctx, "CREATE SCHEMA IF NOT EXISTS "+schemaName)
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(ctx, getCreateSql(schemaName))

	return db, err
}
