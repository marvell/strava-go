module github.com/marvell/strava-go/examples/webhook

go 1.23.2

replace github.com/marvell/strava-go => ../..

require (
	github.com/caarlos0/env/v11 v11.2.2
	github.com/marvell/strava-go v0.0.0-00010101000000-000000000000
)

require (
	golang.org/x/oauth2 v0.22.0 // indirect
	golang.org/x/time v0.6.0 // indirect
)
