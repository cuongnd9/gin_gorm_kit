// Package v1 router v1
package v1

import (
	"github.com/103cuong/gorm_kit/controllers"
	"github.com/gin-gonic/gin"
)

// InitCatRouter initialize cat router
func InitCatRouter(r *gin.RouterGroup) *gin.RouterGroup {
	group := r.Group("/cats")
	{
		group.GET("/", controllers.GetCats)
		group.GET("/:id", controllers.GetCatByID)
		group.POST("/", controllers.CreateCat)
		group.PUT("/:id", controllers.UpdateCat)
		group.DELETE("/:id", controllers.DeleteCat)
	}
	return r
}
