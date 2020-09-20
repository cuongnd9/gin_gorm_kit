// Package migrations migrations for Postgres
package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upCreateUUIDExtension, downCreateUUIDExtension)
}

func upCreateUUIDExtension(tx *sql.Tx) error {
	_, err := tx.Exec(`create extension if not exists "uuid-ossp";`)
	if err != nil {
		return err
	}
	return nil
}

func downCreateUUIDExtension(tx *sql.Tx) error {
	_, err := tx.Exec(`drop extension if exists "uuid-ossp";`)
	if err != nil {
		return err
	}
	return nil
}
