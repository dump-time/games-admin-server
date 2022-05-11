package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"lixiao189/games-admin-server/global"
	"lixiao189/games-admin-server/log"
	"time"
)

var R *gin.Engine

// init global router
func init() {
	R = gin.New()

	// Setup trusted reverse proxies
	if err := R.SetTrustedProxies(global.Config.Serv.TrustedProxies); err != nil {
		log.Fatal(err)
		return
	}

	// Log formatter
	R.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s |%s %d %s| %s |%s %s %s %s | %s | size: %d | %s | %s\n",
			param.TimeStamp.Format(time.RFC1123),
			param.StatusCodeColor(),
			param.StatusCode,
			param.ResetColor(),
			param.ClientIP,
			param.MethodColor(),
			param.Method,
			param.ResetColor(),
			param.Path,
			param.Latency,
			param.BodySize,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	// Panic auto recovery & return 500
	R.Use(gin.Recovery())

	// Setup routers
	v1 := R.Group("/api/v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})

		// TODO: routers...
	}
}
