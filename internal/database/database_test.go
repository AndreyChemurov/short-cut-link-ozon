package database_test

import (
	"context"
	"log"
	"ozon/internal/database"
	"strconv"
	"testing"
	"time"
)

func TestDatabaseCreateFail(t *testing.T) {
	_, err := database.GetDatabase("error-database")
	if err == nil {
		t.Errorf("database creation return wrong response. want: %w, got: nil", err)
	}

	log.Printf("Response: %v", err)
}

func TestDatabaseOpenConnectionFail(t *testing.T) {
	_, err := database.GetDatabase("postgres")
	if err != nil {
		t.Errorf("database open connection return wrong response. want: nil, got: : %w", err)
	}

	log.Printf("Response: %v", err)
}

func TestDatabaseCreate(t *testing.T) {
	db, err := database.GetDatabase("postgres")
	if err != nil {
		t.Errorf("database Create return wrong response. want: nil, got: : %w", err)
	}

	longLink := strconv.Itoa(int(time.Now().UnixNano()))
	shortLink := strconv.Itoa(int(time.Now().UnixNano()))

	if err := db.Create(context.Background(), longLink, shortLink); err != nil {
		t.Errorf("database Create return wrong response. want: nil, got: : %w", err)
	}

	log.Printf("Response: %v", err)
}

func TestDatabaseCheckExistenceLongLink(t *testing.T) {
	db, err := database.GetDatabase("postgres")
	if err != nil {
		t.Errorf("database CheckExistenceLongLink return wrong response. want: nil, got: : %w", err)
	}

	if _, err := db.CheckExistenceLongLink(context.Background(), "long_test_do_not_exist"); err != nil {
		t.Errorf("database CheckExistenceLongLink return wrong response. want: nil, got: : %w", err)
	}

	log.Printf("Response: %v", err)
}

func TestDatabaseCheckExistenceShortLink(t *testing.T) {
	db, err := database.GetDatabase("postgres")
	if err != nil {
		t.Errorf("database CheckExistenceShortLink return wrong response. want: nil, got: : %w", err)
	}

	if _, err := db.CheckExistenceShortLink(context.Background(), "short_test_do_not_exist"); err != nil {
		t.Errorf("database CheckExistenceShortLink return wrong response. want: nil, got: : %w", err)
	}

	log.Printf("Response: %v", err)
}
