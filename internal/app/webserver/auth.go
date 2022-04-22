package webserver

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func IsLoggedIn(c *gin.Context) {
	var session = sessions.Default(c)
	fmt.Println(reflect.TypeOf(session.Get("user_id")))
	var sessionUserIDInt, ok = session.Get("user_id").(int64)
	if !ok {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "unable to parse user_id from session",
			},
		)
		c.Abort()
	}

	var cookieSessionID, _ = c.Cookie("session_id")
	cookieUserID, err := c.Cookie("user_id")
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "unable to parse user_id from cookie",
			},
		)
		c.Abort()
	}

	if cookieSessionID == session.ID() && cookieUserID == fmt.Sprintf("%d", sessionUserIDInt) {
	} else {
		c.Status(http.StatusForbidden)
		c.Writer.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
		c.Abort()
		return
	}
}
