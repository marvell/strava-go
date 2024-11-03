package strava

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// AveragePace calculates the average pace for the lap
func (l Lap) AveragePace() time.Duration {
	return ConvertSpeedToPace(l.AverageSpeed)
}

// MaxPace calculates the maximum pace for the lap
func (l Lap) MaxPace() time.Duration {
	return ConvertSpeedToPace(l.MaxSpeed)
}

// MovingDuration returns the moving time of the lap as a time.Duration
func (l Lap) MovingDuration() time.Duration {
	return time.Duration(l.MovingTime) * time.Second
}

// ElapsedDuration returns the elapsed time of the lap as a time.Duration
func (l Lap) ElapsedDuration() time.Duration {
	return time.Duration(l.ElapsedTime) * time.Second
}

// GetActivityLaps retrieves the laps of an activity
func (c *Client) GetActivityLaps(ctx context.Context, athleteID, activityID uint) ([]*Lap, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/activities/%d/laps", APIBaseURL, activityID), nil)
	if err != nil {
		return nil, fmt.Errorf("could not create request: %w", err)
	}

	body, err := c.call(ctx, athleteID, req, c.maxRetries)
	if err != nil {
		return nil, fmt.Errorf("could not call: %w", err)
	}

	var laps []*Lap
	err = json.Unmarshal(body, &laps)
	if err != nil {
		return nil, err
	}

	return laps, nil
}
