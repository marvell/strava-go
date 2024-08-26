package inmemory

import (
	"context"
	"log/slog"
	"sync"

	"github.com/marvell/strava-go"
	"golang.org/x/oauth2"
)

type TokenStorage struct {
	m sync.Map
}

func (ts *TokenStorage) Get(_ context.Context, athleteID uint) (*oauth2.Token, error) {
	slog.Debug("get token", "athleteID", athleteID)

	v, ok := ts.m.Load(athleteID)
	if !ok {
		return nil, strava.ErrTokenNotFound
	}

	t, ok := v.(*oauth2.Token)
	if !ok {
		return nil, strava.ErrTokenNotFound
	}

	return t, nil
}

func (ts *TokenStorage) Save(_ context.Context, athleteID uint, token *oauth2.Token) error {
	slog.Debug("save token", "athleteID", athleteID)

	ts.m.Store(athleteID, token)
	return nil
}
