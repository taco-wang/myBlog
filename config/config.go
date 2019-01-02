package config

import (
	"github.com/jinzhu/configor"
)

var Config struct{
	APPName string `default:"app name"`

	DB struct {
		Name     string
		User     string `default:"root"`
		Password string `required:"true" env:"DBPassword"`
		Port     uint   `default:"3306"`
		Database string `default:"postgres"`
	}
}

func init(){
	configor.Load(&Config,"./config/config.yml")
}