package polygon

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
func GetPickupAreasWithinArea(config *config.Config, polygonStr string) ([]*gps.Coordinate, error) {
	var polygonArr = strings.Split(polygonStr, ",")

	// first check that we got enough data
	if len(polygonArr) == 8 {
		return nil, errors.New("gps points are malformed (lenght must be 8)")
	}

	// chech that the things in between the commas are actually floats
	for _, point := range polygonArr {
		if _, err := strconv.ParseFloat(point, 64); err != nil {
			return nil, err
		}
	}

	// get polys from db
	var coordinateStrArr, err = db.GetPickupAreasWithinArea(config.DBConn, polygonStr)
	if err != nil {
		return nil, err
	}

	// change the strings into *gps.Coordinate
	var coordinates = make([]*gps.Coordinate, len(coordinateStrArr))
	for i, coordinate := range coordinateStrArr {
		var split = strings.Split(coordinate, " ")
		if len(split) != 2 {
			return nil, errors.New("unable to marshal coordinates from db")
		}

		coordinates[i], err = gps.NewCoordinateFromString(coordinate)
		if len(split) != 2 {
			return nil, err
		}
	}

	return coordinates, nil
}

// SaveArea adds the user's pickup area to the areas table.
// Currently we do not check if this new area overlaps any others.
func SaveArea(config *config.Config, userID int, polygonStr string) error {
	var polygonArr = strings.Split(polygonStr, ",")

	// first check that we got enough data
	if len(polygonArr)%2 != 0 {
		return errors.New("gps points are malformed (lenght is odd, should be even given they are pairs")
	}

	// chech that the things in between the commas are actually floats
	for _, point := range polygonArr {
		if _, err := strconv.ParseFloat(point, 64); err != nil {
			return err
		}
	}

	return db.InsertArea(config.DBConn, userID, polygonStr)
}
