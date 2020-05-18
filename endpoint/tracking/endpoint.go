package tracking

import (
	"context"

	"github.com/aftership/aftership-sdk-go/v2/error"
	"github.com/aftership/aftership-sdk-go/v2/request"
)

// Endpoint provides the interface for all trackings API calls
type Endpoint interface {
	// CreateTracking creates a tracking.
	CreateTracking(ctx context.Context, newTracking NewTrackingRequest) (SingleTrackingData, *error.AfterShipError)

	// DeleteTracking deletes a tracking.
	DeleteTracking(ctx context.Context, param SingleTrackingParam) (SingleTrackingData, *error.AfterShipError)

	// GetTrackings gets tracking results of multiple trackings.
	GetTrackings(ctx context.Context, params MultiTrackingsParams) (MultiTrackingsData, *error.AfterShipError)

	// GetTracking gets tracking results of a single tracking.
	GetTracking(ctx context.Context, param SingleTrackingParam, optionalParams *GetTrackingParams) (SingleTrackingData, *error.AfterShipError)

	// UpdateTracking updates a tracking.
	UpdateTracking(ctx context.Context, param SingleTrackingParam, update UpdateTrackingRequest) (SingleTrackingData, *error.AfterShipError)

	// ReTrack an expired tracking once. Max. 3 times per tracking.
	ReTrack(ctx context.Context, param SingleTrackingParam) (SingleTrackingData, *error.AfterShipError)

	// MarkAsCompleted marks a tracking as completed. The tracking won't auto update until retrack it.
	MarkAsCompleted(ctx context.Context, param SingleTrackingParam, reason MarkTrackingAsCompletedRequest) (SingleTrackingData, *error.AfterShipError)
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
func (impl *EndpointImpl) CreateTracking(ctx context.Context, newTracking NewTrackingRequest) (SingleTrackingData, *error.AfterShipError) {
	var envelope SingleTrackingEnvelope
	err := impl.request.MakeRequest(ctx, "POST", "/trackings", newTracking, &envelope)
	return envelope.Data, err
}

// DeleteTracking deletes a tracking.
func (impl *EndpointImpl) DeleteTracking(ctx context.Context, param SingleTrackingParam) (SingleTrackingData, *error.AfterShipError) {
	url, err := param.BuildTrackingURL("trackings", "")
	if err != nil {
		return SingleTrackingData{}, err
	}

	var envelope SingleTrackingEnvelope
	err = impl.request.MakeRequest(ctx, "DELETE", url, nil, &envelope)
	return envelope.Data, err
}

// GetTrackings gets tracking results of multiple trackings.
func (impl *EndpointImpl) GetTrackings(ctx context.Context, params MultiTrackingsParams) (MultiTrackingsData, *error.AfterShipError) {
	url, err := BuildURLWithQueryString("/trackings", params)
	if err != nil {
		return MultiTrackingsData{}, err
	}

	var envelope MultiTrackingsEnvelope
	err = impl.request.MakeRequest(ctx, "GET", url, nil, &envelope)
	return envelope.Data, err
}

// GetTracking gets tracking results of a single tracking.
func (impl *EndpointImpl) GetTracking(ctx context.Context, param SingleTrackingParam, optionalParams *GetTrackingParams) (SingleTrackingData, *error.AfterShipError) {
	url, err := param.BuildTrackingURL("trackings", "")
	if err != nil {
		return SingleTrackingData{}, err
	}

	url, err = BuildURLWithQueryString(url, optionalParams)
	if err != nil {
		return SingleTrackingData{}, err
	}

	var envelope SingleTrackingEnvelope
	err = impl.request.MakeRequest(ctx, "GET", url, nil, &envelope)
	return envelope.Data, err
}

// UpdateTracking updates a tracking.
func (impl *EndpointImpl) UpdateTracking(ctx context.Context, param SingleTrackingParam, update UpdateTrackingRequest) (SingleTrackingData, *error.AfterShipError) {
	url, err := param.BuildTrackingURL("trackings", "")
	if err != nil {
		return SingleTrackingData{}, err
	}

	var envelope SingleTrackingEnvelope
	err = impl.request.MakeRequest(ctx, "PUT", url, update, &envelope)
	return envelope.Data, err
}

// ReTrack an expired tracking once. Max. 3 times per tracking.
func (impl *EndpointImpl) ReTrack(ctx context.Context, param SingleTrackingParam) (SingleTrackingData, *error.AfterShipError) {
	url, err := param.BuildTrackingURL("trackings", "retrack")
	if err != nil {
		return SingleTrackingData{}, err
	}

	var envelope SingleTrackingEnvelope
	err = impl.request.MakeRequest(ctx, "POST", url, nil, &envelope)
	return envelope.Data, err
}

// MarkAsCompleted marks a tracking as completed. The tracking won't auto update until retrack it.
func (impl *EndpointImpl) MarkAsCompleted(ctx context.Context, param SingleTrackingParam, reason MarkTrackingAsCompletedRequest) (SingleTrackingData, *error.AfterShipError) {
	url, err := param.BuildTrackingURL("trackings", "mark-as-completed")
	if err != nil {
		return SingleTrackingData{}, err
	}

	if reason.Reason != "DELIVERED" && reason.Reason != "LOST" && reason.Reason != "RETURNED_TO_SENDER" {
		return SingleTrackingData{}, error.NewSdkError(error.ErrorTypeHandlerError, `Reason must be Ooe of "DELIVERED", "LOST" or "RETURNED_TO_SENDER"`, param)
	}

	var envelope SingleTrackingEnvelope
	err = impl.request.MakeRequest(ctx, "POST", url, reason, &envelope)
	return envelope.Data, err
}
