package postgres

import (
	"gorm.io/gorm"

	"github.com/marvell/strava-go"
)

type Token struct {
	*gorm.Model
	*strava.Token
}

func (t Token) TableName() string {
	return "strava_tokens"
}
