package control

import (
	"github.com/gin-gonic/gin"
	"myBlog/db"
)

func CommentAdd(c *gin.Context){
	articleId := c.PostForm("articleId")
	comment := c.PostForm("comment")
	commenterId := c.PostForm("commenterId")
	var commenter db.User
	var article db.Article
	db.DB.Model(&commenter).Take(commenterId)
	db.DB.Model()
}
