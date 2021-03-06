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

	var userID, contactAllowed, err = users.Login(config, email, password)
	if err != nil {
		c.Writer.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
		sendJSONError(c, http.StatusForbidden, "unable to login", err)
		return
	}

	// success
	handleSuccess(config, c, userID, contactAllowed, email, password)
}

func CreateUser(config *config.Config, c *gin.Context) {
	var email = c.PostForm("email")
	var password = c.PostForm("password")
	var contactAllowedStr = c.PostForm("contact_allowed")
	var contactAllowed, err = strconv.ParseBool(contactAllowedStr)
	if err != nil {
		sendJSONError(c, http.StatusBadRequest, "unable to parse contact_allowed as bool: "+contactAllowedStr, err)
		return
	}

	userID, err := users.Add(config, email, password, contactAllowed)
	if err != nil {
		sendJSONError(c, http.StatusInternalServerError, "unable to add user", err)
		return
	}

	// success
	handleSuccess(config, c, userID, contactAllowed, email, password)
}

func DeleteUser(config *config.Config, c *gin.Context) {
	var email = c.Param("email")
	var err = users.Remove(config, email)
	if err != nil {
		sendJSONError(c, http.StatusInternalServerError, "unable to remove user", err)
		return
	}

	// delete their session
	var session = sessions.Default(c)
	session.Delete(session.Get("user_id"))

	c.Status(http.StatusOK)
}

func handleSuccess(config *config.Config, c *gin.Context, userID int64, contactAllowed bool, email, password string) {
	var session = sessions.Default(c)
	session.Set("user_id", userID)
	session.Set("contact_allowed", contactAllowed)
	var err = session.Save()
	if err != nil {
		sendJSONError(c, http.StatusInternalServerError, "unable to save session", err)
		return
	}
	c.JSON(
		http.StatusOK,
		gin.H{
			"id":              userID,
			"contact_allowed": contactAllowed,
		},
	)
}
