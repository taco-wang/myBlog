package config

import (
	"github.com/jinzhu/configor"
)

var Config struct {
	APPName string `default:"app name"`

	DB struct {
		Name          string
		User          string `default:"root"`
		Password      string `required:"true" env:"DBPassword"`
		Port          uint   `default:"3306"`
		Database      string `default:"postgres"`
		Addr          string `default:"127.0.0.1"`
		Connectstring string ``
		Log           bool   `default:true`
	}
}

func init() {
	configor.Load(&Config, "./config/config.yml")
}
