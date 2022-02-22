package config

import (
	"github.com/jinzhu/configor"
)

var Config = struct {
	BaseURL string `default:"http://localhost" yaml:"BaseURL"`
	Port    string `default:"8080" yaml:"Port"`

	DB struct {
		User     string `default:"root" yaml:"User"`
		Password string `yaml:"Password"`
		Protocol string `default:"tcp" yaml:"Protocol"`
		Host     string `default:"localhost" yaml:"Host"`
		Port     string `default:"3306" yaml:"Port"`
		Name     string `default:"tinyUrlMock_go" yaml:"Name"`
		Params   string `charset=utf8mb4,utf8&parseTime=True&timeout=5s&readTimeout=5s&writeTimeout=5s&sql_mode=''" yaml:"Params"`
	}

	Redis struct {
		Host string `default:"localhost" yaml:"Host"`
		Port string `default:"6379" yaml:"Port"`
	}
}{}

func Init() {
	err := configor.Load(&Config, "config.yaml")
	if err != nil {
		panic(err)
	}
}
