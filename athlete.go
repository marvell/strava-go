package strava

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) GetAthlete(ctx context.Context, athleteID uint) (*DetailedAthlete, error) {
	req, err := http.NewRequest(http.MethodGet, APIBaseURL+"/athlete", nil)
	if err != nil {
		return nil, fmt.Errorf("could not create request: %w", err)
	}

	body, err := c.call(ctx, athleteID, req, c.maxRetries)
	if err != nil {
		return nil, fmt.Errorf("could not call strava: %w", err)
	}

	var v DetailedAthlete
	err = json.Unmarshal(body, &v)
	if err != nil {
		return nil, err
	}

	return &v, nil
}
