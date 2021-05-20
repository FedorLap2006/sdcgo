package sdcgo

import (
	"golang.org/x/time/rate"
	"net/http"
	"time"
)

// Session represents the connection to SD.C api.
type Session struct {
	Token       string
	Client      *http.Client
	RateLimiter *rate.Limiter
}

// New creates a new session with the given token.
func New(token string) *Session {
	return &Session{
		Token:       "SDC " + token,
		Client:      &http.Client{Timeout: time.Second * 10},
		RateLimiter: rate.NewLimiter(rate.Every(time.Second*2 + time.Second), 1),
	}
}
