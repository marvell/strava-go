package strava

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const (
	AthleteActivitiesPerPage = 100
)

func (c *Client) GetSummaryActivities(ctx context.Context, athleteID uint, from, to time.Time) ([]*SummaryActivity, error) {
	var activities []*SummaryActivity

	for i := 1; ; i++ {
		a, err := c.getSummaryActivities(ctx, athleteID, from, to, i, AthleteActivitiesPerPage)
		if err != nil {
			return nil, err
		}

		activities = append(activities, a...)

		if len(a) < AthleteActivitiesPerPage {
			break
		}
	}

	return activities, nil
}

func (c *Client) GetSummaryActivitiesWithCallback(ctx context.Context, athleteID uint, from, to time.Time, callback func([]*SummaryActivity) error) error {
	for i := 1; ; i++ {
		a, err := c.getSummaryActivities(ctx, athleteID, from, to, i, AthleteActivitiesPerPage)
		if err != nil {
			return err
		}

		if err := callback(a); err != nil {
			return fmt.Errorf("callback failed: %w", err)
		}

		if len(a) < AthleteActivitiesPerPage {
			break
		}
	}

	return nil
}

func (c *Client) UpdateActivity(ctx context.Context, athleteID, activityID uint, name, description string) error {
	reqBody := map[string]string{
		"name":        name,
		"description": description,
	}

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return fmt.Errorf("could not marshal request body: %w", err)
	}

	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/activities/%d", APIBaseURL, activityID), bytes.NewBuffer(jsonBody))
	if err != nil {
		return fmt.Errorf("could not create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	if _, err = c.call(ctx, athleteID, req, c.maxRetries); err != nil {
		return fmt.Errorf("could not call: %w", err)
	}

	return nil
}

func (c *Client) getSummaryActivities(ctx context.Context, athleteID uint, from, to time.Time, page, limit int) ([]*SummaryActivity, error) {
	params := url.Values{}
	params.Add("per_page", fmt.Sprint(limit))
	params.Add("page", fmt.Sprint(page))
	params.Add("after", fmt.Sprint(from.Unix()))
	params.Add("before", fmt.Sprint(to.Unix()))

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/athlete/activities?", APIBaseURL)+params.Encode(), nil)
	if err != nil {
		return nil, fmt.Errorf("could not create request: %w", err)
	}

	body, err := c.call(ctx, athleteID, req, c.maxRetries)
	if err != nil {
		return nil, fmt.Errorf("could not call: %w", err)
	}

	var v []*SummaryActivity
	err = json.Unmarshal(body, &v)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (c *Client) GetDetailedActivity(ctx context.Context, athleteID, activityID uint) (*DetailedActivity, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/activities/%d", APIBaseURL, activityID), nil)
	if err != nil {
		return nil, fmt.Errorf("could not create request: %w", err)
	}

	body, err := c.call(ctx, athleteID, req, c.maxRetries)
	if err != nil {
		return nil, fmt.Errorf("could not call strava: %w", err)
	}

	var v DetailedActivity
	err = json.Unmarshal(body, &v)
	if err != nil {
		return nil, err
	}

	return &v, nil

}
