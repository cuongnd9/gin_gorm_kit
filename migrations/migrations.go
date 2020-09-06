// Package migrations migrations for Postgres
package migrations

import (
	"log"

	"github.com/103cuong/gorm_kit/configs"
	"github.com/pressly/goose"

	// postgres driver
	_ "github.com/lib/pq"
)

// MigrateDB migrate the database
func MigrateDB(command string) {
	dbstring := configs.BuildDSN()
	db, err := goose.OpenDBWithDriver("postgres", dbstring)
	if err != nil {
		log.Fatalf("goose: failed to open DB: %v\n", err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("goose: failed to close DB: %v\n", err)
		}
	}()

	if err := goose.Run(command, db, "./migrations"); err != nil {
		log.Fatalf("goose run failed :%v", err)
	}
}
