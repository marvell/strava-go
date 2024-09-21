package strava

import "errors"

var (
	ErrTokenNotFound = errors.New("token not found")
)

type APIError struct {
	Errors []struct {
		Code     string `json:"code"`
		Field    string `json:"field"`
		Resource string `json:"resource"`
	} `json:"errors"`
	Message string `json:"message"`
}

func (e APIError) Error() string {
	return e.Message
}
