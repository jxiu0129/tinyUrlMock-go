package keys

import (
	"fmt"
	"net/http"
	skey "tinyUrlMock-go/api/services/key"
	"tinyUrlMock-go/lib/apires"
	"tinyUrlMock-go/lib/db"
	"tinyUrlMock-go/lib/errors"

	"github.com/gin-gonic/gin"
)

// empty request example
/*
	func BeInvited(ctx *gin.Context) {
		uid := session.GetUID(ctx)
		count, err := smember.New(db.DBReaderGorm).GetBeInvitedCnt(uid.(uint64))
		if err != nil {
			errors.DBError(ctx, err)
			return
		}

		ctx.JSON(http.StatusOK, apires.Data{
			Base: apires.Base{
				Code:    errors.CODE_OK,
				Message: "ok",
			},
			Data: struct {
				InviteCnt uint `json:"invite_cnt"`
			}{
				InviteCnt: count,
			},
		})
}
*/

type (
	UpdateKeyRequest struct {
		// Url string `form:"url" binding:"required` //!why bad
		Amount int `json:"amount" `
	}
	// ? hide what for
	/* CreateGiftRequest struct {
		ProductID    hide.Uint64 `json:"product_id" binding:"required"`
	} */
	UpdateKeyResponse struct {
		apires.Base
		Data UpdateKeyResponseData `json:"data"`
	}

	UpdateKeyResponseData struct {
		UpdateKeys []string `json:"updateKeys"`
	}
)

// setKeysUsed (update) -> (service/control)
// setKeysUnused (update) -> (service/control)
// url_expired ? () -> url
func SetKeyUsed(ctx *gin.Context) {

	var err error
	// 1. find One From UnusedKey
	key, err := skey.New(db.DBGorm).FindOneUnusedKey()
	if err != nil {
		errors.Throw(ctx, err)
	}

	fmt.Println(key, err)
	updateKeys := []string{key}
	// 2. delete One From UnusedKey
	if err := skey.New(db.DBGorm).DeleteUnusedKeys(updateKeys); err != nil {
		errors.Throw(ctx, err)
	}
	// 3. insertUsedKey
	if err := skey.New(db.DBGorm).InsertUsedKeys(updateKeys); err != nil {
		errors.Throw(ctx, err)
	}
	ctx.JSON(http.StatusOK, &UpdateKeyResponse{
		Base: apires.Base{
			Code:    errors.CODE_OK,
			Message: errors.MessageOK,
		},
		Data: UpdateKeyResponseData{
			UpdateKeys: updateKeys,
		},
	})
}
