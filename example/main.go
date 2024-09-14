package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/caarlos0/env/v11"
	"github.com/marvell/strava-go"
	"github.com/marvell/strava-go/file"
)

type Config struct {
	ID          string `env:"STRAVA_ID,required"`
	Secret      string `env:"STRAVA_SECRET,required"`
	RedirectURL string `env:"STRAVA_REDIRECT_URL,required"`
	AthleteID   uint   `env:"STRAVA_ATHLETE_ID"`
}

func main() {
	slog.SetLogLoggerLevel(slog.LevelDebug)

	var config Config
	if err := env.Parse(&config); err != nil {
		panic(err)
	}

	ctx := context.Background()

	ts, err := file.NewTokenStorage("./tokens")
	if err != nil {
		panic(err)
	}

	cl := strava.NewClient(config.ID, config.Secret, config.RedirectURL, ts)

	athleteID := config.AthleteID
	if athleteID == 0 {
		athleteID = auth(ctx, cl)
	} else {
		_, err := ts.Get(ctx, athleteID)
		if err != nil {
			if !errors.Is(err, strava.ErrTokenNotFound) {
				panic(err)
			}

			id := auth(ctx, cl)
			if id != athleteID {
				slog.Warn(fmt.Sprintf("STRAVA_ATHLETE_ID (%d) != authorized athlete ID (%d)", athleteID, id))
			}
			athleteID = id
		}
	}

	ath, err := cl.GetAthlete(ctx, athleteID)
	if err != nil {
		panic(err)
	}

	slog.Info(fmt.Sprintf("athlete: %+v", ath))
}

func auth(ctx context.Context, cl *strava.Client) uint {
	fmt.Printf("You need to authorize first.\n")
	fmt.Printf("Strava auth URL: %s\n", cl.AuthCodeURL())

	fmt.Print("Enter code: ")
	var stravaCode string
	_, err := fmt.Scanln(&stravaCode)
	if err != nil {
		panic(err)
	}

	athleteID, err := cl.AuthExchange(ctx, stravaCode)
	if err != nil {
		panic(err)
	}

	return athleteID
}
