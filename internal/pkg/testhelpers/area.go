package testhelpers

import "github.com/kmulvey/trashmap/internal/pkg/gps"

func GetRandomArea() *gps.Area {
	var nwPoint = gps.Coordinate{
		Lat:  GetRandomFloat(40.0, 50.0),
		Long: GetRandomFloat(-110.0, -100.0),
	}
	var a = gps.Area{
		Coords: []*gps.Coordinate{
			&nwPoint, //NW
			{Lat: nwPoint.Lat, Long: nwPoint.Long + 1},     // NE
			{Lat: nwPoint.Lat - 1, Long: nwPoint.Long + 1}, // SE
			{Lat: nwPoint.Lat - 1, Long: nwPoint.Long},     //SW
		},
	}

	return &a
}
