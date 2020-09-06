// Package controllers gin's controller
package controllers

import (
	"net/http"

	"github.com/103cuong/gorm_kit/models"
	"github.com/103cuong/gorm_kit/services"
	"github.com/gin-gonic/gin"
)

// GetCategories get all categories
func GetCategories(ctx *gin.Context) {
	var categories []models.Category
	err := services.GetCategories(&categories)
	if err != nil {
		// FIXME: response status + error.
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}
	ctx.JSON(http.StatusOK, categories)
}

// CreateCategory create new category.
func CreateCategory(ctx *gin.Context) {
	var category models.Category
	err := ctx.BindJSON(&category)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	err = services.CreateCategory(&category)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusCreated, category)
}
