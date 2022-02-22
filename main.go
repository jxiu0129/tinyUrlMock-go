package main

import (
	"github.com/gin-gonic/gin"

	"tinyUrlMock-go/api/base"
	"tinyUrlMock-go/api/keys"
	"tinyUrlMock-go/api/url"
	"tinyUrlMock-go/config"
	"tinyUrlMock-go/lib/db"
)

func main() {
	config.Init() //=>會自動執行的init是小寫的
	router := gin.Default()
	// !error, and dont know what for
	// router.Use(ginrequestid.RequestId())

	db.Init()

	base.Route(router) //=> /:redirect(api)

	api := router.Group("/v1")
	keys.Route(api) //=> createNewKey(service)
	url.Route(api)  //=> createTinyUrl(api)

	router.Run(":8080")

	/* from funnow-go
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.Config.Port),
		Handler: router,
		// ReadTimeout:    defReadTimeout,
		// WriteTimeout:   defWriteTimeout,
		// MaxHeaderBytes: defMaxHeaderBytes,
	}

	fmt.Printf("FunNow API Server started listen port: %d", config.Config.Port)

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			// logs.Systemf("Server listen error: %v", err)
			fmt.Printf("sever listen error %v", err)
		}
	}()
	*/

}
