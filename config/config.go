package config

import (
	"github.com/jinzhu/configor"
)

var Config = struct {
	// 坑：tag裡key沒寫清楚會直接不compile然後也不說
	BaseURL string `default:"http://localhost" yaml:"BaseURL"`
	Port    string `default:"8080" yaml:"Port"`

	DB struct {
		User     string `default:"root" yaml:"User"`
		Password string `yaml:"Password"`
		Protocol string `default:"tcp" yaml:"Protocol"`
		Host     string `default:"localhost" yaml:"Host"`
		Port     string `default:"3306" yaml:"Port"`
		Name     string `default:"tinyUrlMock_go" yaml:"Name"`
		Params   string `yaml:"Params"`
	}

	Redis struct {
		MaxIdleConns int    `default:"20" yaml:"MaxIdleConns"`
		Protocol     string `default:"tcp" yaml:"Protocol"`
		Host         string `default:"redis" yaml:"Host"`
		Port         string `default:"6379" yaml:"Port"`
	}
}{}

func Init() {
	err := configor.Load(&Config, "config.yaml")
	if err != nil {
		panic(err)
	}
}
