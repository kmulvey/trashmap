package gps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Coordinate struct {
	Lat  float64
	Long float64
}

func NewCoordinateFromPostGISString(coordinateStr string) (*Coordinate, error) {
	var coordinate = new(Coordinate)
	var err error

	var split = strings.Split(coordinateStr, " ")
	if len(split) != 2 {
		return nil, errors.New("unable to marshal coordinates from db, incomplete pair: " + coordinateStr)
	}

	coordinate.Lat, err = strconv.ParseFloat(split[0], 64)
	if err != nil {
		return nil, fmt.Errorf("unable to marshal coordinates from db, bad Lat: %s, error: %w", split[0], err)
	}

	coordinate.Long, err = strconv.ParseFloat(split[1], 64)
	if err != nil {
		return nil, fmt.Errorf("unable to marshal coordinates from db, bad  Long: %s, error: %w", split[1], err)
	}

	return coordinate, nil
}

// ToPostGISString prints the Coordinate space separated for postgis
func (c *Coordinate) ToPostGISString() string {
	return fmt.Sprintf("%f %f", c.Lat, c.Long)
}
