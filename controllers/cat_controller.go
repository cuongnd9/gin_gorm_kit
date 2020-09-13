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
		ctx.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "no cat found!"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": cats})
}

// CreateCat create new cat.
func CreateCat(ctx *gin.Context) {
	var cat models.Cat
	err := ctx.BindJSON(&cat)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "invalid body"})
		return
	}
	err = services.CreateCat(&cat)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"status": http.StatusInternalServerError, "message": "create cat failed"},
		)
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "data": cat})
}

// GetCatByID get cat by id.
func GetCatByID(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	var cat models.Cat
	err := services.GetCatByID(&cat, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "no cat found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": cat})
}

// UpdateCat update cat.
func UpdateCat(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	var cat models.Cat
	err := services.GetCatByID(&cat, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "no cat found"})
		return
	}
	err = ctx.BindJSON(&cat)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "invalid body"})
		return
	}
	err = services.UpdateCat(&cat)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"status": http.StatusInternalServerError, "message": "update cat failed"},
		)
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
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"status": http.StatusInternalServerError, "message": "delete cat failed"},
		)
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{"status": http.StatusNoContent, "data": true})
}
