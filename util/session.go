package util

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func ContextSession(context *gin.Context) sessions.Session {
 	return sessions.Default(context)
}
