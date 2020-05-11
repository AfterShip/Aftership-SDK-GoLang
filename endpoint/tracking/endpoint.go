package tracking

import (
	"github.com/aftership/aftership-sdk-go/v2/common"
	"github.com/aftership/aftership-sdk-go/v2/error"
	"github.com/aftership/aftership-sdk-go/v2/request"
)

// Endpoint provides the interface for all trackings API calls
type Endpoint interface {
	// CreateTracking creates a tracking.
	CreateTracking(newTracking NewTrackingRequest) (SingleTrackingData, *error.AfterShipError)

	// DeleteTracking deletes a tracking.
	DeleteTracking(param common.SingleTrackingParam) (SingleTrackingData, *error.AfterShipError)

	// GetTrackings gets tracking results of multiple trackings.
	GetTrackings(params MultiTrackingsParams) (MultiTrackingsData, *error.AfterShipError)

	// GetTracking gets tracking results of a single tracking.
	GetTracking(param common.SingleTrackingParam, optionalParams *GetTrackingParams) (SingleTrackingData, *error.AfterShipError)

	// UpdateTracking updates a tracking.
	UpdateTracking(param common.SingleTrackingParam, update UpdateTrackingRequest) (SingleTrackingData, *error.AfterShipError)

	// ReTrack an expired tracking once. Max. 3 times per tracking.
	ReTrack(param common.SingleTrackingParam) (SingleTrackingData, *error.AfterShipError)
}

// EndpointImpl is the implementaion of tracking endpoint
type EndpointImpl struct {
	request request.APIRequest
}

// NewEndpoint creates a instance of tracking endpoint
func NewEndpoint(req request.APIRequest) Endpoint {
	return &EndpointImpl{
		request: req,
	}
}

// CreateTracking creates a new tracking
func (impl *EndpointImpl) CreateTracking(newTracking NewTrackingRequest) (SingleTrackingData, *error.AfterShipError) {
	var envelope SingleTrackingEnvelope
	err := impl.request.MakeRequest("POST", "/trackings", newTracking, &envelope)
	return envelope.Data, err
}

// DeleteTracking deletes a tracking.
func (impl *EndpointImpl) DeleteTracking(param common.SingleTrackingParam) (SingleTrackingData, *error.AfterShipError) {
	url, err := param.BuildTrackingURL("trackings", "")
	if err != nil {
		return SingleTrackingData{}, err
	}

	var envelope SingleTrackingEnvelope
	err = impl.request.MakeRequest("DELETE", url, nil, &envelope)
	return envelope.Data, err
}

// GetTrackings gets tracking results of multiple trackings.
func (impl *EndpointImpl) GetTrackings(params MultiTrackingsParams) (MultiTrackingsData, *error.AfterShipError) {
	url, err := common.BuildURLWithQueryString("/trackings", params)
	if err != nil {
		return MultiTrackingsData{}, err
	}

	var envelope MultiTrackingsEnvelope
	err = impl.request.MakeRequest("GET", url, nil, &envelope)
	return envelope.Data, err
}

// GetTracking gets tracking results of a single tracking.
func (impl *EndpointImpl) GetTracking(param common.SingleTrackingParam, optionalParams *GetTrackingParams) (SingleTrackingData, *error.AfterShipError) {
	url, err := param.BuildTrackingURL("trackings", "")
	if err != nil {
		return SingleTrackingData{}, err
	}

	url, err = common.BuildURLWithQueryString(url, optionalParams)
	if err != nil {
		return SingleTrackingData{}, err
	}

	var envelope SingleTrackingEnvelope
	err = impl.request.MakeRequest("GET", url, nil, &envelope)
	return envelope.Data, err
}

// UpdateTracking updates a tracking.
func (impl *EndpointImpl) UpdateTracking(param common.SingleTrackingParam, update UpdateTrackingRequest) (SingleTrackingData, *error.AfterShipError) {
	url, err := param.BuildTrackingURL("trackings", "")
	if err != nil {
		return SingleTrackingData{}, err
	}

	var envelope SingleTrackingEnvelope
	err = impl.request.MakeRequest("PUT", url, update, &envelope)
	return envelope.Data, err
}

// ReTrack an expired tracking once. Max. 3 times per tracking.
func (impl *EndpointImpl) ReTrack(param common.SingleTrackingParam) (SingleTrackingData, *error.AfterShipError) {
	url, err := param.BuildTrackingURL("trackings", "retrack")
	if err != nil {
		return SingleTrackingData{}, err
	}

	var envelope SingleTrackingEnvelope
	err = impl.request.MakeRequest("POST", url, nil, &envelope)
	return envelope.Data, err
}
