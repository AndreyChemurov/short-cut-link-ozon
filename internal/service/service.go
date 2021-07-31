package service

import (
	"context"

	"ozon/internal/database"
)

func Create(ctx context.Context, longLink string) (string, error) {
	database, err := database.GetDatabase("postgres")
	if err != nil {
		return "", err
	}

	shortLink, err := database.Create(ctx, longLink)
	if err != nil {
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
