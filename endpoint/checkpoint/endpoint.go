package checkpoint

import (
	"github.com/aftership/aftership-sdk-go/v2/common"
	"github.com/aftership/aftership-sdk-go/v2/error"
	"github.com/aftership/aftership-sdk-go/v2/request"
)

// Endpoint provides the interface for all checkpoint API calls
type Endpoint interface {
	// GetLastCheckpoint returns the tracking information of the last checkpoint of a single tracking.
	GetLastCheckpoint(param common.SingleTrackingParam, optionalParams *GetCheckpointParams) (LastCheckpoint, *error.AfterShipError)
}

// EndpointImpl is the implementaion of checkpoint endpoint
type EndpointImpl struct {
	request request.APIRequest
}

// NewEndpoint creates a instance of checkpoint endpoint
func NewEndpoint(req request.APIRequest) Endpoint {
	return &EndpointImpl{
		request: req,
	}
}

// GetLastCheckpoint returns the tracking information of the last checkpoint of a single tracking.
func (impl *EndpointImpl) GetLastCheckpoint(param common.SingleTrackingParam, optionalParams *GetCheckpointParams) (LastCheckpoint, *error.AfterShipError) {
	url, err := buildLastCheckpointURL(param, optionalParams)
	if err != nil {
		return LastCheckpoint{}, err
	}

	var envelope LastCheckpointEnvelope
	err = impl.request.MakeRequest("GET", url, nil, &envelope)
	return envelope.Data, err
}

func buildLastCheckpointURL(param common.SingleTrackingParam, optionalParams *GetCheckpointParams) (string, *error.AfterShipError) {
	checkpointURL, err := param.BuildTrackingURL("last_checkpoint", "")
	if err != nil {
		return "", err
	}

	if optionalParams != nil {
		checkpointURL, err = common.BuildURLWithQueryString(checkpointURL, optionalParams)
		if err != nil {
			return "", err
		}

	}

	return checkpointURL, nil
}
