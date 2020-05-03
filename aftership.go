package aftership

import (
	"github.com/aftership/aftership-sdk-go/v2/checkpoint"
	"github.com/aftership/aftership-sdk-go/v2/common"
	"github.com/aftership/aftership-sdk-go/v2/courier"
	"github.com/aftership/aftership-sdk-go/v2/error"
	"github.com/aftership/aftership-sdk-go/v2/notification"
	"github.com/aftership/aftership-sdk-go/v2/request"
	"github.com/aftership/aftership-sdk-go/v2/response"
	"github.com/aftership/aftership-sdk-go/v2/tracking"
)

// AfterShip is the client for all AfterShip API calls
type AfterShip struct {
	Config         *common.AfterShipConf // The config of AfterShip SDK
	Courier        courier.Endpoint      // The endpoint to get a list of supported couriers.
	Tracking       tracking.Endpoint     // The endpoint to create trackings, update trackings, and get tracking results.
	LastCheckpoint checkpoint.Endpoint   // The endpoint to get tracking information of the last checkpoint of a tracking.
	Notification   notification.Endpoint // The endpoint to get, add or remove contacts (sms or email) to be notified when the status of a tracking has changed.
	RateLimit      *response.RateLimit
}

// NewAfterShip returns the AfterShip client
func NewAfterShip(cfg *common.AfterShipConf) (*AfterShip, *error.AfterShipError) {
	if cfg == nil {
		return nil, error.MakeSdkError(error.ErrorTypeConstructorError, "ConstructorError: config is nil", "")
	}

	if cfg.APIKey == "" {
		return nil, error.MakeSdkError(error.ErrorTypeConstructorError, "ConstructorError: Invalid API key", "")
	}

	if cfg.Endpoint == "" {
		cfg.Endpoint = "https://api.aftership.com/v4"
	}

	if cfg.UserAagentPrefix == "" {
		cfg.UserAagentPrefix = "aftership-sdk-go"
	}

	rateLimit := &response.RateLimit{}
	req := request.NewRequest(cfg, rateLimit)
	return &AfterShip{
		Config:         cfg,
		Courier:        courier.NewEnpoint(req),
		Tracking:       tracking.NewEnpoint(req),
		LastCheckpoint: checkpoint.NewEnpoint(req),
		Notification:   notification.NewEnpoint(req),
		RateLimit:      rateLimit,
	}, nil
}
