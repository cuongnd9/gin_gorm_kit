package migrations

import (
	"database/sql"

	"github.com/103cuong/gorm_kit/configs"
	"github.com/103cuong/gorm_kit/models"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upCreateCategories, downCreateCategories)
}

func upCreateCategories(tx *sql.Tx) error {
	err := configs.DB.AutoMigrate(&models.Category{})
	if err != nil {
		return err
	}
	return nil
}

func downCreateCategories(tx *sql.Tx) error {
	err := configs.DB.Migrator().DropTable(&models.Category{})
	if err != nil {
		return err
	}
	return nil
}
