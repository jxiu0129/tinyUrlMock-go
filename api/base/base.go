package base

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type response struct {
	status  int
	message string
	// data
}

type data struct {
	datas []string
}

func Route(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		fmt.Printf("%v\n", c)
		fmt.Println("hello world")
		// c.JSON(200, )
	})
	router.GET("/:redirect", func(c *gin.Context) {
		fmt.Printf("%v\n", c.Param("redirect"))
	})
	router.GET("/json", func(c *gin.Context) {
		fmt.Printf("%v\n", c.Request)
		var testRes response
		testRes.status = 200
		testRes.message = "test"
		// testRes.data = []string{}
		testRes2 := []string{"hello world", "world"}
		c.JSON(200, testRes2)
	})
}

// router.route("/admin").post(admin.createNewKeys).get(admin.clearAll); -> api/admin

// router.route("/createTinyUrl").post(url.createTinyUrl); -> api/url

// router.route("/:url").get(url.redirectUrl); -> api/base
