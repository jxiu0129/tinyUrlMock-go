package base

import (
	"fmt"
	"net/http"
	ucache "tinyUrlMock-go/api/cache/url"
	"tinyUrlMock-go/api/entities/edb"
	surl "tinyUrlMock-go/api/services/url"
	"tinyUrlMock-go/api/url"
	"tinyUrlMock-go/lib/db"
	"tinyUrlMock-go/lib/errors"
	"tinyUrlMock-go/lib/util"

	"github.com/gin-gonic/gin"
)

type (
	RedirectUrlRequest struct {
		Url string `uri:"redirect" binding:"required"`
	}
)

func (r *RedirectUrlRequest) validate(ctx *gin.Context) error {
	if err := ctx.ShouldBindUri(r); err != nil {
		return errors.ErrInvalidParams.SetError(err)
	}
	if len(r.Url) != 6 {
		errors.Throw(ctx, errors.ErrNoData.Err)
		return nil
	}
	r.Url = ctx.Param("redirect")
	return nil
}

func RedirectUrl(ctx *gin.Context) {
	req := &RedirectUrlRequest{}
	if err := req.validate(ctx); err != nil {
		errors.Throw(ctx, err)
		return
	}

	cacheOriginalUrl, err := ucache.GetOriginalUrl(req.Url)
	if err != nil {
		errors.Throw(ctx, err)
		return
	}

	if cacheOriginalUrl.OriginalUrl != "" {
		urlCheck := &edb.Url{OriginalUrl: cacheOriginalUrl.OriginalUrl}
		if util.IsUrlExpired(cacheOriginalUrl.CreatedAt) {
			if err := url.UrlExpired(urlCheck); err != nil {
				errors.Throw(ctx, errors.ErrNoData.Err)
				return
			}
		}
		fmt.Println("from redis")
		ctx.Redirect(http.StatusFound, "https://"+urlCheck.OriginalUrl)
		return
	}

	// 2. 再從db
	existUrl, err := surl.New(db.DBGorm).FindShortenUrl(req.Url)
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

	if err := ucache.SetOriginalUrl(existUrl); err != nil {
		errors.Throw(ctx, err)
		return
	}

	ctx.Redirect(http.StatusFound, "https://"+existUrl.OriginalUrl)

}
