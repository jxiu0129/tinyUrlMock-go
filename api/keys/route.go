package keys

import (
	"github.com/gin-gonic/gin"
)

// createKey
func Route(r *gin.RouterGroup) {
	g := r.Group("/keys")
	g.GET("/createNewKeys", CreateNewKeys)
	g.GET("/setKeyUsed", SetKeyUsed)
}
