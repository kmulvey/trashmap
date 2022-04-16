package trashapp

import (
	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
	"github.com/kmulvey/trashmap/internal/app/config"
	log "github.com/sirupsen/logrus"
)

func startWebServer(config *config.Config, runLocal bool) {
	r := gin.Default()
	r.PUT("/user", func(c *gin.Context) { CreateUser(config, c) })
	r.DELETE("/user", func(c *gin.Context) { DeleteUser(config, c) })

	if runLocal {
		log.Fatal(r.Run(config.HTTPAddr))
	} else {
		log.Fatal(autotls.Run(r, "example1.com", "example2.com")) // TODO: real hostname .. https://github.com/kmulvey/trashmap/issues/1
	}
}
