package webserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kmulvey/trashmap/internal/app/config"
	"github.com/kmulvey/trashmap/internal/app/polygon"
)

func CreatePolygon(config *config.Config, c *gin.Context) {
	var polygonStr = c.PostForm("polygon")

	var err = polygon.SavePolygon(config, 0, polygonStr)
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
	c.Status(http.StatusOK)
}
