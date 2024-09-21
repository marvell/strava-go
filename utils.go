package strava

import (
	"math"
	"time"
)

// ConvertSpeedToPace converts speed in meters per second to pace (time per kilometer).
// It returns the pace as a time.Duration.
func ConvertSpeedToPace(speed float64) time.Duration {
	return time.Duration(math.Round(1000/speed)) * time.Second
}

// PaceToSpeed converts pace (time per kilometer) to speed in meters per second.
// It returns the speed as a float64.
func PaceToSpeed(pace time.Duration) float64 {
	return 3600. / pace.Seconds()
}
