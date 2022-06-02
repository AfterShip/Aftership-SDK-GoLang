package aftership

import (
	"context"
	"errors"
	"net/http"
)

// Courier is the model describing an AfterShip courier
type Courier struct {
	Slug                   string   `json:"slug"`                      // Unique code of courier
	Name                   string   `json:"name"`                      // Name of courier
	Phone                  string   `json:"phone"`                     // Contact phone number of courier
	OtherName              string   `json:"other_name"`                // Other name of courier
	WebURL                 string   `json:"web_url"`                   // Website link of courier
	RequiredFields         []string `json:"required_fields"`           // The extra fields need for tracking, such as `tracking_account_number`, `tracking_postal_code`, `tracking_ship_date`, `tracking_key`, `tracking_destination_country`
	OptionalFields         []string `json:"optional_fields"`           // the extra fields which are optional for tracking. Basically it's the same as required_fields, but the difference is that only some of the tracking numbers require these fields.
	DefaultLanguage        string   `json:"default_language"`          // Default language of tracking results
	SupportedLanguages     []string `json:"supported_languages"`       // Other supported languages
	ServiceFromCountryISO3 []string `json:"service_from_country_iso3"` // Country code (ISO Alpha-3) where the courier provides service
}

// CourierList is the model describing an AfterShip courier list
type CourierList struct {
	Total    int       `json:"total"`    // Total number of couriers supported by AfterShip.
	Couriers []Courier `json:"couriers"` // Array of Courier describes the couriers information.
}

// CourierDetectionParams contains fields required and optional fields for courier detection
type CourierDetectionParams struct {

	// TrackingNumber of a shipment. Mandatory field.
	TrackingNumber string `json:"tracking_number"`

	// Slug If not specified, AfterShip will automatically detect the courier based on the tracking number format and
	// your selected couriers.
	// Use array to input a list of couriers for auto detect.
	Slug []string `json:"slug,omitempty"`

	AdditionalField

	SlugGroup string `json:"slug_group,omitempty"`
}

// GetCouriers returns a list of couriers activated at your AfterShip account.
func (client *Client) GetCouriers(ctx context.Context) (CourierList, error) {
	var courierList CourierList
	err := client.makeRequest(ctx, http.MethodGet, "/couriers", nil, nil, &courierList)
	return courierList, err
}

// GetAllCouriers returns a list of all couriers.
func (client *Client) GetAllCouriers(ctx context.Context) (CourierList, error) {
	var courierList CourierList
	err := client.makeRequest(ctx, http.MethodGet, "/couriers/all", nil, nil, &courierList)
	return courierList, err
}

// detectCourierRequest is a model for detect courier API request
type detectCourierRequest struct {
	Tracking CourierDetectionParams `json:"tracking"`
}

// DetectCouriers returns a list of matched couriers based on tracking number format
// and selected couriers or a list of couriers.
func (client *Client) DetectCouriers(ctx context.Context, params CourierDetectionParams) (CourierList, error) {
	if params.TrackingNumber == "" {
		return CourierList{}, errors.New(errMissingTrackingNumber)
	}

	var courierList CourierList
	err := client.makeRequest(ctx, http.MethodPost, "/couriers/detect", nil,
		&detectCourierRequest{
			Tracking: params,
		}, &courierList)
	return courierList, err
}
