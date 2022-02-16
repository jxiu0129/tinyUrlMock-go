package admin

import (
	"github.com/gin-gonic/gin"
)

func Route(r *gin.RouterGroup) {
	g := r.Group("/url")
	g.POST("/createTinyUrl", CreateTinyUrl)
}
