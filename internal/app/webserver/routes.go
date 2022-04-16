package trashapp

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kmulvey/trashmap/internal/app/config"
	"github.com/kmulvey/trashmap/internal/app/users"
)

func CreateUser(config *config.Config, c *gin.Context) {
	var email = c.PostForm("email")
	var contactAllowedStr = c.PostForm("contactAllowed")
	var contactAllowed, err = strconv.ParseBool(contactAllowedStr)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error":     "unable to ParseBool: " + contactAllowedStr,
				"raw_error": err.Error(),
			},
		)
		return
	}

	err = users.AddUser(config, email, contactAllowed)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error":     "unable to add user",
				"raw_error": err.Error(),
			},
		)
		return
	}
	c.Status(http.StatusOK)
}

func DeleteUser(config *config.Config, c *gin.Context) {
	var email = c.PostForm("email")
	var err = users.RemoveUser(config, email)
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
