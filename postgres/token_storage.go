package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/marvell/strava-go"
	"golang.org/x/oauth2"
	"gorm.io/gorm"
)

func NewTokenStorage(db *gorm.DB) (*TokenStorage, error) {
	if err := db.AutoMigrate(&Token{}); err != nil {
		return nil, fmt.Errorf("could not auto migrate token: %w", err)
	}

	ts := &TokenStorage{
		db: db,
	}

	return ts, nil
}

type TokenStorage struct {
	db *gorm.DB
}

func (ts *TokenStorage) Get(ctx context.Context, athleteID uint) (*oauth2.Token, error) {
	var t Token
	if err := ts.db.WithContext(ctx).First(&t, athleteID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, strava.ErrTokenNotFound
		}

		return nil, fmt.Errorf("could not get token: %w", err)
	}

	return t.Token, nil
}

func (ts *TokenStorage) Save(ctx context.Context, athleteID uint, token *oauth2.Token) error {
	var t Token
	if err := ts.db.WithContext(ctx).First(&t, athleteID).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("could not get token: %w", err)
		}

		t = Token{
			Model: &gorm.Model{
				ID: athleteID,
			},
		}
	}
	t.Token = token

	if err := ts.db.Save(t).Error; err != nil {
		return fmt.Errorf("could not save token: %w", err)
	}

	return nil
}
