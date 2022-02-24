package base

import (
	"fmt"
	"net/http"
	surl "tinyUrlMock-go/api/services/url"
	"tinyUrlMock-go/api/url"
	"tinyUrlMock-go/lib/db"
	"tinyUrlMock-go/lib/errors"
	"tinyUrlMock-go/lib/util"

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

func RedirectUrl(ctx *gin.Context) {
	redirectUrl := ctx.Param("redirect")
	fmt.Println(redirectUrl)

	// 放進request
	if len(redirectUrl) != 6 {
		errors.Throw(ctx, errors.ErrNoData.Err)
		return
	}

	//todo 1. 先從redis

	// 2. 再從db
	existUrl, err := surl.New(db.DBGorm).FindShortenUrl(redirectUrl)
	if err != nil {
		errors.Throw(ctx, err)
		return
	}
	if util.IsUrlExpired(existUrl.CreatedAt) {
		// url expired
		if err := url.UrlExpired(existUrl); err != nil {
			errors.Throw(ctx, err)
			return
		}
		errors.Throw(ctx, errors.ErrNoData.Err)
		return
	}

	ctx.Redirect(http.StatusFound, "https://"+existUrl.OriginalUrl)

}
