package router

import (
	"github.com/gin-gonic/gin"
	"myBlog/control"
)

func R (e * gin.Engine) {
	// 前端文件
	e.Static("/static","./front")
	e.GET("/ping",control.Ping)
	apiRouter(e)
}


func apiRouter (e *gin.Engine) {
	api := e.Group("/api")
	{
		api.POST("/ArticleAdd",control.Add)
		api.POST("/ArticleDetail",control.ArticleDetail)
	}
}