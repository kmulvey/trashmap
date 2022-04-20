# Trash Map
[![TrashMap](https://github.com/kmulvey/trashmap/actions/workflows/release_build.yml/badge.svg)](https://github.com/kmulvey/trashmap/actions/workflows/release_build.yml) [![Stand With Ukraine](https://raw.githubusercontent.com/vshymanskyy/StandWithUkraine/main/badges/StandWithUkraine.svg)](https://vshymanskyy.github.io/StandWithUkraine)

## Run Locally
```
cd deployments/local/
docker-compose up
```
pgAdmin will come up on http://localhost:8080/browser/

## HTTP routes
| Path       | Method  | Args                                                       | Description |
|------------|---------|------------------------------------------------------------|--------------
| /login     | POST    | email, password                                            | logs the user in
| /user      | PUT     | email, password, contact_allowed                           | creates the use acct
| /user/:id  | DELETE  | user id                                                    | deletes the user
| /area      | PUT     | gps corrdinate json (see below)                            | creates a new pickup area
| /areas     | POST    | gps corrdinate json (see below), must be four points       | returns all the pickup areas within the given area

## Data format
### [gps.Coordinate](https://github.com/kmulvey/trashmap/blob/main/internal/pkg/gps/gps.go#L10)
```
{
  "lat": 40.259822802779816,
  "long": -105.65290936674407
}
```
### [gps.Area](https://github.com/kmulvey/trashmap/blob/main/internal/pkg/gps/area.go#L9)
```
[
  {
    "lat": 40.259822802779816,
    "long": -105.65290936674407
  },
  {
    "lat": 40.26201885227386,
    "long": -105.05519389236237
  }
]
```
### [gps.Map](https://github.com/kmulvey/trashmap/blob/main/internal/pkg/gps/area.go#L8)
```
[
  {
    "Coords": [
      {
        "lat": 40.259822802779816,
        "long": -105.65290936674407
      },
      {
        "lat": 40.26201885227386,
        "long": -105.05519389236237
      },
      {
        "lat": 40.259822802779816,
        "long": -105.65290936674407
      }
    ]
  },
  {
    "Coords": [
      {
        "lat": 39.95833557541779,
        "long": -105.05494458234654
      },
      {
        "lat": 39.93788639054093,
        "long": -105.68899269947714
      },
      {
        "lat": 39.95833557541779,
        "long": -105.05494458234654
      }
    ]
  }
]
```
