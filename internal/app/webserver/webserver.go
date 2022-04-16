package webserver

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/httpsign"
	"github.com/gin-contrib/httpsign/crypto"
	"github.com/gin-contrib/secure"
	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
	"github.com/kmulvey/trashmap/internal/app/config"
	log "github.com/sirupsen/logrus"
)

func StartWebServer(config *config.Config, runLocal bool) {
	var router = gin.Default()

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())

	// CORS
	if runLocal {
		router.Use(cors.Default())
	} else {
		var corsConfig = cors.DefaultConfig()
		corsConfig.AllowOrigins = []string{config.HTTPAddr}
		router.Use(cors.New(corsConfig))
	}

	// compress
	router.Use(gzip.Gzip(gzip.DefaultCompression))

	// secure things a bit
	router.Use(secure.New(secure.Config{
		AllowedHosts:          []string{config.HTTPAddr},
		SSLRedirect:           true,
		STSSeconds:            315360000,
		STSIncludeSubdomains:  true,
		FrameDeny:             true,
		ContentTypeNosniff:    true,
		BrowserXssFilter:      true,
		ContentSecurityPolicy: "default-src 'self'",
		IENoOpen:              true,
		ReferrerPolicy:        "strict-origin-when-cross-origin",
		SSLProxyHeaders:       map[string]string{"X-Forwarded-Proto": "https"},
	}))

	// HTTP signing
	var readKeyID = httpsign.KeyID("read")
	var writeKeyID = httpsign.KeyID("write")
	var secrets = httpsign.Secrets{
		readKeyID: &httpsign.Secret{
			Key:       config.HTTPReadSigningSecret,
			Algorithm: &crypto.HmacSha512{},
		},
		writeKeyID: &httpsign.Secret{
			Key:       config.HTTPWriteSigningSecret,
			Algorithm: &crypto.HmacSha512{},
		},
	}
	auth := httpsign.NewAuthenticator(secrets)
	router.Use(auth.Authenticated())

	// routes
	router.PUT("/user", func(c *gin.Context) { CreateUser(config, c) })
	router.DELETE("/user", func(c *gin.Context) { DeleteUser(config, c) })

	if runLocal {
		log.Fatal(router.Run(config.HTTPAddr))
	} else {
		log.Fatal(autotls.Run(router, config.HTTPAddr))
	}
}
