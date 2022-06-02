package aftership

import (
	"context"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
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

	AdditionalField
}

// LastCheckpoint is the last checkpoint API response
type LastCheckpoint struct {
	ID             string     `json:"id,omitempty"`
	Slug           string     `json:"slug,omitempty"`
	TrackingNumber string     `json:"tracking_number,omitempty"`
	Tag            string     `json:"tag,omitempty"`
	Subtag         string     `json:"subtag,omitempty"`
	SubtagMessage  string     `json:"subtag_message,omitempty"`
	Checkpoint     Checkpoint `json:"checkpoint"`
}

// GetLastCheckpoint returns the tracking information of the last checkpoint of a single tracking.
func (client *Client) GetLastCheckpoint(ctx context.Context, identifier TrackingIdentifier, params GetCheckpointParams) (LastCheckpoint, error) {
	uriPath, err := identifier.URIPath()
	if err != nil {
		return LastCheckpoint{}, errors.Wrap(err, "error getting last checkpoint")
	}

	uriPath = fmt.Sprintf("/last_checkpoint%s", uriPath)
	var lastCheckpoint LastCheckpoint
	err = client.makeRequest(ctx, http.MethodGet, uriPath, params, nil, &lastCheckpoint)
	return lastCheckpoint, err
}
