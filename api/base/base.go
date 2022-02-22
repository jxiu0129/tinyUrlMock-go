package base

import (
	"fmt"
	"net/http"
	surl "tinyUrlMock-go/api/services/url"
	"tinyUrlMock-go/lib/db"
	"tinyUrlMock-go/lib/errors"

	"github.com/gin-gonic/gin"
)

// for practice
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
	router.GET("/json", func(c *gin.Context) {
		fmt.Printf("%v\n", c.Request)
		var testRes response
		testRes.status = 200
		testRes.message = "test"
		// testRes.data = []string{}
		testRes2 := []string{"hello world", "world"}
		c.JSON(200, testRes2)
	})

	// ! real redirect
	router.GET("/:redirect", func(ctx *gin.Context) {
		redirectUrl := ctx.Param("redirect")
		if len(redirectUrl) == 6 {
			//todo 1. 先從redis
			// 2. 再從db
			url, err := surl.New(db.DBGorm).FindExistUrl(surl.FindUrl{ShortenUrl: redirectUrl})
			// fmt.Println(url)
			if err != nil {
				errors.Throw(ctx, err)
				return
			}
			if url != nil {
				fmt.Println("redirect from db")
				ctx.Redirect(http.StatusFound, "https://"+url.OriginalUrl)
			}
		}
	})
}
