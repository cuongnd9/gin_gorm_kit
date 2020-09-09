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
	-- public.categories definition

	-- Drop table
	
	-- DROP TABLE public.categories;
	
	CREATE TABLE public.categories (
		id uuid NOT NULL DEFAULT uuid_generate_v4(),
		created_at timestamptz NULL,
		updated_at timestamptz NULL,
		deleted_at timestamptz NULL,
		name text NULL,
		CONSTRAINT categories_pkey PRIMARY KEY (id)
	);
	CREATE INDEX idx_categories_deleted_at ON public.categories USING btree (deleted_at);
	`)
	if err != nil {
		return err
	}
	return nil
}

func downCreateCategories(tx *sql.Tx) error {
	_, err := tx.Exec(`
	DROP TABLE public.categories;
	`)
	if err != nil {
		return err
	}
	return nil
}
