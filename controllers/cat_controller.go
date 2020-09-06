// Package controllers gin's controller
package controllers

import (
	"net/http"

	"github.com/103cuong/gorm_kit/models"
	"github.com/103cuong/gorm_kit/services"
	"github.com/gin-gonic/gin"
)

// GetCats get all cats.
func GetCats(ctx *gin.Context) {
	var cats []models.Cat
	err := services.GetCats(&cats)
	if err != nil {
		// FIXME: response status + error.
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}
	ctx.JSON(http.StatusOK, cats)
}

// CreateCat create new cat.
func CreateCat(ctx *gin.Context) {
	var cat models.Cat
	err := ctx.BindJSON(&cat)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	err = services.CreateCat(&cat)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusCreated, cat)
}

// GetCatByID get cat by id.
func GetCatByID(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	var cat models.Cat
	err := services.GetCatByID(&cat, id)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, cat)
}

// UpdateCat update cat.
func UpdateCat(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	var cat models.Cat
	err := services.GetCatByID(&cat, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, cat)
		return
	}
	err = ctx.BindJSON(&cat)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	err = services.UpdateCat(&cat)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, cat)
}

// DeleteCat delete cat.
func DeleteCat(ctx *gin.Context) {
	var cat models.Cat
	id := ctx.Params.ByName("id")
	err := services.DeleteCat(&cat, id)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{"deleted": true})
}
