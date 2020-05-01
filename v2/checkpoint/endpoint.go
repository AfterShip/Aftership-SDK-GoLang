package checkpoint

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/aftership/aftership-sdk-go/v2/common"
	"github.com/aftership/aftership-sdk-go/v2/error"
	"github.com/aftership/aftership-sdk-go/v2/request"
)

// Endpoint provides the interface for all checkpoint API calls
type Endpoint interface {
	// GetLastCheckpoint returns the tracking information of the last checkpoint of a single tracking.
	GetLastCheckpoint(param common.SingleTrackingParam, fields string, lang string) (LastCheckpoint, *error.AfterShipError)
}

// EndpointImpl is the implementaion of checkpoint endpoint
type EndpointImpl struct {
	request request.APIRequest
}

// NewEnpoint creates a instance of checkpoint endpoint
func NewEnpoint(req request.APIRequest) Endpoint {
	return &EndpointImpl{
		request: req,
	}
}

// GetLastCheckpoint Return the tracking information of the last checkpoint of a single tracking.
func (impl *EndpointImpl) GetLastCheckpoint(param common.SingleTrackingParam, fields string, lang string) (LastCheckpoint, *error.AfterShipError) {
	url, err := buildLastCheckpointURL(param, fields, lang)

	var envelope LastCheckpointEnvelope
	err = impl.request.MakeRequest("GET", url, nil, &envelope)
	if err != nil {
		return LastCheckpoint{}, err
	}

	return envelope.Data, nil
}

func buildLastCheckpointURL(param common.SingleTrackingParam, fields string, lang string) (string, *error.AfterShipError) {
	checkpointURL, err := param.BuildTrackingURL("last_checkpoint", "")
	if err != nil {
		return "", err
	}

	if fields != "" || lang != "" {
		extraParams := []string{}

		if fields != "" {
			extraParams = append(extraParams, fmt.Sprintf("fields=%s", url.QueryEscape(fields)))
		}

		if lang != "" {
			extraParams = append(extraParams, fmt.Sprintf("lang=%s", url.QueryEscape(lang)))
		}

		checkpointURL += fmt.Sprintf("?%s", strings.Join(extraParams, "&"))
	}

	return checkpointURL, nil
}
