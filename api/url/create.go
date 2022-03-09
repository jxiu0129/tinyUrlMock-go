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
		Url string `form:"url" binding:"required"` //!why bad
	}
	CreateTinyUrlResponse struct {
		apires.Base
		Data CreateTinyUrlResponseData `json:"data"`
	}

	CreateTinyUrlResponseData struct {
		OriginalUrl string `json:"originalUrl"`
		ShortenUrl  string `json:"shortenUrl"` //todo: 會在回傳時欄位變tag裡的，庫，待研究
	}
)

func (req *CreateTinyUrlRequest) Bind(ctx *gin.Context) error {
	if err := ctx.ShouldBind(req); err != nil {
		return err
	}
	// put things in req

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
	// 1.找db有沒有一模一樣已經換過的，有的話直接回傳
	existUrl, err := surl.New(db.DBGorm).FindOriginalUrl(originalUrl)
	if err != nil {
		errors.Throw(ctx, err)
		return

	}
	if existUrl != nil {
		if util.IsUrlExpired(existUrl.CreatedAt) {
			// url expired
			if err := UrlExpired(existUrl); err != nil {
				errors.Throw(ctx, errors.ErrNoData.Err)
				return
			}
		}
		// redis set
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

	// 2.新增一個tinyurl, (key DB)unused => used
	// 2.1 setKeysUsed();
	newkey, err := keys.SetOneKeyUsed()
	if err != nil {
		errors.Throw(ctx, err)
		return
	}
	// 2.2 insert_newURL(url, uniqueKey);
	newUrl := []*edb.Url{
		{ShortenUrl: newkey, OriginalUrl: originalUrl, CreatedAt: time.Now()},
	}
	if err := surl.New(db.DBGorm).InsertUrls(newUrl); err != nil {
		errors.Throw(ctx, err)
		return
	}

	// ?ok ask: &CreateTinyUrlResponse => 吃值吃址都可以
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
