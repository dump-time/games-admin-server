package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddJob(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, nil)
}
