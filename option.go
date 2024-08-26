package strava

import (
	"log/slog"
	"time"

	"golang.org/x/time/rate"
)

func WithLogger(l *slog.Logger) Option {
	return func(c *Client) {
		c.logger = l
	}
}

func WithRateLimiter(lmt *rate.Limiter) Option {
	return func(c *Client) {
		c.lmt = lmt
	}
}

func WithRetries(max uint, delay time.Duration) Option {
	return func(c *Client) {
		c.maxRetries = max
		c.retryDelay = delay
	}
}