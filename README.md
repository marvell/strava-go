# strava-go
A Go library for interacting with the Strava API.

## Installation

```
go get github.com/marvell/strava-go
```

## Usage

```go
import (
	"github.com/marvell/strava-go"
	"github.com/marvell/strava-go/inmemory"
)

ts := &inmemory.TokenStorage{}
cl := strava.NewClient(id, secret, redirectURL, ts)

ath, err := cl.GetAthlete(context.Background(), uint(athleteID))
if err != nil {
	panic(err)
}

slog.Info(fmt.Sprintf("athlete: %+v", ath))
```

## Documentation

For detailed documentation and examples, please visit the [GoDoc](https://pkg.go.dev/github.com/marvell/strava-go) page.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
