package strava

import (
	"net/url"
	"strings"
	"testing"

	"github.com/gookit/goutil/testutil/assert"
)

func TestClient_AuthCodeURL(t *testing.T) {
	// arrange
	clientID := "client_id"
	clientSecret := "client_secret"
	redirectURL := "http://localhost:8080/callback"
	scopes := []string{"activity:read", "activity:write"}
	c := NewClient(clientID, clientSecret, redirectURL, nil, WithScopes(scopes...))
	u, _ := url.Parse("https://www.strava.com/oauth/authorize")
	u.RawQuery = url.Values{
		"access_type":   {"offline"},
		"client_id":     {clientID},
		"redirect_uri":  {redirectURL},
		"response_type": {"code"},
		"scope":         {strings.Join(scopes, " ")},
		"state":         {OAuthStaticState},
	}.Encode()
	want := u.String()

	// act
	got := c.AuthCodeURL()

	// assert
	assert.Equal(t, want, got)
}
