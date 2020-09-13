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
	_, err := tx.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)
	if err != nil {
		return err
	}
	return nil
}

func downCreateUUIDExtension(tx *sql.Tx) error {
	_, err := tx.Exec(`DROP EXTENSION IF NOT EXISTS "uuid-ossp";`)
	if err != nil {
		return err
	}
	return nil
}
