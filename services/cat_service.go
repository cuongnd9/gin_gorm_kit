// Package services contain all business logic
package services

import (
	"github.com/103cuong/gorm_kit/configs"
	"github.com/103cuong/gorm_kit/models"
)

// GetCats fetch all cats.
func GetCats(cats *[]models.Cat) (err error) {
	err = configs.DB.Preload("Category").Find(&cats).Error
	if err != nil {
		return err
	}
	return nil
}

// CreateCat create new cat.
func CreateCat(cat *models.Cat) (err error) {
	err = configs.DB.Create(&cat).Error
	if err != nil {
		return err
	}
	return nil
}

// GetCatByID fetch cat by ID.
func GetCatByID(cat *models.Cat, id string) (err error) {
	err = configs.DB.Preload("Category").Where("id = ?", id).First(&cat).Error
	if err != nil {
		return err
	}
	return nil
}

// UpdateCat update cat.
func UpdateCat(cat *models.Cat) (err error) {
	err = configs.DB.Save(&cat).Error
	if err != nil {
		return err
	}
	return nil
}

// DeleteCat delete cat.
func DeleteCat(cat *models.Cat, id string) (err error) {
	err = configs.DB.Where("id = ?", id).Delete(&cat).Error
	if err != nil {
		return err
	}
	return nil
}
