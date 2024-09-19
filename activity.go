package strava

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	AthleteActivitiesPerPage = 100
)

type ActivityIndex struct {
	ID      uint `json:"id"`
	Athlete struct {
		ID uint `json:"id"`
	} `json:"athlete"`
	StartDate time.Time `json:"start_date"`
}

type Activity struct {
	ActivityIndex
	Data []byte `json:"-"`
}

func (ad *Activity) UnmarshalJSON(b []byte) error {
	if err := json.Unmarshal(b, &ad.ActivityIndex); err != nil {
		return err
	}
	ad.Data = b
	return nil
}

func (c *Client) GetActivities(ctx context.Context, athleteID uint, from, to time.Time) ([]*Activity, error) {
	var activities []*Activity

	for i := 1; ; i++ {
		a, err := c.getActivities(ctx, athleteID, from, to, i)
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

func (c *Client) GetActivitiesWithCallback(ctx context.Context, athleteID uint, from, to time.Time, callback func([]*Activity) error) error {
	for i := 1; ; i++ {
		a, err := c.getActivities(ctx, athleteID, from, to, i)
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

func (c *Client) UpdateActivityDescription(ctx context.Context, athleteID, activityID uint, description string) error {
	reqBody := strings.NewReader(url.Values{
		"description": []string{description},
	}.Encode())

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/activities/%d", APIBaseURL, activityID), reqBody)
	if err != nil {
		return fmt.Errorf("could not create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	if _, err = c.call(ctx, athleteID, req, c.maxRetries); err != nil {
		return fmt.Errorf("could not call: %w", err)
	}

	return nil
}

func (c *Client) getActivities(ctx context.Context, athleteID uint, from, to time.Time, page int) ([]*Activity, error) {
	params := url.Values{}
	params.Add("per_page", fmt.Sprint(AthleteActivitiesPerPage))
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

	var v []*Activity
	err = json.Unmarshal(body, &v)
	if err != nil {
		return nil, err
	}

	return v, nil
}
