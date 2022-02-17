package admin

import (
	"github.com/gin-gonic/gin"
)

// !fix -> createKey
func CreateTinyUrl(ctx *gin.Context) {
	ctx.JSON(200, "OK")
	// ctx.JSON(http.StatusOK, apires.Base{
	// 	Code:    errors.CODE_OK,
	// 	Message: errors.MessageOK,
	// })
}
