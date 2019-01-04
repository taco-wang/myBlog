package main

import (
	"github.com/gin-gonic/gin"
	_ "myBlog/db"
	"myBlog/router"
	"net/http"
)

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		// 加入跨域请求头
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
		c.Header("Access-Control-Allow-Methods", "GET, OPTIONS, POST, PUT, DELETE")
		c.Set("content-type", "application/json")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}

		c.Next()
	}
}

func main() {
	r := gin.Default()
	// 允许跨域
	r.Use(CorsMiddleware())
	// 初始化路由
	router.R(r)
	r.Run(":9999")
}
