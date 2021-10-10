package client

import (
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var random *rand.Rand

func init() {
	random = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// BackoffStrategy specifies a strategy for how long to wait between retries
type BackoffStrategy func(attemptNum int) time.Duration

// backoffStatusCheck checks if the status code is 429 or 503 and if
// the header key "Retry-After" exists, and it returns the time.Duration
// provided in the header
func backoffStatusCheck(resp *http.Response) *time.Duration {
	if resp != nil {
		if resp.StatusCode == http.StatusTooManyRequests || resp.StatusCode == http.StatusServiceUnavailable {
			if s, ok := resp.Header[retryAfterHeaderKey]; ok {
				if sleep, err := strconv.ParseInt(s[0], 10, 64); err == nil {
					m := time.Second * time.Duration(sleep)
					return &m
				}
				if after, err := time.Parse(time.RFC1123, s[0]); err == nil {
					n := time.Until(after)
					return &n
				}
			}
		}
	}
	return nil
}

// DefaultBackoffStrategy always returns 1 second
func DefaultBackoffStrategy(_ int) time.Duration {
	return 1 * time.Second
}

// ExponentialBackoffStrategy returns ever-increasing backoffs by a power of 2
func ExponentialBackoffStrategy(i int) time.Duration {
	return time.Duration(1<<uint(i)) * time.Second
}

// ExponentialJitterBackoffStrategy returns ever-increasing backoffs by a power of 2
// with +/- 0-33% to prevent synchronized requests
func ExponentialJitterBackoffStrategy(i int) time.Duration {
	return jitter(int(1 << uint(i)))
}

// LinearBackoffStrategy returns increasing durations, each a second longer than the last
func LinearBackoffStrategy(i int) time.Duration {
	return time.Duration(i) * time.Second
}

// LinearJitterBackoffStrategy returns increasing durations, each a second longer than the last
// with +/- 0-33% to prevent synchronized requests.
func LinearJitterBackoffStrategy(i int) time.Duration {
	return jitter(i)
}

// jitter keeps the +/- 0-33% logic in one place
func jitter(i int) time.Duration {
	ms := i * 1000

	maxJitter := ms / 3

	// ms ± rand
	ms += random.Intn(2*maxJitter) - maxJitter

	// a jitter of 0 messes up the time.Tick chan
	if ms <= 0 {
		ms = 1
	}

	return time.Duration(ms) * time.Millisecond
}
