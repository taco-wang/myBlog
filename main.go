package main

import (
	"github.com/gin-gonic/gin"
	_ "myBlog/db"
	"myBlog/router"
)

func main() {
	r := gin.Default()
	// 初始化路由
	router.R(r)
	r.Run(":9999")
}
