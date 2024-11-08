package strava

import (
	"bytes"
	"fmt"
	"math"
	"net/http"
	"net/url"
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

func MustParseURL(rawurl string) *url.URL {
	u, err := url.Parse(rawurl)
	if err != nil {
		panic(err)
	}
	return u
}

type responseWriter struct {
	http.ResponseWriter
	statusCode int
	body       bytes.Buffer
}

// WriteHeader captures the status code.
func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

// Write captures the body content.
func (rw *responseWriter) Write(b []byte) (int, error) {
	rw.body.Write(b)
	return rw.ResponseWriter.Write(b)
}

func (rw *responseWriter) String() string {
	statusCode := http.StatusOK
	if rw.statusCode != 0 {
		statusCode = rw.statusCode
	}

	return fmt.Sprintf("status=%d body=%q", statusCode, rw.body.String())
}
