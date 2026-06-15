package riot

import "errors"

var (
	ErrPlayerNotFound = errors.New("player not found")
	ErrMatchNotFound = errors.New("match not found")
	ErrRateLimitExceeded = errors.New("Riot API rate limit exceeded")
	ErrUnknown = errors.New("unknown error occurred")
)
