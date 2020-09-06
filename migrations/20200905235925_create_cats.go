// Package migrations migrations for Postgres
package migrations

import (
	"database/sql"

	"github.com/103cuong/gorm_kit/configs"
	"github.com/103cuong/gorm_kit/models"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upCreateCats, downCreateCats)
}

func upCreateCats(tx *sql.Tx) error {
	err := configs.DB.AutoMigrate(&models.Cat{})
	if err != nil {
		return err
	}
	return nil
}

func downCreateCats(tx *sql.Tx) error {
	err := configs.DB.Migrator().DropTable(&models.Cat{})
	if err != nil {
		return err
	}
	return nil
}
