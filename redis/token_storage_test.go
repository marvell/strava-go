package redis

import (
	"context"
	"errors"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/marvell/strava-go"
	"golang.org/x/oauth2"
)

var testRedisClient *redis.Client

func TestMain(m *testing.M) {
	// Connection options for Redis. Using an environment variable for the address
	// allows flexibility for different test environments (e.g., CI vs local).
	redisAddr := os.Getenv("STRAVA_TEST_REDIS_ADDR")
	if redisAddr == "" {
		redisAddr = "localhost:6379" // Default address
	}

	testRedisClient = redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})

	// Ping Redis to ensure connection before running tests
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := testRedisClient.Ping(ctx).Err(); err != nil {
		fmt.Printf("Could not connect to Redis at %s: %v\n", redisAddr, err)
		fmt.Println("Skipping Redis integration tests. Ensure Redis is running and accessible.")
		// os.Exit(1) // Or skip tests if Redis is essential
		return // Exit TestMain without running tests if Redis is not available
	}

	// Run tests
	code := m.Run()

	os.Exit(code)
}

func flushRedis(ctx context.Context, t *testing.T) {
	if err := testRedisClient.FlushDB(ctx).Err(); err != nil {
		t.Fatalf("Failed to flush Redis: %v", err)
	}
}

func TestNewTokenStorage(t *testing.T) {
	tsDefault := NewTokenStorage(testRedisClient, "")
	if tsDefault.prefix != "strava:token:" {
		t.Errorf("Expected default prefix 'strava:token:', got '%s'", tsDefault.prefix)
	}

	customPrefix := "mycustom:prefix:"
	tsCustom := NewTokenStorage(testRedisClient, customPrefix)
	if tsCustom.prefix != customPrefix {
		t.Errorf("Expected custom prefix '%s', got '%s'", customPrefix, tsCustom.prefix)
	}
}

func TestTokenStorage_SaveAndGet(t *testing.T) {
	ctx := context.Background()
	flushRedis(ctx, t)

	ts := NewTokenStorage(testRedisClient, "test:strava:token:")
	athleteID := uint(12345)
	expiry := time.Now().Add(1 * time.Hour).Round(time.Second) // Round for consistent comparison

	token := &strava.Token{
		Token: &oauth2.Token{
			AccessToken:  "test_access_token",
			RefreshToken: "test_refresh_token",
			Expiry:       expiry,
			TokenType:    "Bearer",
		},
		AthleteID: athleteID,
		Scope:     "read,activity:read_all",
	}

	err := ts.Save(ctx, token)
	if err != nil {
		t.Fatalf("Save() error = %v, wantErr nil", err)
	}

	retrievedToken, err := ts.Get(ctx, athleteID)
	if err != nil {
		t.Fatalf("Get() error = %v, wantErr nil", err)
	}

	if retrievedToken.AccessToken != token.AccessToken {
		t.Errorf("Retrieved AccessToken = %s, want %s", retrievedToken.AccessToken, token.AccessToken)
	}
	if retrievedToken.RefreshToken != token.RefreshToken {
		t.Errorf("Retrieved RefreshToken = %s, want %s", retrievedToken.RefreshToken, token.RefreshToken)
	}
	// Compare time with tolerance or after rounding due to potential precision differences in storage/retrieval
	if !retrievedToken.Expiry.Equal(token.Expiry) {
		t.Errorf("Retrieved Expiry = %v, want %v", retrievedToken.Expiry, token.Expiry)
	}
	if retrievedToken.AthleteID != token.AthleteID {
		t.Errorf("Retrieved AthleteID = %d, want %d", retrievedToken.AthleteID, token.AthleteID)
	}
	if retrievedToken.Scope != token.Scope {
		t.Errorf("Retrieved Scope = %s, want %s", retrievedToken.Scope, token.Scope)
	}
}

func TestTokenStorage_Get_NotFound(t *testing.T) {
	ctx := context.Background()
	flushRedis(ctx, t)

	ts := NewTokenStorage(testRedisClient, "test:strava:token:")
	athleteID := uint(67890)

	_, err := ts.Get(ctx, athleteID)
	if !errors.Is(err, strava.ErrTokenNotFound) {
		t.Fatalf("Get() error = %v, want %v", err, strava.ErrTokenNotFound)
	}
}

func TestTokenStorage_Save_Overwrite(t *testing.T) {
	ctx := context.Background()
	flushRedis(ctx, t)

	ts := NewTokenStorage(testRedisClient, "test:strava:token:")
	athleteID := uint(11223)
	initialToken := &strava.Token{
		Token: &oauth2.Token{
			AccessToken: "initial_access_token",
		},
		AthleteID: athleteID,
	}
	updatedToken := &strava.Token{
		Token: &oauth2.Token{
			AccessToken: "updated_access_token",
		},
		AthleteID: athleteID,
	}

	// Save initial token
	if err := ts.Save(ctx, initialToken); err != nil {
		t.Fatalf("Save() initial token error = %v", err)
	}

	// Save updated token for the same athlete
	if err := ts.Save(ctx, updatedToken); err != nil {
		t.Fatalf("Save() updated token error = %v", err)
	}

	retrievedToken, err := ts.Get(ctx, athleteID)
	if err != nil {
		t.Fatalf("Get() after overwrite error = %v", err)
	}

	if retrievedToken.AccessToken != updatedToken.AccessToken {
		t.Errorf("Retrieved AccessToken = %s, want %s (updated)", retrievedToken.AccessToken, updatedToken.AccessToken)
	}
}
