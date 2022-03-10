package keys

import (
	"github.com/gin-gonic/gin"
)

func Route(r *gin.RouterGroup) {
	g := r.Group("/keys")
	g.GET("/createNewKeys", CreateNewKeys)
}
