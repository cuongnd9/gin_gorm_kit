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
		ctx.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "no category found!"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": categories})
}

// CreateCategory create new category.
func CreateCategory(ctx *gin.Context) {
	var category models.Category
	err := ctx.BindJSON(&category)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "invalid body"})
		return
	}
	err = services.CreateCategory(&category)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"status": http.StatusInternalServerError, "message": "create category failed"},
		)
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "data": category})
}
