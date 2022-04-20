package webserver

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/kmulvey/trashmap/internal/app/areas"
	"github.com/kmulvey/trashmap/internal/app/config"
)

// CreatePickupArea handler takes a POST'd gps string and
// adds it to the database.
func CreatePickupArea(config *config.Config, c *gin.Context) {
	var polygonStr = c.PostForm("polygon")
	var session = sessions.Default(c)
	var userIDIFace = session.Get("user_id")
	var userID, ok = userIDIFace.(int)
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

	var id, err = areas.SaveArea(config, userID, polygonStr)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error":     "unable to remove user",
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
