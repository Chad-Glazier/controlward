package riot

import (
	"errors"
	"net/http"
)

var (
	ErrBadRequest = errors.New("riot: bad request")
	ErrUnauthorized = errors.New("riot: unauthorized request")
	ErrForbidden = errors.New("riot: request forbidden")
	ErrNotFound = errors.New("riot: requested data not found")
	ErrMethodNotAllowed = errors.New("riot: request method not allowed")
	ErrUnsupportedMediaType = errors.New("riot: unsupported media type")
	ErrRateLimit = errors.New("riot: api rate limit exceeded")
	ErrInternal = errors.New("riot: internal server error")
	ErrBadGateway = errors.New("riot: bad gateway")
	ErrServiceUnavailable = errors.New("riot: service unavailable")
	ErrGatewayTimeout = errors.New("riot: gateway timeout")
	ErrUnknown = errors.New("riot: unknown error")
)

// Maps an error response from the Riot API to a predefined error.
func riotError(resp *http.Response) error {
	switch resp.StatusCode {
	case http.StatusBadRequest:
		return ErrBadRequest
	case http.StatusUnauthorized:
		return ErrUnauthorized
	case http.StatusForbidden:
		return ErrForbidden
	case http.StatusNotFound:
		return ErrNotFound
	case http.StatusMethodNotAllowed:
		return ErrMethodNotAllowed
	case http.StatusUnsupportedMediaType:
		return ErrUnsupportedMediaType
	case http.StatusTooManyRequests:
		return ErrRateLimit
	case http.StatusInternalServerError:
		return ErrInternal
	case http.StatusBadGateway:
		return ErrBadGateway
	case http.StatusServiceUnavailable:
		return ErrServiceUnavailable
	case http.StatusGatewayTimeout:
		return ErrGatewayTimeout
	default:
		return ErrUnknown
	}
}
