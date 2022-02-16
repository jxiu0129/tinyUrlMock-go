package url

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func RedirectUrl(ctx *gin.Context) {
	fmt.Printf("redirect url => %v", ctx.Param("redirectUrl"))
	ctx.Redirect(302, "http://google.com")
}
