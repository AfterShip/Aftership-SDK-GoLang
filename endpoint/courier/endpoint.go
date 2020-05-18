package courier

import (
	"context"

	"github.com/aftership/aftership-sdk-go/v2/error"
	"github.com/aftership/aftership-sdk-go/v2/request"
)

// Endpoint provides the interface for all courier API calls
type Endpoint interface {

	// GetCouriers returns a list of couriers activated at your AfterShip account.
	GetCouriers(ctx context.Context) (List, *error.AfterShipError)

	// GetAllCouriers returns a list of all couriers.
	GetAllCouriers(ctx context.Context) (List, *error.AfterShipError)

	// DetectCouriers returns a list of matched couriers based on tracking number format
	// and selected couriers or a list of couriers.
	DetectCouriers(ctx context.Context, req DetectCourierRequest) (DetectList, *error.AfterShipError)
}

// EndpointImpl is the implementaion of courier endpoint
type EndpointImpl struct {
	request request.APIRequest
}

// NewEndpoint creates a instance of courier endpoint
func NewEndpoint(req request.APIRequest) Endpoint {
	return &EndpointImpl{
		request: req,
	}
}

// GetCouriers returns a list of couriers activated at your AfterShip account.
func (impl *EndpointImpl) GetCouriers(ctx context.Context) (List, *error.AfterShipError) {
	var envelope Envelope
	err := impl.request.MakeRequest(ctx, "GET", "/couriers", nil, &envelope)
	return envelope.Data, err
}

// GetAllCouriers returns a list of all couriers.
func (impl *EndpointImpl) GetAllCouriers(ctx context.Context) (List, *error.AfterShipError) {
	var envelope Envelope
	err := impl.request.MakeRequest(ctx, "GET", "/couriers/all", nil, &envelope)
	return envelope.Data, err
}

// DetectCouriers returns a list of matched couriers based on tracking number format
// and selected couriers or a list of couriers.
func (impl *EndpointImpl) DetectCouriers(ctx context.Context, req DetectCourierRequest) (DetectList, *error.AfterShipError) {
	if req.Tracking.TrackingNumber == "" {
		return DetectList{}, error.NewSdkError(error.ErrorTypeHandlerError, "HandlerError: Invalid TrackingNumber", "")
	}

	var envelope DetectEnvelope
	err := impl.request.MakeRequest(ctx, "POST", "/couriers/detect", req, &envelope)
	return envelope.Data, err
}
