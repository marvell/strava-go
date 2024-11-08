package main

import (
	"context"
	"log/slog"
	"net/http"
	"os/signal"
	"syscall"

	"github.com/caarlos0/env/v11"
	"github.com/marvell/strava-go"
	"github.com/marvell/strava-go/file"
)

type Config struct {
	BindAddr           string `env:"BIND_ADDR" envDefault:":8000"`
	StravaID           string `env:"STRAVA_ID,required"`
	StravaSecret       string `env:"STRAVA_SECRET,required"`
	WebhookCallbackURL string `env:"WEBHOOK_CALLBACK_URL,required"`
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	slog.SetLogLoggerLevel(slog.LevelDebug)

	var config Config
	err := env.Parse(&config)
	failIfError(err)

	ts, err := file.NewTokenStorage("../tokens")
	failIfError(err)

	cl := strava.NewClient(config.StravaID, config.StravaSecret, "", ts,
		strava.WithWebhookCallbackURL(config.WebhookCallbackURL),
		strava.WithDebug(),
	)

	http.HandleFunc("/callback", cl.WebhookCallback)

	go func() {
		err := http.ListenAndServe(config.BindAddr, nil)
		failIfError(err)
	}()

	err = cl.InitWebhook(ctx)
	failIfError(err)

	err = cl.RegisterEventHandler(func(event strava.Event) error {
		slog.Info("got new event", "event", event)
		return nil
	})
	failIfError(err)

	<-ctx.Done()

	// clean up
	err = cl.CloseWebhook(context.Background())
	failIfError(err)
}

func failIfError(err error) {
	if err != nil {
		panic(err)
	}
}
