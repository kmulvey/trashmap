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

func GetPolygonsWithinArea(db *sql.DB, points string) ([]string, error) {
	// https://www.postgis.net/docs/ST_Within.html
	updateStmt := `SELECT id, user_id, ST_AsText(polygon) as poly FROM areas WHERE ST_Contains(ST_GeomFromText('POLYGON(($1))'), polygon)`
	var rs, err = db.Query(updateStmt, points)
	if err != nil {
		return nil, err
	}

	var polygons []string
	for rs.Next() {
		var polygon string
		err = rs.Scan(&polygon)
		if err != nil {
			return nil, err
		}
		polygons = append(polygons, polygon)
	}

	return polygons, err
}
