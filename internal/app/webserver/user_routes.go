package webserver

import (
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/kmulvey/trashmap/internal/app/config"
	"github.com/kmulvey/trashmap/internal/app/users"
)

func Login(config *config.Config, c *gin.Context) {
	var email = c.PostForm("email")
	var password = c.PostForm("password")

	var id, contactAllowed, err = users.Login(config, email, password)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error":     "unable to login",
				"raw_error": err.Error(),
			},
		)
		return
	}

	// success
	var session = sessions.Default(c)
	session.Set("user_id", id)
	c.JSON(
		http.StatusOK,
		gin.H{
			"id":              id,
			"contact_allowed": contactAllowed,
		},
	)
}

func CreateUser(config *config.Config, c *gin.Context) {
	var email = c.PostForm("email")
	var password = c.PostForm("password")
	var contactAllowedStr = c.PostForm("contact_allowed")
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

	userID, err := users.Add(config, email, password, contactAllowed)
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

	var session = sessions.Default(c)
	session.Set("user_id", userID)

	c.Status(http.StatusOK)
}

func DeleteUser(config *config.Config, c *gin.Context) {
	var email = c.Param("email")
	var err = users.Remove(config, email)
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

	// delete their session
	var session = sessions.Default(c)
	session.Delete(session.Get("user_id"))

	c.Status(http.StatusOK)
}
