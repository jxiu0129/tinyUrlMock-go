package config

import (
	"github.com/jinzhu/configor"
)

// *會先抓yaml才抓這裡的default

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
		Debug    bool   `default:"false" yaml:"Debug"`
	}

	Redis struct {
		MaxIdleConns int    `default:"20" yaml:"MaxIdleConns"`
		Protocol     string `default:"tcp" yaml:"Protocol"`
		Host         string `default:"redis" yaml:"Host"`
		Port         string `default:"6379" yaml:"Port"`
	}

	RateLimiter struct {
		Base           int64 `default:"10" yaml:"AdminAuth"`
		AdminAuth      int64 `env:"RATE_LIMITER_ADMIN_AUTH" default:"5" yaml:"AdminAuth"`
		BossNowAuth    int64 `env:"RATE_LIMITER_BOSSNOW_AUTH" default:"5" yaml:"BossNowAuth"`
		FunNowAuth     int64 `env:"RATE_LIMITER_FUNNOW_AUTH" default:"5" yaml:"FunNowAuth"`
		FunBookReserve int64 `env:"RATE_LIMITER_FUNBOOK_RESERVE" default:"5" yaml:"FunBookReserve"`
	}
}{}

func Init() {
	err := configor.Load(&Config, "config.yaml")
	if err != nil {
		panic(err)
	}
}
