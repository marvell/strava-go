package strava

import (
	"context"
	"fmt"

	"golang.org/x/oauth2"
)

const (
	OAuthStaticState = "strava-go"
)

type Token struct {
	*oauth2.Token
	AthleteID uint   `json:"athlete_id"`
	Scope     string `json:"scope"`
}

func (c *Client) AuthRedirectURL() string {
	return c.oacfg.RedirectURL
}

func (c *Client) AuthCodeURL(redirectURL string, scopes []string) string {
	oacfg := c.oacfg
	oacfg.RedirectURL = redirectURL
	if redirectURL != "" {
		oacfg.RedirectURL = redirectURL
	}
	if len(scopes) > 0 {
		oacfg.Scopes = scopes
	}

	return oacfg.AuthCodeURL(OAuthStaticState, oauth2.AccessTypeOffline)
}

func (c *Client) AuthExchange(ctx context.Context, code, scope, state string) (uint, error) {
	if state != OAuthStaticState {
		return 0, fmt.Errorf("invalid state: %s", state)
	}

	oauthToken, err := c.oacfg.Exchange(ctx, code)
	if err != nil {
		return 0, fmt.Errorf("could not exchange code for token: %w", err)
	}

	extra, ok := oauthToken.Extra("athlete").(map[string]any)
	if !ok {
		return 0, fmt.Errorf("could not get athlete data from token")
	}
	athleteID := uint(extra["id"].(float64))

	token := &Token{
		Token:     oauthToken,
		AthleteID: athleteID,
		Scope:     scope,
	}

	if err := c.tstore.Save(ctx, token); err != nil {
		return 0, fmt.Errorf("could not save token: %w", err)
	}

	return athleteID, nil
}
