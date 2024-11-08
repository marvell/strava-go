package strava

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

func (c *Client) InitWebhook(ctx context.Context) error {
	subs, err := c.GetSubscriptions(ctx)
	if err != nil {
		return fmt.Errorf("get subscriptions: %w", err)
	}

	// Clean up legacy subscriptions
	for _, sub := range subs {
		if err := c.DeleteSubscription(ctx, sub.ID); err != nil {
			return fmt.Errorf("delete subscription: %w", err)
		}
	}

	// Create a new subscription
	subID, err := c.CreateSubscription(ctx)
	if err != nil {
		return fmt.Errorf("create subscription: %w", err)
	}
	c.subscriptionID = subID

	return nil
}

func (c *Client) CloseWebhook(ctx context.Context) error {
	if c.subscriptionID == 0 {
		return nil
	}

	return c.DeleteSubscription(ctx, c.subscriptionID)
}

func (c *Client) CreateSubscription(ctx context.Context) (uint, error) {
	if c.webhookCallbackURL == "" {
		return 0, fmt.Errorf("webhook callback URL is not set")
	}

	endpoint := MustParseURL(APIBaseURL + "/push_subscriptions")
	endpoint.RawQuery = url.Values{
		"client_id":     {c.oacfg.ClientID},
		"client_secret": {c.oacfg.ClientSecret},
		"callback_url":  {c.webhookCallbackURL},
		"verify_token":  {c.webhookVerifyToken()},
	}.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint.String(), nil)
	if err != nil {
		return 0, fmt.Errorf("create request: %w", err)
	}

	body, err := c.call(ctx, 0, req, 0)
	if err != nil {
		return 0, fmt.Errorf("call: %w", err)
	}

	var v struct {
		ID uint `json:"id"`
	}

	err = json.Unmarshal(body, &v)
	if err != nil {
		return 0, fmt.Errorf("unmarshal response: %w", err)
	}

	return v.ID, nil
}

func (c *Client) DeleteSubscription(ctx context.Context, id uint) error {
	endpoint := MustParseURL(fmt.Sprintf("%s/push_subscriptions/%d", APIBaseURL, id))
	endpoint.RawQuery = url.Values{
		"client_id":     {c.oacfg.ClientID},
		"client_secret": {c.oacfg.ClientSecret},
	}.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, endpoint.String(), nil)
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}

	_, err = c.call(ctx, 0, req, 0)
	if err != nil {
		return fmt.Errorf("call: %w", err)
	}

	return nil
}

type Subscription struct {
	ID            uint      `json:"id"`
	ResourceState uint      `json:"resource_state"`
	ApplicationID uint      `json:"application_id"`
	CallbackURL   string    `json:"callback_url"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (c *Client) GetSubscriptions(ctx context.Context) ([]*Subscription, error) {
	endpoint := MustParseURL(APIBaseURL + "/push_subscriptions")
	endpoint.RawQuery = url.Values{
		"client_id":     {c.oacfg.ClientID},
		"client_secret": {c.oacfg.ClientSecret},
	}.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	body, err := c.call(ctx, 0, req, 0)
	if err != nil {
		return nil, fmt.Errorf("call: %w", err)
	}

	var v []*Subscription
	err = json.Unmarshal(body, &v)
	if err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}

	return v, nil
}

type EventHandler func(event Event) error

func (c *Client) RegisterEventHandler(handler EventHandler) error {
	if c.subscriptionID == 0 {
		return fmt.Errorf("webhook is not initialized")
	}

	c.eventHandlersLock.Lock()
	c.eventHandlers = append(c.eventHandlers, handler)
	c.eventHandlersLock.Unlock()

	return nil
}

func (c *Client) webhookVerifyToken() string {
	return fmt.Sprintf("strava-go-%s", c.oacfg.ClientID)
}

func (c *Client) WebhookCallback(w http.ResponseWriter, r *http.Request) {
	// Dump request
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		c.logger.WarnContext(r.Context(), "dump request with error", slog.Any("error", err))
	} else {
		c.logger.DebugContext(r.Context(), "webhook callback request", slog.Any("dump", dump))
	}

	// Create response wrapper
	rw := &responseWriter{ResponseWriter: w}

	switch r.Method {
	case http.MethodGet:
		c.webhookValidation(rw, r)
	case http.MethodPost:
		c.webhookEvent(rw, r)
	default:
		http.Error(rw, "method not allowed", http.StatusMethodNotAllowed)
	}

	c.logger.DebugContext(r.Context(), "webhook callback response", slog.String("dump", rw.String()))
}

func (c *Client) webhookValidation(w http.ResponseWriter, r *http.Request) {
	// Verify the token matches
	if token := r.URL.Query().Get("hub.verify_token"); token != c.webhookVerifyToken() {
		http.Error(w, "invalid verification token", http.StatusBadRequest)
		return
	}

	// Echo back the challenge
	challenge := r.URL.Query().Get("hub.challenge")

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]string{
		"hub.challenge": challenge,
	})
}

type Event struct {
	ObjectType     EventObjectType   `json:"object_type"`
	ObjectID       uint              `json:"object_id"`
	AspectType     EventAspectType   `json:"aspect_type"`
	Updates        map[string]string `json:"updates,omitempty"`
	OwnerID        uint              `json:"owner_id"`
	SubscriptionID uint              `json:"subscription_id"`
	EventTime      uint              `json:"event_time"`
}

type EventObjectType string

const (
	EventObjectTypeActivity EventObjectType = "activity"
	EventObjectTypeAthlete  EventObjectType = "athlete"
)

type EventAspectType string

const (
	EventAspectTypeCreate EventAspectType = "create"
	EventAspectTypeUpdate EventAspectType = "update"
	EventAspectTypeDelete EventAspectType = "delete"
)

func (c *Client) webhookEvent(w http.ResponseWriter, r *http.Request) {
	var event Event
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	c.eventHandlersLock.RLock()
	defer c.eventHandlersLock.RUnlock()

	for _, handler := range c.eventHandlers {
		go func(handler EventHandler) {
			if err := handler(event); err != nil {
				c.logger.Error("webhook event handled with error", "error", err)
			}
		}(handler)
	}

	w.WriteHeader(http.StatusOK)
}
