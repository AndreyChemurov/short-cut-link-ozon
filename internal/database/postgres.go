package database

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq" //
)

// postgres - реализация интерфейса Database для postgres.
type postgres struct {
	db *sql.DB
}

var instance *postgres

func newPostgres() Database {
	if instance != nil {
		return instance
	}

	var (
		host     = os.Getenv("POSTGRES_HOST")
		port     = os.Getenv("POSTGRES_PORT")
		user     = os.Getenv("POSTGRES_USER")
		password = os.Getenv("POSTGRES_PASSWORD")
		dbname   = os.Getenv("POSTGRES_DB")
	)

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil
	}

	err = db.Ping()
	if err != nil {
		return nil
	}

	instance = &postgres{
		db: db,
	}

	return instance
}

func (p *postgres) Create(ctx context.Context, longLink, shortLink string) (err error) {
	// Begin transaction
	tx, err := p.db.BeginTx(ctx, nil)
	if err != nil {
		return
	}

	// Defer will be ignored on commit
	defer func() {
		if err != nil {
			if rerr := tx.Rollback(); rerr != nil {
				err = fmt.Errorf("error on create: %v. error on rollback: %v", err, rerr)
			}
		}
	}()

	// Prepare to insert into "links" table
	stmt, err := tx.PrepareContext(ctx, "INSERT INTO links VALUES (DEFAULT, $1, $2);")
	if err != nil {
		return
	}

	defer func() {
		if closeErr := stmt.Close(); closeErr != nil {
			err = fmt.Errorf("error on create: %v. error on stmt close: %v", err, closeErr)
		}
	}()

	// Add short and long links to table
	if _, err = stmt.ExecContext(ctx, shortLink, longLink); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		return
	}

	return
}

func (p *postgres) Get(ctx context.Context, shortLink string) (longLink string, err error) {
	// Begin transaction
	tx, err := p.db.BeginTx(ctx, nil)
	if err != nil {
		return
	}

	// Defer will be ignored on commit
	defer func() {
		if err != nil {
			if rerr := tx.Rollback(); rerr != nil {
				err = fmt.Errorf("error on create: %v. error on rollback: %v", err, rerr)
			}
		}
	}()

	// Prepare to select
	stmt, err := tx.Prepare("SELECT long_link FROM links WHERE short_link = $1;")
	if err != nil {
		return
	}

	if err = stmt.QueryRowContext(ctx, shortLink).Scan(&longLink); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		return
	}

	return
}

func (p *postgres) CheckExistenceShortLink(ctx context.Context, shortLink string) (exists bool, err error) {
	// Begin transaction
	tx, err := p.db.BeginTx(ctx, nil)
	if err != nil {
		return
	}

	// Defer will be ignored on commit
	defer func() {
		if err != nil {
			if rerr := tx.Rollback(); rerr != nil {
				err = fmt.Errorf("error on create: %v. error on rollback: %v", err, rerr)
			}
		}
	}()

	// Prepare to select
	stmt, err := tx.Prepare("SELECT short_link FROM links WHERE long_link = $1;")
	if err != nil {
		return
	}

	if err = stmt.QueryRowContext(ctx, shortLink).Err(); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		return
	}

	return
}

func (p *postgres) CheckExistenceLongLink(ctx context.Context, longLink string) (exists bool, err error) {
	// Begin transaction
	tx, err := p.db.BeginTx(ctx, nil)
	if err != nil {
		return
	}

	// Defer will be ignored on commit
	defer func() {
		if err != nil {
			if rerr := tx.Rollback(); rerr != nil {
				err = fmt.Errorf("error on create: %v. error on rollback: %v", err, rerr)
			}
		}
	}()

	// Prepare to select
	stmt, err := tx.Prepare("SELECT EXISTS(SELECT 1 FROM links WHERE long_link = $1 LIMIT 1);")
	if err != nil {
		return
	}

	if err = stmt.QueryRowContext(ctx, longLink).Scan(&exists); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		return
	}

	return
}
