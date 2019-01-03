package db

import (
	"fmt"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
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
	Author   *User
	Content  string
	Block    *Block
}

func (s Article) String() string {
	return fmt.Sprintf("Story<%d %s %s>", s.ID, s.Title, s.Author)
}

type Comment struct {
	gorm.Model
	Author    *User
	Comment   string
	Commenter *User
}

func (s Comment) String() string {
	return fmt.Sprintf("Comment<%d %s %s %s>", s.ID, s.Comment, s.Author, s.Commenter)
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
	//DB = pg.Connect(&pg.Options{
	//	User:     config.Config.DB.User,
	//	Password:  config.Config.DB.Password,
	//	Database: config.Config.DB.Database,
	//})
	//err:= createSchema(DB)
	//if err != nil {
	//	panic(err)
	//}
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

func createSchema(db *pg.DB) error {
	for _, model := range TABLES {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			Temp: false,
			IfNotExists:true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
