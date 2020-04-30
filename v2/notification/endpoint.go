package notification

import (
	"github.com/aftership/aftership-sdk-go/v2/error"
	"github.com/aftership/aftership-sdk-go/v2/request"
	"github.com/aftership/aftership-sdk-go/v2/tracking"
)

// Endpoint provides the interface for all notifications handling API calls
type Endpoint interface {
	// AddNotification Adds notifications to a tracking number.
	AddNotification(param tracking.SingleTrackingParam, data Data) (Data, *error.AfterShipError)

	// RemoveNotification Removes notifications from a tracking number.
	RemoveNotification(param tracking.SingleTrackingParam, data Data) (Data, *error.AfterShipError)

	// GetNotification Get contact information for the users to notify when the tracking changes. Please note that only customer receivers will be returned.
	// Any email, sms or webhook that belongs to the Store will not be returned.
	GetNotification(param tracking.SingleTrackingParam) (Data, *error.AfterShipError)
}

// EndpointImpl is the implementaion of notification endpoint
type EndpointImpl struct {
	request request.APIRequest
}

// NewEnpoint creates a instance of notification endpoint
func NewEnpoint(req request.APIRequest) Endpoint {
	return &EndpointImpl{
		request: req,
	}
}

// AddNotification Adds notifications to a tracking number.
func (impl *EndpointImpl) AddNotification(param tracking.SingleTrackingParam, data Data) (Data, *error.AfterShipError) {
	url, err := tracking.BuildTrackingURL(param, "notifications", "add")
	if err != nil {
		return Data{}, err
	}

	var envelope Envelope
	err = impl.request.MakeRequest("POST", url, data, &envelope)
	if err != nil {
		return Data{}, err
	}
	return envelope.Data, nil
}

// RemoveNotification Removes notifications from a tracking number.
func (impl *EndpointImpl) RemoveNotification(param tracking.SingleTrackingParam, data Data) (Data, *error.AfterShipError) {
	url, err := tracking.BuildTrackingURL(param, "notifications", "remove")
	if err != nil {
		return Data{}, err
	}

	var envelope Envelope
	err = impl.request.MakeRequest("POST", url, data, &envelope)
	if err != nil {
		return Data{}, err
	}
	return envelope.Data, nil
}

// GetNotification Get contact information for the users to notify when the tracking changes. Please note that only customer receivers will be returned.
// Any email, sms or webhook that belongs to the Store will not be returned.
func (impl *EndpointImpl) GetNotification(param tracking.SingleTrackingParam) (Data, *error.AfterShipError) {
	url, err := tracking.BuildTrackingURL(param, "notifications", "")
	if err != nil {
		return Data{}, err
	}

	var envelope Envelope
	err = impl.request.MakeRequest("GET", url, nil, &envelope)
	if err != nil {
		return Data{}, err
	}
	return envelope.Data, nil
}
