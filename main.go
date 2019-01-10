package main

import (
	"github.com/gin-gonic/gin"
	_ "myBlog/db"
	"myBlog/middlewares"
	"myBlog/router"
)



func main() {
	r := gin.Default()
	// 允许跨域
	r.Use(middlewares.CorsMiddleware())
	// 初始化路由
	router.R(r)
	r.Run(":9999")
}
