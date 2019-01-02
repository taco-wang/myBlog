package db

import (
	"fmt"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"myBlog/config"
)

type User struct {
	Id     int64
	Name   string
	Emails []string
}

func (u User) String() string {
	return fmt.Sprintf("User<%d %s %v>", u.Id, u.Name, u.Emails)
}

type Article struct {
	Id       int64
	Title    string
	AuthorId int64
	Author   *User
	Content  string
	Block    *Block
}

func (s Article) String() string {
	return fmt.Sprintf("Story<%d %s %s>", s.Id, s.Title, s.Author)
}

type Comment struct {
	Id        int64
	Author    *User
	Comment   string
	Commenter *User
}

func (s Comment) String() string {
	return fmt.Sprintf("Comment<%d %s %s %s>", s.Id, s.Comment, s.Author, s.Commenter)
}

type Block struct {
	Id   int64
	Name string
}

func (s Block) String() string {
	return fmt.Sprintf("Block<%d %s %s %s>", s.Id, s.Name)
}

var DB *pg.DB

func init() {
	DB = pg.Connect(&pg.Options{
		User:     config.Config.DB.User,
		Password:  config.Config.DB.Password,
		Database: config.Config.DB.Database,
	})
	err:= createSchema(DB)
	if err != nil {
		panic(err)
	}
}

func createSchema(db *pg.DB) error {
	for _, model := range []interface{}{
		(*User)(nil), (*Article)(nil), (*Comment)(nil), (*Block)(nil),
	} {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			Temp: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
