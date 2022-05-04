package main

import (
	"fmt"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"log"
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
	server := endless.NewServer(":8083", router)

	// daemon mode
	if *daemonMode {
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
