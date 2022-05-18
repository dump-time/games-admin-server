package log

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Info(v interface{}) {
	log.Println("[\033[1;32mINFO\033[0m] ", v)
}

func Warn(v interface{}) {
	log.Println("[\033[1;33mWARN\033[0m] ", v)
}

func Error(v interface{}) {
	log.Println("[\033[1;31mERROR\033[0m] ", v)
}

func Fatal(v interface{}) {
	log.Fatalln("[\033[1;30;41mFATAL\033[0m] ", v)
}

func Formatter(param gin.LogFormatterParams) string {
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
}
