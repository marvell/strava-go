package strava

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net"
	"net/http"
	"net/http/httputil"
	"sync"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/time/rate"
)

const (
	OAuthBaseURL = "https://www.strava.com/oauth"
	APIBaseURL   = "https://www.strava.com/api/v3"

	HTTPClientTimeout = 5 * time.Second
)

type TokenStorage interface {
	Get(ctx context.Context, athleteID uint) (*Token, error)
	Save(ctx context.Context, token *Token) error
}

type Option func(*Client)

func NewClient(id, secret, redirectURL string, ts TokenStorage, opts ...Option) *Client {
	oacfg := oauth2.Config{
		ClientID:     id,
		ClientSecret: secret,
		Endpoint: oauth2.Endpoint{
			AuthURL:  OAuthBaseURL + "/authorize",
			TokenURL: OAuthBaseURL + "/token",
		},
		RedirectURL: redirectURL,
		Scopes:      []string{"read"},
	}

	c := &Client{
		oacfg:  oacfg,
		tstore: ts,
		lmt:    nil,
		logger: slog.Default(),
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

type Client struct {
	transport *http.Transport

	oacfg  oauth2.Config
	tstore TokenStorage

	lmt *rate.Limiter

	maxRetries uint
	retryDelay time.Duration

	webhookCallbackURL string
	subscriptionID     uint
	eventHandlers      []EventHandler
	eventHandlersLock  sync.RWMutex

	logger *slog.Logger
	debug  bool
}

func (c *Client) call(ctx context.Context, athleteID uint, req *http.Request, retries uint) ([]byte, error) {
	if c.lmt != nil && !c.lmt.Allow() {
		c.logger.Warn("rate limit exceeded: waiting...")

		if err := c.lmt.Wait(ctx); err != nil {
			return nil, fmt.Errorf("rate limiter: wait: %w", err)
		}
	}

	httpClient, err := c.getHttpClientFor(ctx, athleteID)
	if err != nil {
		return nil, fmt.Errorf("get http client for %d athlete: %w", athleteID, err)
	}

	if c.debug {
		reqDump, err := httputil.DumpRequestOut(req, false)
		if err != nil {
			c.logger.WarnContext(ctx, "dump request", slog.Any("error", err))
		} else {
			c.logger.DebugContext(ctx, fmt.Sprintf("strava api request: %s %s", req.Method, req.URL.Path), slog.Any("dump", reqDump))
		}
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		if err, ok := err.(net.Error); ok && err.Timeout() {
			if retries > 0 {
				c.logger.WarnContext(ctx, "request timeout: retrying", slog.Any("error", err))

				time.Sleep(c.retryDelay)
				return c.call(ctx, athleteID, req, retries-1)
			}
		}

		return nil, err
	}
	defer resp.Body.Close()

	if c.debug {
		respDump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			c.logger.WarnContext(ctx, "dump response", slog.Any("error", err))
		} else {
			c.logger.DebugContext(ctx, fmt.Sprintf("strava api response: %s %s", req.Method, req.URL.Path), slog.Any("dump", respDump))
		}
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= http.StatusBadRequest {
		var details any = body

		var apiErr Fault
		if err := json.Unmarshal(body, &apiErr); err == nil {
			details = apiErr
		}

		return nil, fmt.Errorf("API error (status code %d): %v", resp.StatusCode, details)
	}

	return body, nil
}

func (c *Client) getHttpClient(_ context.Context) *http.Client {
	hc := &http.Client{
		Timeout: HTTPClientTimeout,
	}
	if c.transport != nil {
		hc.Transport = c.transport
	}

	return hc
}

func (c *Client) getHttpClientFor(ctx context.Context, athleteID uint) (*http.Client, error) {
	if athleteID == 0 {
		return c.getHttpClient(ctx), nil
	}

	token, err := c.token(ctx, athleteID)
	if err != nil {
		return nil, err
	}

	if c.transport != nil {
		ctx = context.WithValue(ctx, oauth2.HTTPClient, &http.Client{
			Transport: c.transport,
		})
	}

	hc := c.oacfg.Client(ctx, token)
	hc.Timeout = HTTPClientTimeout
	return hc, nil
}

func (c *Client) token(ctx context.Context, athleteID uint) (*oauth2.Token, error) {
	token, err := c.tstore.Get(ctx, athleteID)
	if err != nil {
		return nil, fmt.Errorf("get token from %T: %w", c.tstore, err)
	}

	if !token.Valid() {
		oauthToken, err := c.oacfg.TokenSource(ctx, token.Token).Token()
		if err != nil {
			return nil, fmt.Errorf("refresh token: %w", err)
		}
		token.Token = oauthToken

		err = c.tstore.Save(ctx, token)
		if err != nil {
			return nil, fmt.Errorf("save token to %T: %w", c.tstore, err)
		}
	}

	return token.Token, nil
}
