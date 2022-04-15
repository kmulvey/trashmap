package main

import (
	"log"

	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
)

func startWebServer(addr string, runLocal bool) {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	if runLocal {
		r.Run(addr)
	} else {
		log.Fatal(autotls.Run(r, "example1.com", "example2.com")) // TODO: real hostname .. #1
	}
}
