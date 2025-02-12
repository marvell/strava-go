package strava

import (
	"log/slog"
	"net/http"
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

func WithDebug() Option {
	return func(c *Client) {
		c.debug = true
	}
}

func WithTransport(t *http.Transport) Option {
	return func(c *Client) {
		c.transport = t
	}
}

func WithWebhookCallbackURL(url string) Option {
	return func(c *Client) {
		c.webhookCallbackURL = url
	}
}

func WithScopes(scopes ...string) Option {
	return func(c *Client) {
		c.oacfg.Scopes = scopes
	}
}
