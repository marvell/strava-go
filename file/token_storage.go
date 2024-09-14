package file

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/marvell/strava-go"
	"golang.org/x/oauth2"
)

type TokenStorage struct {
	storageDir string
}

func NewTokenStorage(storageDir string) (*TokenStorage, error) {
	if err := os.MkdirAll(storageDir, 0755); err != nil {
		return nil, fmt.Errorf("create storage directory: %w", err)
	}
	return &TokenStorage{storageDir: storageDir}, nil
}

func (ts *TokenStorage) Get(_ context.Context, athleteID uint) (*oauth2.Token, error) {
	slog.Debug("get token", "athleteID", athleteID)

	data, err := os.ReadFile(ts.filename(athleteID))
	if err != nil {
		if os.IsNotExist(err) {
			return nil, strava.ErrTokenNotFound
		}
		return nil, fmt.Errorf("read token file: %w", err)
	}

	var token oauth2.Token
	if err := json.Unmarshal(data, &token); err != nil {
		return nil, fmt.Errorf("unmarshal token: %w", err)
	}

	return &token, nil
}

func (ts *TokenStorage) Save(_ context.Context, athleteID uint, token *oauth2.Token) error {
	slog.Debug("save token", "athleteID", athleteID)

	data, err := json.Marshal(token)
	if err != nil {
		return fmt.Errorf("marshal token: %w", err)
	}

	if err := os.WriteFile(ts.filename(athleteID), data, 0600); err != nil {
		return fmt.Errorf("write token file: %w", err)
	}

	return nil
}

func (ts *TokenStorage) filename(athleteID uint) string {
	return filepath.Join(ts.storageDir, fmt.Sprintf("%d.json", athleteID))
}
