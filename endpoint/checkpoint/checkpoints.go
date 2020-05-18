package checkpoint

import (
	"github.com/aftership/aftership-sdk-go/v2/endpoint/tracking"
	"github.com/aftership/aftership-sdk-go/v2/response"
)

// GetCheckpointParams is the additional parameters in checkpoint query
type GetCheckpointParams struct {
	// List of fields to include in the response. Use comma for multiple values.
	// Fields to include:slug,created_at,checkpoint_time,city,coordinates,country_iso3,
	// country_name,message,state,tag,zip
	// Default: none, Example: city,tag
	Fields string `url:"fields,omitempty" json:"fields,omitempty"`

	// Support Chinese to English translation for china-ems  and  china-post  only (Example: en)
	Lang string `url:"lang,omitempty" json:"lang,omitempty"`
}

// LastCheckpoint is the last checkpoint API response
type LastCheckpoint struct {
	ID             string              `json:"id,omitempty"`
	Slug           string              `json:"slug,omitempty"`
	TrackingNumber string              `json:"tracking_number"`
	Tag            string              `json:"tag"`
	Checkpoint     tracking.Checkpoint `json:"checkpoint"`
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
