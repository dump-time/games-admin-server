package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
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
	server.BeforeBegin = func(add string) {
		// Saving pid info
		pid := os.Getpid()
		log.Printf("Actual pid is %d", pid)

		// TODO 这个地方的 ./dist 要使用配置文件来读取
		file, _ := os.OpenFile("./dist/pid.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
		defer file.Close()

		writer := bufio.NewWriter(file)
		_, _ = writer.WriteString(fmt.Sprintf("%v", pid))
		_ = writer.Flush()
	}
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
