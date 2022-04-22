package webserver

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func IsLoggedIn(c *gin.Context) {
	var session = sessions.Default(c)
	var sessionUserIDInt, ok = session.Get("user_id").(int64)
	if !ok {
		sendJSONError(c, http.StatusForbidden, "unable to parse user_id from session", nil)
		c.Abort()
		return
	}

	var cookieSessionID, _ = c.Cookie("session_id")
	cookieUserID, err := c.Cookie("user_id")
	if err != nil {
		sendJSONError(c, http.StatusBadRequest, "unable to parse user_id from cookie", err)
		c.Abort()
		return
	}

	if cookieSessionID == session.ID() && cookieUserID == fmt.Sprintf("%d", sessionUserIDInt) {
	} else {
		c.Status(http.StatusForbidden)
		c.Writer.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
		c.Abort()
		return
	}
}
