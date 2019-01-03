package control

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/orm"
	"log"
	"myBlog/db"
	"strconv"
)

var START int
var END int

const (
	ArticleMinLength = 10
)


func AddArticle(a *db.Article) *db.Article {
	db.DB.Create(a)
	return a
}

func Detail(a *db.Article) *db.Article {
	db.DB.Select(a)
	return a
}

func filter(q *orm.Query) (*orm.Query, error) {
	q = q.Where(" limit ? ", END-START).Where(" offset ? ", START)
	return q, nil
}

func ArticlePage(c *gin.Context)  {

	START = c.GetInt("start")
	if START < 0 {
		START = 0
	}
	END = c.GetInt("end")
	if END < 1{
		END = 10
	}
	var articles []db.Article
	db.DB.Limit(END - START).Offset(START).Find(&articles)
	if len(articles) == 0 {
		articles = []db.Article{}
	}
	c.JSON(200,gin.H{
		"articles":articles,
	})
}

func ArticleAdd(c *gin.Context) {
	content := c.PostForm("content")
	title := c.PostForm("title")
	if len(content) < ArticleMinLength {
		c.JSON(100, gin.H{
			"message": "文章内容必须大于十个字符",
		})
	}
	userId, err := strconv.ParseInt(c.PostForm("user"), 10, 64)
	if err != nil {
		c.JSON(100, gin.H{
			"message": "userId 需要",
		})
	}
	// 先查询用户信息　
	user := &db.User{}
	db.DB.Find(user,userId)
	article := &db.Article{
		Author:*user,
		Content:content,
		AuthorId:userId,
		Title:title,
	}
	AddArticle(article)
	c.JSON(200,gin.H{
		"message":"ok",
		"article":article,
	})
}

func ArticleDetail(c *gin.Context){
	articleId,err := strconv.ParseInt( c.PostForm("articleId"),10,64 )
	if err != nil {
		c.JSON(100, gin.H{
			"message": "userId 需要",
		})
	}
	article := &db.Article{}
	db.DB.Model(article).Related("Author").Take(article,articleId)
	log.Println(article)
	c.JSON(200,gin.H{
		"article":article,
	})
}


