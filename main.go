package main

import (
	"fmt"
	"lixiao189/games-admin-server/global"
	"lixiao189/games-admin-server/util"
	"log"
	"os"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initial functions 
	global.InitFlag()

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Start server gracefully
	server := endless.NewServer(":8083", router)

	// daemon mode
	if *global.DaemonMode {
		server.BeforeBegin = func(add string) {
			// stdout pid
			pid := os.Getpid()
			util.InfoLog(fmt.Sprintf("Deamon started: %v", pid))	
		}
	}

	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
