package main

import (
	"github.com/103cuong/gorm_kit/configs"
	"github.com/103cuong/gorm_kit/migrations"
	"github.com/103cuong/gorm_kit/routers"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	var err error
	configs.DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  configs.BuildDSN(),
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// migrate database.
	migrations.MigrateDB("up")

	r := routers.InitRouter()
	err = r.Run(":9000")
	if err != nil {
		panic("failed to run server")
	}
}
