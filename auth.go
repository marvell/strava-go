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
	Scope string
}

func (c *Client) AuthRedirectURL() string {
	return c.oacfg.RedirectURL
}

func (c *Client) AuthCodeURL() string {
	return c.AuthCodeURLWithRedirectURL("")
}

func (c *Client) AuthCodeURLWithRedirectURL(redirectURL string) string {
	oacfg := c.oacfg
	if redirectURL != "" {
		oacfg.RedirectURL = redirectURL
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
		Token: oauthToken,
		Scope: scope,
	}

	if err := c.tstore.Save(ctx, athleteID, token); err != nil {
		return 0, fmt.Errorf("could not save token: %w", err)
	}

	return athleteID, nil
}
