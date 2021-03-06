package gps

import (
	"encoding/json"
	"strings"
)

type Map []*Area
type Area struct {
	Coords []*Coordinate
}

func NewAreaFromJSONString(areaStr string) (*Area, error) {
	var area = new(Area)
	var err = json.Unmarshal([]byte(areaStr), &area.Coords)
	if err != nil {
		return nil, err
	}

	// enforce postgis format
	if !area.Validate() {
		area.Coords = append(area.Coords, area.Coords[0])
	}

	return area, nil
}

func NewAreaFromPostGISString(areaStr string) (*Area, error) {
	var areaSplit = strings.Split(strings.TrimSpace(areaStr), ",")
	var area = new(Area)
	area.Coords = make([]*Coordinate, len(areaSplit))

	for i, a := range areaSplit {
		var err error
		area.Coords[i], err = NewCoordinateFromPostGISString(strings.TrimSpace(a))
		if err != nil {
			return nil, err
		}
	}

	// enforce postgis format
	if !area.Validate() {
		area.Coords = append(area.Coords, area.Coords[0])
	}

	return area, nil
}

// Validate ensures that the first and last coordinates are the same,
// just the way postgis likes it
func (a *Area) Validate() bool {
	var lastIndex = len(a.Coords) - 1
	if a.Coords[0].Lat == a.Coords[lastIndex].Lat && a.Coords[0].Long == a.Coords[lastIndex].Long {
		return true
	}
	return false
}

// CoordinatesToPostGISString prints the Coordinate space separated for postgis
func (a *Area) CoordinatesToPostGISString() string {
	var builder = strings.Builder{}

	for i, coord := range a.Coords {
		if i != 0 {
			builder.WriteString(",")
		}
		builder.WriteString(coord.ToPostGISString())
	}
	return builder.String()
}

// ToJSON removes the final point from the area
// returns the resulting JSON. This is done because
// the front end does not need the duplicate point as
// postgis does
func (a *Area) ToJSON() (string, error) {
	var coords = a.Coords[:len(a.Coords)-1]

	var js, err = json.Marshal(coords)
	return string(js), err
}

func (m *Map) ToJSON() (string, error) {
	var js, err = json.Marshal(m)
	return string(js), err
}
