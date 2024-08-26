package strava

import (
	"context"
	"fmt"

	"golang.org/x/oauth2"
)

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
	return oacfg.AuthCodeURL("state", oauth2.AccessTypeOffline)
}

func (c *Client) AuthExchange(ctx context.Context, code string) (uint, error) {
	token, err := c.oacfg.Exchange(ctx, code)
	if err != nil {
		return 0, fmt.Errorf("could not exchange code for token: %w", err)
	}

	extra, ok := token.Extra("athlete").(map[string]any)
	if !ok {
		return 0, fmt.Errorf("could not get athlete data from token")
	}
	athleteID := uint(extra["id"].(float64))

	if err := c.tstore.Save(ctx, athleteID, token); err != nil {
		return 0, fmt.Errorf("could not save token: %w", err)
	}

	return athleteID, nil
}
