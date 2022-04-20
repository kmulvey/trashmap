# Trash Map
[![TrashMap](https://github.com/kmulvey/trashmap/actions/workflows/release_build.yml/badge.svg)](https://github.com/kmulvey/trashmap/actions/workflows/release_build.yml) [![Stand With Ukraine](https://raw.githubusercontent.com/vshymanskyy/StandWithUkraine/main/badges/StandWithUkraine.svg)](https://vshymanskyy.github.io/StandWithUkraine)

## HTTP routes
| Path       | Method  | Args                                                       | Description |
|------------|---------|------------------------------------------------------------|--------------
| /login     | POST    | email, password                                            | logs the user in
| /user      | PUT     | email, password, contact_allowed                           | creates the use acct
| /user/:id  | DELETE  | user id                                                    | deletes the user
| /area      | PUT     | gps corrdinate json (see below)                            | creates a new pickup area
| /areas     | POST    | gps corrdinate json (see below), must be four points       | returns all the pickup areas within the given area
