package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"strconv"

	"github.com/marvell/strava-go"
	"github.com/marvell/strava-go/inmemory"
)

func main() {
	slog.SetLogLoggerLevel(slog.LevelDebug)

	id := os.Getenv("STRAVA_ID")
	secret := os.Getenv("STRAVA_SECRET")
	redirectURL := os.Getenv("STRAVA_REDIRECT_URL")

	athleteID, err := strconv.ParseUint(os.Getenv("STRAVA_ATHLETE_ID"), 10, 64)
	if err != nil {
		panic(err)
	}

	ts := &inmemory.TokenStorage{}
	cl := strava.NewClient(id, secret, redirectURL, ts)

	ctx := context.Background()

	ath, err := cl.GetAthlete(ctx, uint(athleteID))
	if err != nil {
		panic(err)
	}

	slog.Info(fmt.Sprintf("athlete: %+v", ath))
}
