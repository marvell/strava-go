package strava

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Lap struct {
	ID                 int64        `json:"id"`
	ResourceState      int          `json:"resource_state"`
	Activity           MetaActivity `json:"activity"`
	Athlete            MetaAthlete  `json:"athlete"`
	Name               string       `json:"name"`
	ElapsedTime        int          `json:"elapsed_time"`
	MovingTime         int          `json:"moving_time"`
	StartDate          time.Time    `json:"start_date"`
	StartDateLocal     time.Time    `json:"start_date_local"`
	Distance           float64      `json:"distance"`
	StartIndex         int          `json:"start_index"`
	EndIndex           int          `json:"end_index"`
	TotalElevationGain float64      `json:"total_elevation_gain"`
	AverageSpeed       float64      `json:"average_speed"`
	MaxSpeed           float64      `json:"max_speed"`
	AverageCadence     float64      `json:"average_cadence"`
	DeviceWatts        bool         `json:"device_watts"`
	AverageWatts       float64      `json:"average_watts"`
	LapIndex           int          `json:"lap_index"`
	Split              int          `json:"split"`
	PaceZone           int          `json:"pace_zone"`
	AverageHeartrate   float64      `json:"average_heartrate"`
	MaxHeartrate       float64      `json:"max_heartrate"`
}

type MetaActivity struct {
	ID            int64  `json:"id"`
	Visibility    string `json:"visibility"`
	ResourceState int    `json:"resource_state"`
}

type MetaAthlete struct {
	ID            int64 `json:"id"`
	ResourceState int   `json:"resource_state"`
}

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
