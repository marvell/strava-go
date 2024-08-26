package postgres

import (
	"golang.org/x/oauth2"
	"gorm.io/gorm"
)

type Token struct {
	*gorm.Model
	*oauth2.Token
}

func (t Token) TableName() string {
	return "strava_tokens"
}
