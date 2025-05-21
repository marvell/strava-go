package redis

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/marvell/strava-go"
)

// TokenStorage implements the strava.TokenStorage interface using Redis.
type TokenStorage struct {
	client *redis.Client
	prefix string // e.g., "strava:token:"
}

var _ strava.TokenStorage = (*TokenStorage)(nil)

// NewTokenStorage creates a new Redis TokenStorage.
func NewTokenStorage(client *redis.Client, keyPrefix string) *TokenStorage {
	if keyPrefix == "" {
		keyPrefix = "strava:token:"
	}
	return &TokenStorage{
		client: client,
		prefix: keyPrefix,
	}
}

func (ts *TokenStorage) key(athleteID uint) string {
	return fmt.Sprintf("%s%d", ts.prefix, athleteID)
}

// Get retrieves a token from Redis.
func (ts *TokenStorage) Get(ctx context.Context, athleteID uint) (*strava.Token, error) {
	key := ts.key(athleteID)
	val, err := ts.client.Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, strava.ErrTokenNotFound
		}
		return nil, fmt.Errorf("redis get token for athlete %d (key %s): %w", athleteID, key, err)
	}

	var token strava.Token
	if err := json.Unmarshal([]byte(val), &token); err != nil {
		return nil, fmt.Errorf("redis unmarshal token for athlete %d (key %s): %w", athleteID, key, err)
	}

	return &token, nil
}

// Save stores a token in Redis.
func (ts *TokenStorage) Save(ctx context.Context, token *strava.Token) error {
	key := ts.key(token.AthleteID)
	val, err := json.Marshal(token)
	if err != nil {
		return fmt.Errorf("redis marshal token for athlete %d (key %s): %w", token.AthleteID, key, err)
	}

	if err := ts.client.Set(ctx, key, val, 0).Err(); err != nil { // 0 means no expiration
		return fmt.Errorf("redis set token for athlete %d (key %s): %w", token.AthleteID, key, err)
	}

	return nil
}
