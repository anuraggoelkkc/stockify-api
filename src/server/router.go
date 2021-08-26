package server

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"stockify-api/src/handlers"
	"stockify-api/src/middlewares"
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

	router.Use(middlewares.AuthMiddleware())

	user := new(handlers.UserHandler)

	router.POST("/addUser", user.AddUser)

	alert := new(handlers.AlertHandler)

	router.POST("/addAlert", alert.AddAlert)

	router.GET("/removeAlert/:userID", alert.RemoveAlert)

	router.GET("/alertList/:userID", alert.AlertList)

	metadata := new(handlers.MetadataHandler)

	router.GET("/reloadInstruments", metadata.ReloadInstruments)

	router.GET("/trendingList", metadata.TrendingInstruments)

	return router

}
