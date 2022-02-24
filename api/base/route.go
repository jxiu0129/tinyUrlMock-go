package base

import (
	"github.com/gin-gonic/gin"
)

func Route(r *gin.RouterGroup) {
	g := r.Group("/")
	g.GET("/:redirect", RedirectUrl)
}
