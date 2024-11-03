package strava

import (
	"errors"
	"fmt"
)

var (
	ErrTokenNotFound = errors.New("token not found")
)

type APIError struct {
	Status   int    `json:"status"`
	Resource string `json:"resource"`
	Field    string `json:"field"`
	Code     string `json:"code"`
}

func (e APIError) Error() string {
	return fmt.Sprintf("%s.%s: %s", e.Resource, e.Field, e.Code)
}

type APIErrors struct {
	Errors  []APIError `json:"errors"`
	Message string     `json:"message"`
}

func (e APIErrors) Error() string {
	return e.Message
}
