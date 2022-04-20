package main

import (
	"flag"

	"github.com/kmulvey/trashmap/internal/app/config"
	"github.com/kmulvey/trashmap/internal/app/webserver"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	var runLocal bool
	flag.BoolVar(&runLocal, "run-local", false, "run the webserver locally (no https)")
	flag.Parse()

	var config, err = config.NewConfig("", "", "", "", 0, "", "")
	if err != nil {
		log.Fatal(err)
	}

	webserver.StartWebServer(config, runLocal)
}
