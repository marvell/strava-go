package inmemory

import (
	"context"
	"log/slog"
	"sync"

	"github.com/marvell/strava-go"
)

type TokenStorage struct {
	m sync.Map
}

var _ strava.TokenStorage = (*TokenStorage)(nil)

func (ts *TokenStorage) Get(_ context.Context, athleteID uint) (*strava.Token, error) {
	slog.Debug("get token", "athleteID", athleteID)

	v, ok := ts.m.Load(athleteID)
	if !ok {
		return nil, strava.ErrTokenNotFound
	}

	t, ok := v.(*strava.Token)
	if !ok {
		return nil, strava.ErrTokenNotFound
	}

	return t, nil
}

func (ts *TokenStorage) Save(_ context.Context, token *strava.Token) error {
	slog.Debug("save token", "athleteID", token.AthleteID)

	ts.m.Store(token.AthleteID, token)
	return nil
}
