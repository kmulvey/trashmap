package db

import (
	"database/sql"
)

func InsertArea(db *sql.DB, userID int, points string) error {
	//ST_GeometryFromText('POLYGON((50.6373 3.0750,50.6374 3.0750,50.6374 3.0749,50.63 3.07491,50.6373 3.0750))')
	updateStmt := `INSERT INTO areas(user_id, polygon) VALUES($1, ST_GeometryFromText('POLYGON(($2))))`
	var _, err = db.Exec(updateStmt, userID, points)
	return err
}

func GetPolygonsWithinArea(db *sql.DB, points string) error {
	updateStmt := `INSERT INTO areas(user_id, polygon) VALUES($1, ST_GeometryFromText('POLYGON(($2))))`
	var _, err = db.Exec(updateStmt, points)
	return err
}
