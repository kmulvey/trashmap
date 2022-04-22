package webserver

import "github.com/gin-gonic/gin"

func sendJSONError(c *gin.Context, status int, errorContext string, err error) {
	c.JSON(
		status,
		gin.H{
			"error":     errorContext,
			"raw_error": err.Error(),
		},
	)
}
