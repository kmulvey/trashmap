package polygon

import (
	"errors"
	"strconv"
	"strings"

	"github.com/kmulvey/trashmap/internal/app/config"
	"github.com/kmulvey/trashmap/internal/app/db"
)

func GetPolygonsWithinArea(config *config.Config, polygonStr string) error {
	var polygonArr = strings.Split(polygonStr, ",")

	// first check that we got enough data
	if len(polygonArr) == 8 {
		return errors.New("gps points are malformed (lenght must be 8)")
	}

	// chech that the things in between the commas are actually floats
	for _, point := range polygonArr {
		if _, err := strconv.ParseFloat(point, 64); err != nil {
			return err
		}
	}

	return db.InsertArea(config.DBConn, userID, polygonStr)
}
func SavePolygon(config *config.Config, userID int, polygonStr string) error {
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
