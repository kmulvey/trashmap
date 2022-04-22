package testhelpers

import "github.com/kmulvey/trashmap/internal/pkg/gps"

func GetRandomCoordinate() *gps.Coordinate {
	return &gps.Coordinate{
		Lat:  GetRandomFloat(40.0, 50.0),
		Long: GetRandomFloat(-110.0, -100.0),
	}
}
