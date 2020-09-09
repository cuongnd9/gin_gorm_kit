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
	-- public.cats definition

	-- Drop table
	
	-- DROP TABLE public.cats;
	
	CREATE TABLE public.cats (
		id uuid NOT NULL DEFAULT uuid_generate_v4(),
		created_at timestamptz NULL,
		updated_at timestamptz NULL,
		deleted_at timestamptz NULL,
		"name" text NULL,
		color text NULL,
		category_id uuid NULL,
		CONSTRAINT cats_pkey PRIMARY KEY (id)
	);
	CREATE INDEX idx_cats_deleted_at ON public.cats USING btree (deleted_at);
	
	
	-- public.cats foreign keys
	
	ALTER TABLE public.cats ADD CONSTRAINT fk_categories_cats FOREIGN KEY (category_id) REFERENCES categories(id) ON UPDATE CASCADE ON DELETE SET NULL;
	ALTER TABLE public.cats ADD CONSTRAINT fk_cats_category FOREIGN KEY (category_id) REFERENCES categories(id) ON UPDATE CASCADE ON DELETE SET NULL;
	`)
	if err != nil {
		return err
	}
	return nil
}

func downCreateCats(tx *sql.Tx) error {
	_, err := tx.Exec(`
	DROP TABLE public.cats;
	`)
	if err != nil {
		return err
	}
	return nil
}
