package webserver

import (
	"net/http"

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
	var userID, ok = userIDIFace.(int64)
	if !ok {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error":     "unable to user get user_id from session",
				"raw_error": "",
			},
		)
		return
	}

	var polygon, err = gps.NewAreaFromJSONString(areaStr)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error":     "unable to unmarshal gps json data",
				"raw_error": err.Error(),
			},
		)
		return
	}

	id, err := areas.SaveArea(config, userID, polygon)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error":     "unable to save area",
				"raw_error": err.Error(),
			},
		)
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
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error":     "unable to unmarshal gps json data",
				"raw_error": err.Error(),
			},
		)
		return
	}

	pickupAreas, err := areas.GetPickupAreasWithinArea(config, polygon)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error":     "unable to get areas from db",
				"raw_error": err.Error(),
			},
		)
		return
	}

	pickupAreasJSON, err := pickupAreas.ToJSON()
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error":     "unable to marshal areas to JSON",
				"raw_error": err.Error(),
			},
		)
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
