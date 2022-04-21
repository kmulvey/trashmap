package db

import (
	"database/sql"
	"fmt"
	"strings"
)

// InsertArea adds the user's pickup area to the areas table and returns new ID.
// Currently we do not check if this new area overlaps any others.
// we cannot use rs.LastInsertId() because its not supported by pq
// https://github.com/lib/pq/issues/24
func InsertArea(db *sql.DB, schema string, userID int64, polygon string) error {
	// we use fmt.Sprintf to fill in the polygon string because the sql driver
	// cant see variables within '', which are required by ST_GeometryFromText.
	// The polygon value has been processed quite a bit by this point anyway.
	var stmt = fmt.Sprintf(`INSERT INTO %s.areas(user_id, polygon) VALUES($1, ST_GeometryFromText('POLYGON((%s))'))`, schema, polygon)
	var _, err = db.Exec(stmt, userID)
	return err
}

// GetAreaID returns the id of an area given userID and the area
func GetAreaID(db *sql.DB, schema string, userID int64, polygon string) (int64, error) {
	// we use fmt.Sprintf to fill in the polygon string because the sql driver
	// cant see variables within '', which are required by ST_GeometryFromText.
	// The polygon value has been processed quite a bit by this point anyway.
	var stmt = fmt.Sprintf(`SELECT id FROM %s.areas where user_id=$1 and polygon=ST_GeometryFromText('POLYGON((%s))')`, schema, polygon)
	var id int64
	var err = db.QueryRow(stmt, userID).Scan(&id)
	if err != nil {
		return -1, err
	}

	return id, err
}

// DeleteArea removes the user's pickup area from the areas table.
func DeleteArea(db *sql.DB, schema string, userID int64) error {
	var stmt = fmt.Sprintf(`DELETE FROM %s.areas WHERE user_id=$1`, schema)
	var _, err = db.Exec(stmt, userID)
	return err
}

// GetPickupAreasWithinArea takes an area as a set of GPS points and
// returns all the trash pickup areas within it.
func GetPickupAreasWithinArea(db *sql.DB, schema string, polygon string) ([]string, error) {
	// we use fmt.Sprintf to fill in the polygon string because the sql driver
	// cant see variables within '', which are required by ST_GeometryFromText.
	// The polygon value has been processed quite a bit by this point anyway.
	// https://www.postgis.net/docs/ST_Within.html
	var stmt = fmt.Sprintf(`SELECT id, user_id, ST_AsText(polygon) as poly FROM %s.areas WHERE ST_Contains(ST_GeomFromText('POLYGON((%s))'), polygon)`, schema, polygon)
	var rs, err = db.Query(stmt)
	if err != nil {
		return nil, err
	}

	var polygons []string
	var rsPolygon string
	var id int64
	var userID int64
	for rs.Next() {
		err = rs.Scan(&id, &userID, &rsPolygon)
		if err != nil {
			return nil, err
		}
		rsPolygon = strings.ReplaceAll(rsPolygon, "POLYGON((", "")
		rsPolygon = strings.ReplaceAll(rsPolygon, "))", "")
		polygons = append(polygons, rsPolygon)
	}

	return polygons, err
}
