package server

import (
	"log"
	"strconv"
	"time"

	"github.com/baldurstod/go-steam-assetclassinfo/api"
	"github.com/baldurstod/go-steam-assetclassinfo/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/secure"
	"github.com/gin-gonic/gin"
)

var ReleaseMode = "true"

func StartServer(config config.Config) {
	engine := initEngine(config)
	var err error

	log.Printf("Listening on port %d\n", config.HTTPS.Port)
	err = engine.RunTLS(":"+strconv.Itoa(config.HTTPS.Port), config.HttpsCertFile, config.HttpsKeyFile)
	log.Fatal(err)
}

func initEngine(config config.Config) *gin.Engine {
	if ReleaseMode == "true" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	r.SetTrustedProxies(nil)

	r.Use(cors.New(cors.Config{
		AllowMethods:    []string{"POST"},
		AllowHeaders:    []string{"Origin", "Content-Length", "Content-Type"},
		AllowAllOrigins: false,
		AllowOrigins:    config.AllowOrigins,
		MaxAge:          12 * time.Hour,
	}))

	r.Use(secure.New(secure.Config{
		SSLRedirect:           true,
		STSSeconds:            315360000,
		ContentSecurityPolicy: "default-src 'self'",
		ContentTypeNosniff:    true,
		ReferrerPolicy:        "strict-origin-when-cross-origin",
		SSLProxyHeaders:       map[string]string{"X-Forwarded-Proto": "https"},
	}))

	r.POST("/api", api.ApiHandler)

	return r
}
