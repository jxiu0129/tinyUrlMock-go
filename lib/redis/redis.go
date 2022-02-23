package redis

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
	"tinyUrlMock-go/config"

	"github.com/gin-contrib/cache/persistence"
	goredis "github.com/go-redis/redis/v7"
	"github.com/gomodule/redigo/redis"
)

var (
	Pool       *redis.Pool
	CacheStore *persistence.RedisStore
	Client     *goredis.Client
)

func Init() {

	redisConfig := config.Config.Redis
	addr := fmt.Sprintf("%v:%v", redisConfig.Host, redisConfig.Port)
	Pool = newPool(addr)
	CacheStore = persistence.NewRedisCacheWithPool(Pool, time.Nanosecond)
	cleanupHook()

	Client = goredis.NewClient(&goredis.Options{
		Addr: addr,
	})
	fmt.Printf("%v at %v:%v\n", Client.Ping(), redisConfig.Host, redisConfig.Port)
}

func newPool(addr string) *redis.Pool {

	return &redis.Pool{

		MaxIdle:     config.Config.Redis.MaxIdleConns,
		IdleTimeout: 240 * time.Second,

		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", addr)
			if err != nil {
				return nil, err
			}
			return c, err
		},

		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func cleanupHook() {

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-c
		Pool.Close()
	}()
}
