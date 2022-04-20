package db

import (
	"database/sql"
)

// InsertArea adds the user's pickup area to the areas table and returns new ID.
// Currently we do not check if this new area overlaps any others.
func InsertArea(db *sql.DB, userID int, points string) (int64, error) {
	//var updateStmt = `
	//	WITH new_row AS (
	//		INSERT INTO areas(user_id, polygon) VALUES($1, ST_GeometryFromText('POLYGON(($2))));
	//		RETURNING *
	//	)
	//	SELECT id from new_row;
	//	`
	var stmt = `INSERT INTO areas(user_id, polygon) VALUES($1, ST_GeometryFromText('POLYGON(($2))))`
	var rs, err = db.Exec(stmt, userID, points)
	if err != nil {
		return -1, err
	}
	return rs.LastInsertId()
}

// DeleteArea removes the user's pickup area from the areas table.
func DeleteArea(db *sql.DB, userID int) error {
	var stmt = `DELETE FROM areas WHERE user_id="$1"`
	var _, err = db.Exec(stmt, userID)
	return err
}

// GetPickupAreasWithinArea takes an area as a set of GPS points and
// returns all the trash pickup areas within it.
func GetPickupAreasWithinArea(db *sql.DB, points string) ([]string, error) {
	// https://www.postgis.net/docs/ST_Within.html
	var stmt = `SELECT id, user_id, ST_AsText(polygon) as poly FROM areas WHERE ST_Contains(ST_GeomFromText('POLYGON(($1))'), polygon)`
	var rs, err = db.Query(stmt, points)
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
