package base

import (
	"net/http"
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

	redirectUrl := ctx.Param("redirect")

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
