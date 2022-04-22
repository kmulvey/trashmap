package webserver

import "github.com/gin-gonic/gin"

func sendJSONError(c *gin.Context, status int, errorContext string, err error) {
	var errStr string
	if err != nil {
		errStr = err.Error()
	}
	c.JSON(
		status,
		gin.H{
			"error":     errorContext,
			"raw_error": errStr,
		},
	)
}
