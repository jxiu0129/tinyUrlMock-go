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

type (
	CreateNewKeysRequest struct {
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

	res, err := skey.New(db.DBGorm).InsertNewUnusedKeys(newKeys)
	if err != nil {
		errors.Throw(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, &CreateNewKeysResponse{
		Base: apires.Base{
			Code:    errors.CODE_OK,
			Message: res,
		},
		Data: CreateNewKeysResponseData{
			NewKeysCount: len(newKeys),
			NewKeys:      newKeys,
		},
	})
}
