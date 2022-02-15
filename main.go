package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello world")
	})

	router.Run(":8080")

	// srv := &http.Server{
	// 	Addr:           fmt.Sprintf(":%d", config.Config.Port),
	// 	Handler:        router,
	// 	// ReadTimeout:    defReadTimeout,
	// 	// WriteTimeout:   defWriteTimeout,
	// 	// MaxHeaderBytes: defMaxHeaderBytes,
	// }

	// fmt.Printf("FunNow API Server started listen port: %d", config.Config.Port)

	// go func() {
	// 	// service connections
	// 	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
	// 		// logs.Systemf("Server listen error: %v", err)
	// 		fmt.Printf("sever listen error %v", err)
	// 	}
	// }()
}
