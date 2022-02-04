package aftership

import (
	"encoding/json"
)

// Error messages
const (
	errEmptyAPIKey                 = "invalid credentials: API Key must not be empty"
	errMissingTrackingNumber       = "tracking number is empty and must be provided"
	errMissingTrackingID           = "tracking id is empty and must be provided"
	errMissingSlugOrTrackingNumber = "slug or tracking number is empty, both of them must be provided"
	errExceedRateLimt              = "rate limit is exceeded, please wait util %s"
)

// APIError is the error in AfterShip API calls
type APIError struct {
	Code    int    `json:"code"`
	Type    string `json:"type"`
	Message string `json:"message"`
	Path    string `json:"path"`

	// HTTPStatusCode represents the original HTTP status code which was returned by the request from the AfterShip's API
	HTTPStatusCode int `json:"http_status_code"`
}

// Error serializes the error object to JSON and returns it as a string.
func (e *APIError) Error() string {
	ret, _ := json.Marshal(e)
	return string(ret)
}

// TooManyRequestsError is the too many requests error in AfterShip API calls
type TooManyRequestsError struct {
	APIError
	RateLimit *RateLimit `json:"rate_limit"`
}

// Error serializes the error object to JSON and returns it as a string.
func (e *TooManyRequestsError) Error() string {
	ret, _ := json.Marshal(e)
	return string(ret)
}
