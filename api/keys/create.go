package keys

import (
	"net/http"
	"regexp"
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
		Amount int `form:"amount" `
	}
	CreateNewKeysResponse struct {
		apires.Base
		Data CreateNewKeysResponseData `json:"data"`
	}

	CreateNewKeysResponseData struct {
		NewKeysCount int      `json:"newKeysCount"`
		NewKeys      []string `json:"newKeys"`
	}
)

// todo not finish yet
func (r *CreateNewKeysRequest) validate(ctx *gin.Context) error {
	if err := ctx.ShouldBindQuery(r); err != nil {
		return errors.ErrInvalidParams.SetError(err)
	}
	match, err := regexp.MatchString(`^[1-9]\d*$`, strconv.Itoa(r.Amount))
	if err != nil {
		return err
	}
	if !match {
		return errors.ErrInvalidParams.SetError(err)
	}
	return nil
}

func CreateNewKeys(ctx *gin.Context) {
	req := &CreateNewKeysRequest{}
	if err := req.validate(ctx); err != nil {
		errors.Throw(ctx, err)
		return
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
	for len(newKeys) < req.Amount {
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
