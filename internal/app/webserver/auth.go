package webserver

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func IsLoggedIn(c *gin.Context) {
	var session = sessions.Default(c)
	var sessionUserIDInt, ok = session.Get("user_id").(int64)
	if !ok {
		sendJSONError(c, http.StatusForbidden, "", nil)
		c.Abort()
		return
	}

	if session.ID() == "" && sessionUserIDInt == 0 {
		c.Status(http.StatusForbidden)
		c.Writer.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
		c.Abort()
		return
	}
}
