package control

import (
	"github.com/gin-gonic/gin"
	"myBlog/state"
)

func Ping(context *gin.Context) {
	context.JSON(state.HTTPOKCODE, gin.H{
		"message": "ok",
	})
}
