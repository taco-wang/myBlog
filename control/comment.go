package control

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"myBlog/db"
	"myBlog/state"
	"strconv"
)

func CommentAdd(c *gin.Context) {
	articleId, _ := strconv.ParseUint(c.PostForm("articleId"), 10, 64)
	comment := c.PostForm("comment")
	commenterId, _ := strconv.ParseUint(c.PostForm("commenterId"), 10, 64)
	//db.DB.Model(&commenter).Take(commenterId)
	comments := &db.Comment{
		Comment: comment,
		Commenter: db.User{
			Model: gorm.Model{ID: uint(commenterId),},
		},
		Article: db.Article{
			Model: gorm.Model{ID: uint(articleId),},
		},
	}
	db.DB.Create(&comments)
	c.JSON(state.HTTPOKCODE, gin.H{
		comment: comments,
	})
}

func CommentDetail(c *gin.Context){
	commentId := c.GetInt("commentId")
	comment :=db.Comment{
		Model:gorm.Model{ID:uint(commentId),},
	}
	db.DB.First(&comment)
	c.JSON(state.HTTPOKCODE, gin.H{
		"comment": comment,
	})
}
