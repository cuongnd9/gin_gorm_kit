package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upCreateCategories, downCreateCategories)
}

func upCreateCategories(tx *sql.Tx) error {
	_, err := tx.Exec(`
	create table if not exists categories
	(
		id uuid default uuid_generate_v4() not null
			constraint categories_pkey
				primary key,
		created_at timestamp with time zone,
		updated_at timestamp with time zone,
		deleted_at timestamp with time zone,
		name text
	);
	
	alter table categories owner to postgres;
	
	create index if not exists idx_categories_deleted_at
		on categories (deleted_at);
	`)
	if err != nil {
		return err
	}
	return nil
}

func downCreateCategories(tx *sql.Tx) error {
	_, err := tx.Exec(`
	drop table if exists categories;
	`)
	if err != nil {
		return err
	}
	return nil
}
