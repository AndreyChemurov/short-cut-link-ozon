package service

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"ozon/internal/database"
)

const (
	alphabet       = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"
	linkLength     = 10
)

// Rules for short link:
//	- 10 letters long
//	- low/high register, digits, underscore only
//	- unique for each original link
func createShortLink(longLink string) (string, error) {
	rand.Seed(time.Now().UnixNano())

	unique := make([]byte, linkLength)

	for i := range unique {
		unique[i] = alphabet[rand.Intn(len(alphabet))]
	}

	return "http://ozon-test-task/" + string(unique), nil
}

func Create(ctx context.Context, longLink string) (string, error) {
	// Create unique link for the first time
	shortLink, err := createShortLink(longLink)
	if err != nil {
		return "", err
	}

	database, err := database.GetDatabase("postgres")
	if err != nil {
		return "", err
	}

	// Check if that long link was already used to create short one
	exists, err := database.CheckExistenceLongLink(ctx, longLink)
	if err != nil {
		return "", err
	}

	if exists {
		return "", fmt.Errorf("short link from link \"%s\" was already created", longLink)
	}

	// Check link uniqueness
	for {
		exists, err := database.CheckExistenceShortLink(ctx, shortLink)
		if err != nil {
			return "", err
		}

		if !exists {
			break
		}

		shortLink, err = createShortLink(longLink)
		if err != nil {
			return "", err
		}
	}

	if err := database.Create(ctx, longLink, shortLink); err != nil {
		return "", err
	}

	return shortLink, nil
}

func Get(ctx context.Context, shortLink string) (string, error) {
	database, err := database.GetDatabase("postgres")
	if err != nil {
		return "", err
	}

	longLink, err := database.Get(ctx, shortLink)
	if err != nil {
		return "", err
	}

	return longLink, nil
}
