// Package migrations migrations for Postgres
package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upCreateCats, downCreateCats)
}

func upCreateCats(tx *sql.Tx) error {
	_, err := tx.Exec(`
	create table if not exists cats
	(
		id uuid default uuid_generate_v4() not null
			constraint cats_pkey
				primary key,
		created_at timestamp with time zone,
		updated_at timestamp with time zone,
		deleted_at timestamp with time zone,
		name text,
		color text,
		category_id uuid
			constraint fk_categories_cats
				references categories
					on update cascade on delete set null
			constraint fk_cats_category
				references categories
					on update cascade on delete set null
	);
	
	alter table cats owner to postgres;
	
	create index if not exists idx_cats_deleted_at
		on cats (deleted_at);
	`)
	if err != nil {
		return err
	}
	return nil
}

func downCreateCats(tx *sql.Tx) error {
	_, err := tx.Exec(`
	drop table if exists cats;
	`)
	if err != nil {
		return err
	}
	return nil
}
