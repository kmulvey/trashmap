package trashapp

import (
	"log"

	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
	"github.com/kmulvey/trashmap/internal/app/config"
)

func startWebServer(config *config.Config, runLocal bool) {
	r := gin.Default()
	r.PUT("/user", func(c *gin.Context) { CreateUser(config, c) })
	r.DELETE("/user", DeleteUser)
	r.POST("/user", UpdateUser)

	if runLocal {
		r.Run(addr)
	} else {
		log.Fatal(autotls.Run(r, "example1.com", "example2.com")) // TODO: real hostname .. https://github.com/kmulvey/trashmap/issues/1
	}
}
