package keys

import (
	"fmt"
	skey "tinyUrlMock-go/api/services/key"
	"tinyUrlMock-go/lib/apires"
	"tinyUrlMock-go/lib/db"
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
func SetOneKeyUsed() (string, error) {

	var err error
	// 1. find One From UnusedKey
	key, err := skey.New(db.DBGorm).FindOneUnusedKey()
	if err != nil {
		return "", err
	}

	fmt.Println(key, err)
	updateKeys := []string{key}
	// 2. delete One From UnusedKey
	if err := skey.New(db.DBGorm).DeleteUnusedKeys(updateKeys); err != nil {
		return "", err
	}
	// 3. insertUsedKey
	if err := skey.New(db.DBGorm).InsertUsedKeys(updateKeys); err != nil {
		return "", err
	}
	// ctx.JSON(http.StatusOK, &UpdateKeyResponse{
	// 	Base: apires.Base{
	// 		Code:    errors.CODE_OK,
	// 		Message: errors.MessageOK,
	// 	},
	// 	Data: UpdateKeyResponseData{
	// 		UpdateKeys: updateKeys,
	// 	},
	// })
	return key, nil
}

// todo 沒測過
func SetKeyUnused(keys []string) error {

	// keys => ['ws231w', 'dqwdw2',...]
	// var err error
	if err := skey.New(db.DBGorm).DeleteUsedKeys(keys); err != nil {
		return err
	}
	if err := skey.New(db.DBGorm).InsertUnusedKeys(keys); err != nil {
		return err
	}
	return nil
}
