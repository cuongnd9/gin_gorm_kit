// Package routers gin's routers
package routers

import (
	v1 "github.com/103cuong/gorm_kit/routers/v1"
	"github.com/gin-gonic/gin"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	apiV1 := r.Group("/api/v1")
	v1.InitCatRouter(apiV1)
	v1.InitCategoryRouter(apiV1)
	return r
}
