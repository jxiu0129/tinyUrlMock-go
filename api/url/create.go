package url

import (
	"fmt"
	"net/http"
	"time"
	ucache "tinyUrlMock-go/api/cache/url"
	"tinyUrlMock-go/api/entities/edb"
	"tinyUrlMock-go/api/keys"
	surl "tinyUrlMock-go/api/services/url"
	"tinyUrlMock-go/lib/apires"
	"tinyUrlMock-go/lib/db"
	"tinyUrlMock-go/lib/errors"
	"tinyUrlMock-go/lib/util"

	"github.com/gin-gonic/gin"
)

type (
	CreateTinyUrlRequest struct {
		Url string `form:"url" binding:"required"`
	}
	CreateTinyUrlResponse struct {
		apires.Base
		Data CreateTinyUrlResponseData `json:"data"`
	}

	CreateTinyUrlResponseData struct {
		OriginalUrl string `json:"originalUrl"`
		ShortenUrl  string `json:"shortenUrl"`
	}
)

func (req *CreateTinyUrlRequest) Bind(ctx *gin.Context) error {
	if err := ctx.ShouldBind(req); err != nil {
		return err
	}
	return nil
}

func CreateTinyUrl(ctx *gin.Context) {
	originalUrl := ctx.Request.PostFormValue("url")
	req := CreateTinyUrlRequest{}
	var err error
	if err := req.Bind(ctx); err != nil {
		errors.Throw(ctx, err)
		return
	}
	// 0.先從redis找
	cacheShortenUrl, err := ucache.GetUniqueKey(originalUrl)
	if err != nil {
		errors.Throw(ctx, err)
		return
	}
	if cacheShortenUrl.ShortenUrl != "" {
		urlCheck := &edb.Url{ShortenUrl: cacheShortenUrl.ShortenUrl}
		if util.IsUrlExpired(cacheShortenUrl.CreatedAt) {
			if err := UrlExpired(urlCheck); err != nil {
				errors.Throw(ctx, errors.ErrNoData.Err)
				return
			}
		}
		ctx.JSON(http.StatusOK, &CreateTinyUrlResponse{
			Base: apires.Base{
				Code:    errors.CODE_OK,
				Message: errors.MessageOK,
			},
			Data: CreateTinyUrlResponseData{
				ShortenUrl:  "http://localhost:8080/" + cacheShortenUrl.ShortenUrl,
				OriginalUrl: originalUrl,
			},
		})
		fmt.Println("from redis")
		return
	}
	existUrl, err := surl.New(db.DBGorm).FindOriginalUrl(originalUrl)
	if err != nil {
		errors.Throw(ctx, err)
		return

	}
	if existUrl != nil {
		if util.IsUrlExpired(existUrl.CreatedAt) {
			if err := UrlExpired(existUrl); err != nil {
				errors.Throw(ctx, errors.ErrNoData.Err)
				return
			}
		}
		setCacheData := &edb.Url{
			ShortenUrl:  existUrl.ShortenUrl,
			OriginalUrl: originalUrl,
			CreatedAt:   existUrl.CreatedAt,
		}
		if err := ucache.SetUniqueKey(setCacheData); err != nil {
			errors.Throw(ctx, err)
			return
		}

		ctx.JSON(http.StatusOK, &CreateTinyUrlResponse{
			Base: apires.Base{
				Code:    errors.CODE_OK,
				Message: errors.MessageOK,
			},
			Data: CreateTinyUrlResponseData{
				ShortenUrl:  "http://localhost:8080/" + existUrl.ShortenUrl,
				OriginalUrl: originalUrl,
			},
		})
		fmt.Println("from db")
		return

	}

	newkey, err := keys.SetOneKeyUsed()
	if err != nil {
		errors.Throw(ctx, err)
		return
	}
	newUrl := []*edb.Url{
		{ShortenUrl: newkey, OriginalUrl: originalUrl, CreatedAt: time.Now()},
	}
	if err := surl.New(db.DBGorm).InsertUrls(newUrl); err != nil {
		errors.Throw(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, &CreateTinyUrlResponse{
		Base: apires.Base{
			Code:    errors.CODE_OK,
			Message: errors.MessageOK,
		},
		Data: CreateTinyUrlResponseData{
			ShortenUrl:  "http://localhost:8080/" + newkey,
			OriginalUrl: originalUrl,
		},
	})

}
