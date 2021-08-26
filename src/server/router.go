package server

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"stockify-api/src/handlers"
	"time"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	//router.Use(gin.Logger())
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST, OPTIONS, GET, PUT"},
		AllowHeaders:     []string{"Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With"},
		AllowCredentials: true,
		AllowWildcard:    true,
		MaxAge:           12 * time.Hour,
	}))

	router.Use(gin.Recovery())

	health := new(handlers.HealthHandler)

	router.GET("/health", health.Status)

//	router.Use(middlewares.AuthMiddleware())

	user := handlers.NewUserHandler()

	router.POST("/addUser", user.AddUser)

	alert := handlers.NewAlertHandler()

	router.POST("/addAlert", alert.AddAlert)

	router.GET("/removeAlert/:alertID", alert.RemoveAlert)

	router.GET("/alertList/:userID", alert.AlertList)

	metadata := handlers.NewMetadataHandler()

	router.GET("/reloadInstruments", metadata.ReloadInstruments)

	router.GET("/trendingList", metadata.TrendingInstruments)

	router.GET("/instrumentDetails/:exchange/:symbol", metadata.FetchInstrumentDetails)

	return router

}
