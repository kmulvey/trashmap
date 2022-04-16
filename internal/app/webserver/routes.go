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
			http.StatusOK,
			gin.H{
				"code":      http.StatusBadRequest,
				"error":     "unable to ParseBool: " + contactAllowedStr,
				"raw_error": err.Error(),
			},
		)
	}
	err = users.AddUser(config, email, contactAllowed)
	if err != nil {
		c.JSON(
			http.StatusOK,
			gin.H{
				"code":      http.StatusInternalServerError,
				"error":     "unable to add user",
				"raw_error": err.Error(),
			},
		)
	}
}
func UpdateUser(*gin.Context) {

}
func DeleteUser(*gin.Context) {

}
