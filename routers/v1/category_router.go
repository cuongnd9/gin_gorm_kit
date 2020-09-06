// Package v1 router v1
package v1

import (
	"github.com/103cuong/gorm_kit/controllers"
	"github.com/gin-gonic/gin"
)

// InitCategoryRouter initialize category router
func InitCategoryRouter(r *gin.RouterGroup) *gin.RouterGroup {
	group := r.Group("/categories")
	{
		group.GET("/", controllers.GetCategories)
		group.POST("/", controllers.CreateCategory)
	}
	return r
}
