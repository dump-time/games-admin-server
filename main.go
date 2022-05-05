package main

import (
	"fmt"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"lixiao189/games-admin-server/global"
	"lixiao189/games-admin-server/log"
	"net"
	"os"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Start server gracefully
	server := endless.NewServer(global.Config.Serv.Addr, router)

	// daemon mode
	if *global.DaemonMode {
		server.BeforeBegin = func(add string) {
			// stdout pid
			pid := os.Getpid()
			log.Info(fmt.Sprintf("Deamon started: %v", pid))
		}
	}

	// Start server
	if err := server.ListenAndServe(); err != nil {
		switch err.(type) {
		case *net.OpError:
			log.Warn(err)
		default:
			log.Fatal(err)
		}
	}

}
