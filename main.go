package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	"tinyUrlMock-go/api/base"
	"tinyUrlMock-go/api/keys"
	"tinyUrlMock-go/api/url"
	"tinyUrlMock-go/config"
	"tinyUrlMock-go/lib/db"
	"tinyUrlMock-go/lib/redis"
	"tinyUrlMock-go/middleware"
)

const (
	defReadTimeout    = 10 * time.Second
	defWriteTimeout   = 30 * time.Second
	defMaxHeaderBytes = 1 << 20
)

func main() {
	config.Init() //=>會自動執行的init是小寫的
	router := gin.Default()
	// !error, and dont know what for
	// router.Use(ginrequestid.RequestId())

	// redis
	redis.Init()

	db.Init()

	// rate limiter
	router.Use(middleware.RateLimiterByIP(middleware.DefRateLimiterPeriod, config.Config.RateLimiter.Base))

	base.Route(router) //=> /:redirect(api)

	api := router.Group("/v1")
	keys.Route(api) //=> createNewKey(service)
	url.Route(api)  //=> createTinyUrl(api)

	/*
		*golang 的route一開始就會編譯好，不是這樣抓
		router.GET("*", func(ctx *gin.Context) {
			ctx.JSON(http.StatusNotFound, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
		})
	*/

	// router.Run(fmt.Sprintf(":%v", config.Config.Port))

	srv := &http.Server{
		Addr:           fmt.Sprintf(":%v", config.Config.Port),
		Handler:        router,
		ReadTimeout:    defReadTimeout,
		WriteTimeout:   defWriteTimeout,
		MaxHeaderBytes: defMaxHeaderBytes,
	}

	fmt.Printf("tinyUrlMock-go Server started listen port: %v\n", config.Config.Port)

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			// logs.Systemf("Server listen error: %v", err)
			fmt.Printf("sever listen error %v\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	sig := <-quit
	fmt.Printf("Shutdown Server with signal %v", sig)

	ctx, cancel := context.WithTimeout(context.Background(), defReadTimeout)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Printf("Server Shutdown err: %v\n", err)
	}

}
