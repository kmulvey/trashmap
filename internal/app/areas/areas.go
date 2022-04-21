package areas

import (
	"fmt"

	"github.com/kmulvey/trashmap/internal/app/config"
	"github.com/kmulvey/trashmap/internal/app/db"
	"github.com/kmulvey/trashmap/internal/pkg/gps"
)

// GetPickupAreasWithinArea takes an area as a set of GPS points and
// returns all the trash pickup areas within it in the gps.Coordinate format.
func GetPickupAreasWithinArea(config *config.Config, points *gps.Area) (gps.Map, error) {

	// get polys from db
	var areasStrArr, err = db.GetPickupAreasWithinArea(config.DBConn, points.CoordinatesToPostGISString())
	if err != nil {
		return nil, err
	}

	// change the strings into *gps.Coordinate
	var areas = make([]*gps.Area, len(areasStrArr))
	for i, area := range areasStrArr {
		var err error
		areas[i], err = gps.NewAreaFromPostGISString(area)
		if err != nil {
			return nil, fmt.Errorf("unable to marshal coordinates from db: %w", err)
		}
	}

	return areas, nil
}

// SaveArea adds the user's pickup area to the areas table.
// Currently we do not check if this new area overlaps any others.
func SaveArea(config *config.Config, userID int64, polygon *gps.Area) (int64, error) {
	var err = db.InsertArea(config.DBConn, userID, polygon.CoordinatesToPostGISString())
	if err != nil {
		return -1, err
	}

	return db.GetAreaID(config.DBConn, userID, polygon.CoordinatesToPostGISString())
}

// ReomveArea removes the user's pickup area to the areas table.
func RemoveArea(config *config.Config, userID int64, polygon *gps.Area) error {
	return db.DeleteArea(config.DBConn, userID)
}
