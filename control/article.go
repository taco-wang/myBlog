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

type Article interface {
	Add(article *db.Article) *db.Article
	Detail(article *db.Article) *db.Article
	ArticlePage(q *orm.Query) []*db.Article
}

type MyArticle struct {
	db.Article
}

func AddArticle(a *db.Article) *db.Article {
	err := db.DB.Insert(a)
	if err != nil {
		log.Println("error", err)
	}
	return a
}

func Detail(a *db.Article) *db.Article {
	err := db.DB.Select(a)
	if err != nil {
		log.Println("error", err)
	}
	return a
}

func filter(q *orm.Query) (*orm.Query, error) {
	q = q.Where(" limit ", END-START).Where(" offset ", START)
	return q, nil
}

func ArticlePage(a db.Article, start int, end int) []db.Article {
	START = start
	END = end
	var articles []db.Article
	err := db.DB.Model(&articles).Apply(filter).Select()
	if err != nil {
		log.Println("error", err)
	}
	return articles
}

func handleError(err error) error {
	if err != nil {
		return err
	}
	return nil
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
	user := &db.User{
		Id:userId,
	}
	db.DB.Select(user)
	article := &db.Article{
		Author:user,
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
	article := &db.Article{
		Id:articleId,
	}
	db.DB.Select(article)
	log.Println(article)
	c.JSON(200,gin.H{
		"article":article,
	})
}
