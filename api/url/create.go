package url

import (
	"fmt"
	"net/http"
	"time"
	"tinyUrlMock-go/api/entities/edb"
	"tinyUrlMock-go/api/keys"
	surl "tinyUrlMock-go/api/services/url"
	"tinyUrlMock-go/lib/apires"
	"tinyUrlMock-go/lib/db"
	"tinyUrlMock-go/lib/errors"

	"github.com/gin-gonic/gin"
)

type (
	CreateTinyUrlRequest struct {
		Url string `form:"url" binding:"required` //!why bad
	}
	// ? hide what for
	/* CreateGiftRequest struct {
		ProductID    hide.Uint64 `json:"product_id" binding:"required"`
	} */
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

// !fix -> createTinyUrl
func CreateTinyUrl(ctx *gin.Context) {
	originalUrl := ctx.Request.PostFormValue("url")
	req := CreateTinyUrlRequest{}
	var err error
	if err := req.Bind(ctx); err != nil {
		errors.Throw(ctx, err)
		return
	}
	// todo 0.先從redis找
	// todo url會過期
	// 1.找db有沒有一模一樣已經換過的，有的話直接回傳
	existUrl, err := surl.New(db.DBGorm).FindExistUrl(originalUrl)
	// fmt.Println("here ====> ", existUrl, err)
	if err != nil {
		errors.Throw(ctx, err)
		return

	}
	if existUrl != nil {
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

// todo
// func urlExpired(url string) error {

// }
