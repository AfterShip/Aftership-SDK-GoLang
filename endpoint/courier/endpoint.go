package courier

import (
	"github.com/aftership/aftership-sdk-go/v2/error"
	"github.com/aftership/aftership-sdk-go/v2/request"
)

// Endpoint provides the interface for all courier API calls
type Endpoint interface {

	// GetCouriers returns a list of couriers activated at your AfterShip account.
	GetCouriers() (List, *error.AfterShipError)

	// GetAllCouriers returns a list of all couriers.
	GetAllCouriers() (List, *error.AfterShipError)

	// DetectCouriers returns a list of matched couriers based on tracking number format
	// and selected couriers or a list of couriers.
	DetectCouriers(req DetectCourierRequest) (DetectList, *error.AfterShipError)
}

// EndpointImpl is the implementaion of courier endpoint
type EndpointImpl struct {
	request request.APIRequest
}

// NewEnpoint creates a instance of courier endpoint
func NewEnpoint(req request.APIRequest) Endpoint {
	return &EndpointImpl{
		request: req,
	}
}

// GetCouriers returns a list of couriers activated at your AfterShip account.
func (impl *EndpointImpl) GetCouriers() (List, *error.AfterShipError) {
	var envelope Envelope
	err := impl.request.MakeRequest("GET", "/couriers", nil, &envelope)
	return envelope.Data, err
}

// GetAllCouriers returns a list of all couriers.
func (impl *EndpointImpl) GetAllCouriers() (List, *error.AfterShipError) {
	var envelope Envelope
	err := impl.request.MakeRequest("GET", "/couriers/all", nil, &envelope)
	return envelope.Data, err
}

// DetectCouriers returns a list of matched couriers based on tracking number format
// and selected couriers or a list of couriers.
func (impl *EndpointImpl) DetectCouriers(req DetectCourierRequest) (DetectList, *error.AfterShipError) {
	if req.Tracking.TrackingNumber == "" {
		return DetectList{}, error.MakeSdkError(error.ErrorTypeHandlerError, "HandlerError: Invalid TrackingNumber", "")
	}

	var envelope DetectEnvelope
	err := impl.request.MakeRequest("POST", "/couriers/detect", req, &envelope)
	return envelope.Data, err
}
