# Trash Map
[![TrashMap](https://github.com/kmulvey/trashmap/actions/workflows/release_build.yml/badge.svg)](https://github.com/kmulvey/trashmap/actions/workflows/release_build.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/kmulvey/trashmap)](https://goreportcard.com/report/github.com/kmulvey/trashmap) [![Stand With Ukraine](https://raw.githubusercontent.com/vshymanskyy/StandWithUkraine/main/badges/StandWithUkraine.svg)](https://vshymanskyy.github.io/StandWithUkraine) [![Go Reference](https://pkg.go.dev/badge/github.com/kmulvey/trashmap.svg)](https://pkg.go.dev/github.com/kmulvey/trashmap)

## Run Locally
```
cd deployments/local/
docker-compose up -d
cd ../../
go run cmd/trashmap/main.go
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
  },
  {
    "lat": 39.95833557541779,
    "long": -105.05494458234654
  },
    {
    "lat": 39.93788639054093,
    "long": -105.68899269947714
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
        "lat": 39.95833557541779,
        "long": -105.05494458234654
      },
        {
        "lat": 39.93788639054093,
        "long": -105.68899269947714
      }
    ]
  },
  {
    "Coords": [
      {
        "lat": 36.24244932613963,
        "long": -112.40547934312802
      },
      {
        "lat": 36.223086062103704,
        "long": -111.81857028617547
      },
      {
        "lat": 36.00588425331162,
        "long": -111.83171900919251
      },
        {
        "lat": 36.02265323852763,
        "long": -112.34279966186926
      }
    ]
  }
]
```
## curl examples
### Create User
```
curl --cookie-jar jar.txt -v -XPUT localhost:8000/user -F 'contact_allowed=true' -F 'email=me@gmail.com' -F 'password=password'
```
### Login
```
curl --cookie-jar jar.txt -v -XPOST localhost:8000/login -F 'email=me@gmail.com' -F 'password=password'
```
### Create Area
```
curl -b jar.txt -v -XPUT localhost:8000/area -F 'user_id=1' -F 'area=[{"lat":40.259822802779816,"long":-105.65290936674407},{"lat":40.26201885227386,"long":-105.05519389236237},{"lat":39.95833557541779,"long":-105.05494458234654},{"lat":39.93788639054093,"long":-105.68899269947714}]'
curl -b jar.txt -v -XPUT localhost:8000/area -F 'user_id=1' -F 'area=[{"lat":36.24244932613963,"long":-112.40547934312802},{"lat":36.223086062103704,"long":-111.81857028617547},{"lat":36.00588425331162,"long":-111.83171900919251},{"lat":36.02265323852763,"long":-112.34279966186926}]'
```
### Get all areas within a wider area
```
curl -b jar.txt -v -XPOST localhost:8000/areas -F 'area=[{"lat":40.99837454922104,"long":-109.05421673800787},{"lat":41.0004139327614,"long":-102.05758033526878},{"lat":37.015379641011805,"long":-109.05244682441632},{"lat":37.004114233165524,"long":-102.04615148062369}]'
```
