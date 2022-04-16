package db

import (
	"database/sql"
	"fmt"
	"strings"
)

func InsertArea(db *sql.DB, userID int, points []float64) error {
	//ST_GeometryFromText('POLYGON((50.6373 3.0750,50.6374 3.0750,50.6374 3.0749,50.63 3.07491,50.6373 3.0750))')
	updateStmt := `INSERT INTO areas(user_id, polygon) VALUES($1, ST_GeometryFromText('POLYGON(($2))))`
	var _, err = db.Exec(updateStmt, userID, floatSliceToString(points))
	return err
}

func floatSliceToString(points []float64) string {
	var builder = strings.Builder{}

	for i, point := range points {
		// this if() places the commas correctly so we dont have an extra trailing one
		if i == 0 {
			builder.WriteString(fmt.Sprintf("%f", point))
		} else {
			builder.WriteString(fmt.Sprintf(",%f", point))
		}
	}

	return builder.String()
}
