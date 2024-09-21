# strava-go

[![Strava API Docs](https://img.shields.io/badge/Strava%20API-Reference-orange)](https://developers.strava.com/docs/reference/)

strava-go is a Go library for interacting with the Strava API. It provides a simple and efficient way to authenticate and make requests to Strava's API endpoints. The library is designed to work with the official Strava API, which is documented in the [Strava API and SDK Reference](https://developers.strava.com/docs/reference/). For interactive exploration and testing of API endpoints, you can use the [Strava API Playground](https://developers.strava.com/playground/), which provides a user-friendly interface to try out various API calls and see their responses.

## Features

- OAuth2 authentication
- Multiple token storage options
- Rate limiting support
- Configurable retries
- Debugging options

## Installation

To install strava-go, use `go get`:

```
go get github.com/marvell/strava-go
```

## Usage

Here's a basic example of how to use the strava-go library:

```go
import (
    "context"
    "fmt"
    "log/slog"

    "github.com/marvell/strava-go"
    "github.com/marvell/strava-go/file"
)

func main() {
    ctx := context.Background()

    // Create a token storage
    ts, err := file.NewTokenStorage("./tokens")
    if err != nil {
        panic(err)
    }

    // Create a new Strava client
    cl := strava.NewClient(
        "your-client-id",
        "your-client-secret",
        "your-redirect-url",
        ts,
        strava.WithDebug(), // Optional: Enable debug mode
    )

    // Authenticate (if needed)
    athleteID := auth(ctx, cl)

    // Get athlete information
    ath, err := cl.GetAthlete(ctx, athleteID)
    if err != nil {
        panic(err)
    }

    slog.Info(fmt.Sprintf("athlete: %+v", ath))
}

func auth(ctx context.Context, cl *strava.Client) uint {
    // ... (authentication process)
}
```

## Configuration

The `NewClient` function accepts several options to customize the client's behavior:

- `WithLogger`: Set a custom logger
- `WithRateLimiter`: Set a rate limiter
- `WithRetries`: Configure retry behavior
- `WithDebug`: Enable debug mode

## Token Storage

The library provides several token storage implementations to suit different needs:

### File-based Storage

Store tokens in local files:

```go
import "github.com/marvell/strava-go/file"

ts, err := file.NewTokenStorage("./tokens")
```

### PostgreSQL Storage

Store tokens in a PostgreSQL database:

```go
import (
    "github.com/marvell/strava-go/postgres"
    "gorm.io/gorm"
)

db, err := gorm.Open(/* your database connection */)
ts, err := postgres.NewTokenStorage(db)
```

### In-Memory Storage

Store tokens in memory (useful for testing or short-lived applications):

```go
import "github.com/marvell/strava-go/inmemory"

ts := &inmemory.TokenStorage{}
```

### Custom Storage

You can implement your own token storage by satisfying the `TokenStorage` interface:

```go
type TokenStorage interface {
    Get(ctx context.Context, athleteID uint) (*oauth2.Token, error)
    Save(ctx context.Context, athleteID uint, token *oauth2.Token) error
}
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
