package router

import (
	"github.com/gin-gonic/gin"
	"myBlog/control"
)

func R (e * gin.Engine) {
	// 前端文件
	e.Static("/static","./font")
	e.GET("/ping",control.Ping)
}