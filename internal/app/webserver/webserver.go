package webserver

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/secure"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/postgres"
	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/kmulvey/trashmap/internal/app/config"
	log "github.com/sirupsen/logrus"
)

const runLocal = "http://localhost"

func StartWebServer(config *config.Config) error {
	var router = gin.Default()

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())

	// compress
	router.Use(gzip.Gzip(gzip.DefaultCompression))

	// CORS
	var corsConfig = cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{config.HTTPAddr}
	router.Use(cors.New(corsConfig))

	// secure headers
	router.Use(secure.New(secure.Config{
		//		AllowedHosts:          []string{config.HTTPAddr},
		//		SSLRedirect:           true,
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

	// session
	var sqlDB = stdlib.OpenDB(*config.DBConn.Config())
	store, err := postgres.NewStore(sqlDB, []byte(config.PasswordSalt))
	if err != nil {
		return err
	}
	// auth'd routes
	router.Use(sessions.Sessions("web-session", store))
	router.DELETE("/user/:id", IsLoggedIn, func(c *gin.Context) { DeleteUser(config, c) })
	router.POST("/areas", IsLoggedIn, func(c *gin.Context) { GetPickupAreasWithinArea(config, c) })
	router.PUT("/area", IsLoggedIn, func(c *gin.Context) { CreatePickupArea(config, c) })

	// open routes
	router.StaticFS("/assets", http.Dir("./web"))
	router.POST("/login", func(c *gin.Context) { Login(config, c) })
	router.PUT("/user", func(c *gin.Context) { CreateUser(config, c) })

	if config.HTTPAddr == runLocal {
		log.Fatal(router.RunTLS(":8000", "./keys/cert.pem", "./keys/key.pem"))
	} else {
		log.Fatal(autotls.Run(router, config.HTTPAddr))
	}

	return nil
}
