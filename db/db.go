package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"myBlog/config"
)
var TABLES = []interface{} {
	&User{},&Article{},&Comment{},&Block{},
}

type User struct {
	gorm.Model
	Name   string
	Email string
	Password string
}

func (u User) String() string {
	return fmt.Sprintf("User<%d %s %v>", u.ID, u.Name, u.Email)
}

type Article struct {
	gorm.Model
	Title    string
	AuthorId int64
	Author   User
	Content  string
	Block    Block
}

func (s Article) String() string {
	return fmt.Sprintf("Story<%d %s %s>", s.ID, s.Title, s.Author)
}

type Comment struct {
	gorm.Model
	Comment   string
	Commenter User
	Article Article
}

func (s Comment) String() string {
	return fmt.Sprintf("Comment<%d %s %s %s>", s.ID, s.Comment, s.Commenter)
}

type Block struct {
	gorm.Model
	Name string
}

func (s Block) String() string {
	return fmt.Sprintf("Block<%d %s %s %s>", s.ID, s.Name)
}

var DB *gorm.DB

func init() {
	var err error
	DB,err = gorm.Open("postgres", config.Config.DB.Connectstring)
	if err != nil {panic(err)}
	migrate(DB)
	DB.LogMode(config.Config.DB.Log)
}

func migrate(db *gorm.DB) {
	for _,model := range TABLES {
			db.AutoMigrate(model)
	}
}

