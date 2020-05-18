package notification

import (
	"context"

	"github.com/aftership/aftership-sdk-go/v2/endpoint/tracking"
	"github.com/aftership/aftership-sdk-go/v2/error"
	"github.com/aftership/aftership-sdk-go/v2/request"
)

// Endpoint provides the interface for all notifications handling API calls
type Endpoint interface {
	// AddNotification adds notifications to a single tracking.
	AddNotification(ctx context.Context, param tracking.SingleTrackingParam, data Data) (Data, *error.AfterShipError)

	// RemoveNotification removes notifications from a single tracking.
	RemoveNotification(ctx context.Context, param tracking.SingleTrackingParam, data Data) (Data, *error.AfterShipError)

	// GetNotification gets contact information for the users to notify when the tracking changes. Please note that only customer receivers will be returned.
	// Any email, sms or webhook that belongs to the Store will not be returned.
	GetNotification(ctx context.Context, param tracking.SingleTrackingParam) (Data, *error.AfterShipError)
}

// EndpointImpl is the implementaion of notification endpoint
type EndpointImpl struct {
	request request.APIRequest
}

// NewEndpoint creates a instance of notification endpoint
func NewEndpoint(req request.APIRequest) Endpoint {
	return &EndpointImpl{
		request: req,
	}
}

// AddNotification adds notifications to a single tracking.
func (impl *EndpointImpl) AddNotification(ctx context.Context, param tracking.SingleTrackingParam, data Data) (Data, *error.AfterShipError) {
	url, err := param.BuildTrackingURL("notifications", "add")
	if err != nil {
		return Data{}, err
	}

	var envelope Envelope
	err = impl.request.MakeRequest(ctx, "POST", url, data, &envelope)
	return envelope.Data, err
}

// RemoveNotification removes notifications from a single tracking.
func (impl *EndpointImpl) RemoveNotification(ctx context.Context, param tracking.SingleTrackingParam, data Data) (Data, *error.AfterShipError) {
	url, err := param.BuildTrackingURL("notifications", "remove")
	if err != nil {
		return Data{}, err
	}

	var envelope Envelope
	err = impl.request.MakeRequest(ctx, "POST", url, data, &envelope)
	return envelope.Data, err
}

// GetNotification gets contact information for the users to notify when the tracking changes. Please note that only customer receivers will be returned.
// Any email, sms or webhook that belongs to the Store will not be returned.
func (impl *EndpointImpl) GetNotification(ctx context.Context, param tracking.SingleTrackingParam) (Data, *error.AfterShipError) {
	url, err := param.BuildTrackingURL("notifications", "")
	if err != nil {
		return Data{}, err
	}

	var envelope Envelope
	err = impl.request.MakeRequest(ctx, "GET", url, nil, &envelope)
	return envelope.Data, err
}
