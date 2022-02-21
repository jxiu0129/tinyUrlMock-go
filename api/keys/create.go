package keys

import (
	"net/http"
	"strconv"
	skey "tinyUrlMock-go/api/services/key"
	"tinyUrlMock-go/lib/apires"
	"tinyUrlMock-go/lib/db"
	"tinyUrlMock-go/lib/errors"

	"github.com/gin-gonic/gin"
	gonanoid "github.com/matoous/go-nanoid/v2"
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
	CreateNewKeysRequest struct {
		// Url string `form:"url" binding:"required` //!why bad
		Amount int `json:"amount" `
	}
	// ? hide what for
	/* CreateGiftRequest struct {
		ProductID    hide.Uint64 `json:"product_id" binding:"required"`
	} */
	CreateNewKeysResponse struct {
		apires.Base
		Data CreateNewKeysResponseData `json:"data"`
	}

	CreateNewKeysResponseData struct {
		NewKeysCount int      `json:"newKeysCount"`
		NewKeys      []string `json:"newKeys"`
	}
)

// createNewKeys (create) -> (service/create)
// setKeysUsed (update) -> (service/control)
// setKeysUnused (update) -> (service/control)
// url_expired ? () -> url
func CreateNewKeys(ctx *gin.Context) {
	// todo 要套用request struct
	// todo 要加transaction
	amount, err := strconv.Atoi(ctx.Query("amount"))
	if err != nil {
		errors.Throw(ctx, err)
	}

	currentKeys := make(map[string]bool)
	newKeys := []string{}

	// 1. check keys from current Used/Unused db
	allUsedKeys, err := skey.New(db.DBGorm).SearchAllUsedKeys()
	if err != nil {
		errors.Throw(ctx, err)
		return
	}
	for _, entry := range allUsedKeys {
		currentKeys[entry.UniqueKey] = true
	}

	allUnusedKeys, err := skey.New(db.DBGorm).SearchAllUnusedKeys()
	if err != nil {
		errors.Throw(ctx, err)
		return
	}
	for _, entry := range allUnusedKeys {
		currentKeys[entry.UniqueKey] = true
	}

	// 2. create new keys
	for len(newKeys) < amount {
		id, err := gonanoid.Generate("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz", 6)
		if err != nil {
			errors.Throw(ctx, err)
			return
		}

		if _, check := currentKeys[id]; check {
			continue
		}
		newKeys = append(newKeys, id)
		currentKeys[id] = true

	}
	// 3. insert

	res, err := skey.New(db.DBGorm).InsertNewUnusedKeys(newKeys)
	if err != nil {
		errors.Throw(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, &CreateNewKeysResponse{
		Base: apires.Base{
			Code: errors.CODE_OK,
			// Message: errors.MessageOK,
			Message: res,
		},
		Data: CreateNewKeysResponseData{
			NewKeysCount: len(newKeys),
			NewKeys:      newKeys,
		},
	})
}
