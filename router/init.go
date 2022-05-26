package router

import (
	"os"

	"github.com/dump-time/games-admin-server/global"
	"github.com/dump-time/games-admin-server/log"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"

	"github.com/gin-gonic/gin"
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
	R.Use(gin.LoggerWithFormatter(log.Formatter))

	// Panic auto recovery & return 500
	R.Use(gin.Recovery())

	// Add session middleware
	redisHost := global.Config.Redis.Hostname
	redisPort := global.Config.Redis.Port
	redisPass := global.Config.Redis.Password
	redisSecret := global.Config.Redis.Secret
	store, err := redis.NewStore(10, "tcp", redisHost+":"+redisPort, redisPass, []byte(redisSecret))
	if err != nil {
		log.Fatal("Loading redis error")
		log.Fatal(err)
		os.Exit(-1)
	}
	R.Use(sessions.Sessions("admin-server-session", store))

	// Setup routers
	v1 := R.Group("/api/v1")
	basicAPI := initBasicRouter(v1)
	teamAPI := initTeamRouter(basicAPI)
	{
		initVolunteerRouter(teamAPI)
		initJobRouter(teamAPI)
	}
}
