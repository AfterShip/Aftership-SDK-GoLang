package notification

import (
	"github.com/aftership/aftership-sdk-go/v2/common"
	"github.com/aftership/aftership-sdk-go/v2/error"
	"github.com/aftership/aftership-sdk-go/v2/request"
)

// Endpoint provides the interface for all notifications handling API calls
type Endpoint interface {
	// AddNotification Adds notifications to a tracking number.
	AddNotification(param common.SingleTrackingParam, data Data) (Data, *error.AfterShipError)

	// RemoveNotification Removes notifications from a tracking number.
	RemoveNotification(param common.SingleTrackingParam, data Data) (Data, *error.AfterShipError)

	// GetNotification Get contact information for the users to notify when the tracking changes. Please note that only customer receivers will be returned.
	// Any email, sms or webhook that belongs to the Store will not be returned.
	GetNotification(param common.SingleTrackingParam) (Data, *error.AfterShipError)
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
func (impl *EndpointImpl) AddNotification(param common.SingleTrackingParam, data Data) (Data, *error.AfterShipError) {
	url, err := param.BuildTrackingURL("notifications", "add")
	if err != nil {
		return Data{}, err
	}

	var envelope Envelope
	err = impl.request.MakeRequest("POST", url, data, &envelope)
	return envelope.Data, err
}

// RemoveNotification Removes notifications from a tracking number.
func (impl *EndpointImpl) RemoveNotification(param common.SingleTrackingParam, data Data) (Data, *error.AfterShipError) {
	url, err := param.BuildTrackingURL("notifications", "remove")
	if err != nil {
		return Data{}, err
	}

	var envelope Envelope
	err = impl.request.MakeRequest("POST", url, data, &envelope)
	return envelope.Data, err
}

// GetNotification Get contact information for the users to notify when the tracking changes. Please note that only customer receivers will be returned.
// Any email, sms or webhook that belongs to the Store will not be returned.
func (impl *EndpointImpl) GetNotification(param common.SingleTrackingParam) (Data, *error.AfterShipError) {
	url, err := param.BuildTrackingURL("notifications", "")
	if err != nil {
		return Data{}, err
	}

	var envelope Envelope
	err = impl.request.MakeRequest("GET", url, nil, &envelope)
	return envelope.Data, err
}
