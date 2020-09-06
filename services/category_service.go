// Package services contain all business logic
package services

import (
	"github.com/103cuong/gorm_kit/configs"
	"github.com/103cuong/gorm_kit/models"
)

// GetCategories fetch all categories
func GetCategories(categories *[]models.Category) (err error) {
	err = configs.DB.Preload("Cats").Find(&categories).Error
	if err != nil {
		return err
	}
	return nil
}

// CreateCategory create new category
func CreateCategory(category *models.Category) (err error) {
	err = configs.DB.Create(&category).Error
	if err != nil {
		return err
	}
	return nil
}
