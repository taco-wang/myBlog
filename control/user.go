package control

import (
	"github.com/gin-gonic/gin"
	"myBlog/db"
)

func UserAdd(c *gin.Context) {
	nickName := c.DefaultPostForm("nickName", "asynmous")
	email:= c.PostForm("email")
	password := c.PostForm("password")
	//var user User
	user := db.User{
		Name:nickName,
		Email:email,
		Password:password,
	}

	db.DB.Create(&user)
	c.JSON(200,gin.H{
		"userId": user.ID,
		"user": user,
	})
}

func UserDetail(c *gin.Context) {
	id := c.GetInt64("userId")
	user:=db.User{}
	db.DB.Take(user,id)
	c.JSON(200,gin.H{
		"user":user,
	})
}
