package url

import (
	"net/http"
	"tinyUrlMock-go/lib/apires"
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
	// fmt.Printf("redirect url => %v", ctx.Param("redirectUrl"))
	// ctx.Redirect(302, "http://google.com")
	originalUrl := ctx.Request.PostFormValue("url")
	req := CreateTinyUrlRequest{}
	if err := req.Bind(ctx); err != nil {
		errors.Throw(ctx, err)
		return
	}
	// todo logic

	// ?ok ask: &CreateTinyUrlResponse => 吃值吃址都可以
	ctx.JSON(http.StatusOK, &CreateTinyUrlResponse{
		Base: apires.Base{
			Code:    errors.CODE_OK,
			Message: errors.MessageOK,
		},
		Data: CreateTinyUrlResponseData{
			// OrderID: hide.Uint64(ordersID),
			ShortenUrl:  "http://localhost:8080/cooool",
			OriginalUrl: originalUrl,
		},
	})

}
