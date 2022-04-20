package areas

import (
	"errors"
	"strconv"
	"strings"

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
			return nil, errors.New("unable to marshal coordinates from db")
		}
	}

	return areas, nil
}

// SaveArea adds the user's pickup area to the areas table.
// Currently we do not check if this new area overlaps any others.
func SaveArea(config *config.Config, userID int, polygonStr string) (int64, error) {
	var polygonArr = strings.Split(polygonStr, ",")

	// first check that we got enough data
	if len(polygonArr)%2 != 0 {
		return -1, errors.New("gps points are malformed (lenght is odd, should be even given they are pairs")
	}

	// chech that the things in between the commas are actually floats
	for _, point := range polygonArr {
		if _, err := strconv.ParseFloat(point, 64); err != nil {
			return -1, err
		}
	}

	return db.InsertArea(config.DBConn, userID, polygonStr)
}
