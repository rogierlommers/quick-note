package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rogierlommers/quick-note/backend/api"
	cfg "github.com/rogierlommers/quick-note/backend/config"
	"github.com/rogierlommers/quick-note/backend/mailer"
)

func main() {

	// read config and make globally available
	cfg.ReadConfig()

	// gin mode
	if cfg.Settings.Mode == "PRO" || cfg.Settings.Mode == "PRODUCTION" {
		log.Println("enabling production mode")
		gin.SetMode(gin.ReleaseMode)
	}

	// create mailer instance
	m := mailer.NewMailer()

	// create router
	router := gin.New()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "PATCH"},
		AllowHeaders:     []string{"Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers"},
		ExposeHeaders:    []string{"Content-Length"},
		MaxAge:           12 * time.Hour,
		AllowCredentials: true,
	}))

	// add routers
	api.AddRoutes(router, m)

	// start serving
	if err := http.ListenAndServe(":3000", router); err != nil {
		log.Fatal(err)
	}

}
