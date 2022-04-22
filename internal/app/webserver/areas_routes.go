package webserver

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/kmulvey/trashmap/internal/app/areas"
	"github.com/kmulvey/trashmap/internal/app/config"
	"github.com/kmulvey/trashmap/internal/pkg/gps"
)

// CreatePickupArea handler takes a POST'd gps string and
// adds it to the database.
func CreatePickupArea(config *config.Config, c *gin.Context) {
	var areaStr = c.PostForm("area")
	var session = sessions.Default(c)
	var userIDIFace = session.Get("user_id")
	fmt.Println(reflect.TypeOf(userIDIFace))
	var userID, ok = userIDIFace.(int64)
	if !ok {
		sendJSONError(c, http.StatusInternalServerError, "unable to user get user_id from session", nil)
		return
	}

	var polygon, err = gps.NewAreaFromJSONString(areaStr)
	if err != nil {
		sendJSONError(c, http.StatusBadRequest, "unable to unmarshal gps json data", err)
		return
	}

	id, err := areas.SaveArea(config, userID, polygon)
	if err != nil {
		sendJSONError(c, http.StatusInternalServerError, "unable to save area", err)
		return
	}
	c.JSON(
		http.StatusOK,
		gin.H{
			"id": id,
		},
	)
}

// GetPickupAreasWithinArea handler takes a POST'd gps string and
// adds returns all the pickup areas within that area.
func GetPickupAreasWithinArea(config *config.Config, c *gin.Context) {
	var areaStr = c.PostForm("area")
	var polygon, err = gps.NewAreaFromJSONString(areaStr)
	if err != nil {
		sendJSONError(c, http.StatusBadRequest, "unable to unmarshal gps json data", err)
		return
	}

	pickupAreas, err := areas.GetPickupAreasWithinArea(config, polygon)
	if err != nil {
		sendJSONError(c, http.StatusInternalServerError, "unable to get areas from db", err)
		return
	}

	pickupAreasJSON, err := pickupAreas.ToJSON()
	if err != nil {
		sendJSONError(c, http.StatusInternalServerError, "unable to marshal areas to JSON", err)
		return
	}

	// all good
	c.JSON(
		http.StatusOK,
		gin.H{
			"pickup_areas": string(pickupAreasJSON),
		},
	)
}
