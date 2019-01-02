package control

import (
	"github.com/gin-gonic/gin"
	"myBlog/db"
)

func UserAdd(c *gin.Context) {
	nickName := c.DefaultPostForm("nickName", "asynmous")
	email:= c.PostForm("email")
	password := c.PostForm("password")
	user:= &db.User{
		Name:nickName,
		Email:email,
		Password:password,
	}
	db.DB.Insert(user)
	c.JSON(200,gin.H{
		"userId": user.Id,
	})
}

func UserDetail(c *gin.Context) {
	id := c.GetInt64("userId")
	user:=db.User{
		Id:id,
	}
	db.DB.Select(user)
	c.JSON(200,gin.H{
		"user":user,
	})
}
