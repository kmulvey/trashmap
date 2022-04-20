package db

import (
	"database/sql"
)

// InsertArea adds the user's pickup area to the areas table.
// Currently we do not check if this new area overlaps any others.
func InsertArea(db *sql.DB, userID int, points string) error {
	updateStmt := `INSERT INTO areas(user_id, polygon) VALUES($1, ST_GeometryFromText('POLYGON(($2))))`
	var _, err = db.Exec(updateStmt, userID, points)
	return err
}

// GetPickupAreasWithinArea takes an area as a set of GPS points and
// returns all the trash pickup areas within it.
func GetPickupAreasWithinArea(db *sql.DB, points string) ([]string, error) {
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
