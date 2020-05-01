package checkpoint

import "github.com/aftership/aftership-sdk-go/v2/response"

// Checkpoint represents a checkpoint returned by the Aftership API
type Checkpoint struct {
	Slug           string   `json:"slug,omitempty"`
	CreatedAt      string   `json:"created_at,omitempty"`
	CheckPointTime string   `json:"checkpoint_time,omitempty"`
	City           string   `json:"city,omitempty"`
	Coordinates    []string `json:"coordinates,omitempty"`
	CountryIso3    string   `json:"country_iso3,omitempty"`
	CountryName    string   `json:"country_name,omitempty"`
	Message        string   `json:"message,omitempty"`
	State          string   `json:"state,omitempty"`
	Location       string   `json:"location,omitempty"`
	Tag            string   `json:"tag,omitempty"`
	Zip            string   `json:"zip,omitempty"`
}

// LastCheckpoint is the last checkpoint API response
type LastCheckpoint struct {
	ID             string     `json:"id,omitempty"`
	Slug           string     `json:"slug,omitempty"`
	TrackingNumber string     `json:"tracking_number"`
	Tag            string     `json:"tag"`
	Checkpoint     Checkpoint `json:"checkpoint"`
}

// LastCheckpointEnvelope is the message envelope for the last checkpoint API responses
type LastCheckpointEnvelope struct {
	Meta response.Meta  `json:"meta"`
	Data LastCheckpoint `json:"data"`
}

// GetMeta returns the response meta
func (e *LastCheckpointEnvelope) GetMeta() response.Meta {
	return e.Meta
}
