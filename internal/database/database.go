package database

import (
	"context"
	"fmt"
)

type Database interface {
	Create(ctx context.Context, longLink string) (shortLink string, err error)
	Get(ctx context.Context, shortLink string) (longLink string, err error)
}

var (
	databases = map[string]Database{
		"postgres": newPostgres(),
	}
)

// GetDatabase - фабричный метод (здесь только для postgres)
// получения объекта базы данных.
func GetDatabase(driver string) (Database, error) {
	database, found := databases[driver]
	if !found {
		return nil, fmt.Errorf("no database with driver %s", driver)
	}

	return database, nil
}
