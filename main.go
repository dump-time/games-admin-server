package main

import (
	"bufio"
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
			// Saving pid info
			pid := os.Getpid()
			log.Printf("Actual pid is %d", pid)

			// TODO 这个地方的 ./dist 要使用配置文件来读取
			file, err := os.OpenFile("./dist/pid.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
			if err != nil {
				fmt.Println(err.Error())
			}
			defer func(file *os.File) {
				err := file.Close()
				if err != nil {
					// TODO
				}
			}(file)

			writer := bufio.NewWriter(file)
			_, _ = writer.WriteString(fmt.Sprintf("%v", pid))
			_ = writer.Flush()
		}
	}

	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
