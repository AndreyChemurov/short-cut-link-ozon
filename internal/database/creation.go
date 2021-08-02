package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

// createDatabase - SQL string for creating table.
const createDatabase string = `
BEGIN;

CREATE TABLE IF NOT EXISTS "links" (
	"id" BIGSERIAL NOT NULL PRIMARY KEY,
	"short_link" TEXT NOT NULL CHECK ("short_link" <> '') UNIQUE,
	"long_link" TEXT NOT NULL CHECK ("long_link" <> '') UNIQUE
);

CREATE INDEX IF NOT EXISTS links_short_idx ON links USING hash (short_link);
CREATE INDEX IF NOT EXISTS links_long_idx ON links USING hash (long_link);

COMMIT;
`

// CreateTableAndIndex - creates table and index for "short-linker".
func CreateTableAndIndex() error {
	var (
		host     = os.Getenv("POSTGRES_HOST")
		port     = os.Getenv("POSTGRES_PORT")
		user     = os.Getenv("POSTGRES_USER")
		password = os.Getenv("POSTGRES_PASSWORD")
		dbname   = os.Getenv("POSTGRES_DB")
	)

	log.Println("Creating database table if not exists...")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Println("Server failed on opening connection")
		return err
	}

	defer db.Close()

	_, err = db.Exec(createDatabase)

	if err != nil {
		log.Println("Server failed on creating table")
		return err
	}

	log.Println("Creating database table done")
	return nil
}
