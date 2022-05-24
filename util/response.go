package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func standardResp(context *gin.Context, code int, msg string, data interface{}) {
	context.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

// SuccessResp the successful response
func SuccessResp(context *gin.Context, data interface{}) {
	standardResp(context, 0, "ok", data)
}

// FailedResp the failed response
func FailedResp(context *gin.Context, code int, msg string) {
	standardResp(context, code, msg, nil)
}

func ParamsErrResp(context *gin.Context) {
	FailedResp(context, 4001, "parameter error!")
}

func NotAllowedResp(context *gin.Context) {
	FailedResp(context, 4002, "not allow!")
}


func NotLoginResp(context *gin.Context) {
	FailedResp(context, 4003, "not login!")
}

func NotFoundResp(context *gin.Context) {
	FailedResp(context, 404, "Not Found")
}
