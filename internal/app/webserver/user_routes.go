package webserver

import (
	"fmt"
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

	var userID, contactAllowed, err = users.Login(config, email, password)
	if err != nil {
		c.Writer.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
		c.JSON(
			http.StatusForbidden,
			gin.H{
				"error":     "unable to login",
				"raw_error": err.Error(),
			},
		)
		return
	}

	// success
	var session = sessions.Default(c)
	session.Set("user_id", userID)
	session.Set("contact_allowed", contactAllowed)
	session.Save()
	c.Request.SetBasicAuth(email, password)
	c.SetCookie("session_id", session.ID(), 3600, "/", config.HTTPAddr, true, true)
	c.SetCookie("user_id", fmt.Sprintf("%d", userID), 3600, "/", config.HTTPAddr, true, true)
	c.JSON(
		http.StatusOK,
		gin.H{
			"id":              userID,
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
				"error":     "unable to parse contact_allowed as bool: " + contactAllowedStr,
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
	session.Set("contact_allowed", contactAllowed)
	session.Save()
	c.Request.SetBasicAuth(email, password)
	c.SetCookie("session_id", session.ID(), 3600, "/", config.HTTPAddr, true, true)
	c.SetCookie("user_id", fmt.Sprintf("%d", userID), 3600, "/", config.HTTPAddr, true, true)
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
