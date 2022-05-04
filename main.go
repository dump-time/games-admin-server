package main

import (
	"fmt"
	"lixiao189/games-admin-server/global"
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
			fmt.Printf("[INFO] Deamon started: %v\n", pid)
		}
	}

	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
