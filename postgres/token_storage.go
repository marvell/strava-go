package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/marvell/strava-go"
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

var _ strava.TokenStorage = (*TokenStorage)(nil)

func (ts *TokenStorage) Get(ctx context.Context, athleteID uint) (*strava.Token, error) {
	var t Token
	if err := ts.db.WithContext(ctx).First(&t, athleteID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, strava.ErrTokenNotFound
		}

		return nil, fmt.Errorf("could not get token: %w", err)
	}

	return t.Token, nil
}

func (ts *TokenStorage) Save(ctx context.Context, token *strava.Token) error {
	var t Token
	if err := ts.db.WithContext(ctx).First(&t, token.AthleteID).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("could not get token: %w", err)
		}

		t = Token{
			Model: &gorm.Model{
				ID: token.AthleteID,
			},
		}
	}
	t.Token = token

	if err := ts.db.Save(t).Error; err != nil {
		return fmt.Errorf("could not save token: %w", err)
	}

	return nil
}
