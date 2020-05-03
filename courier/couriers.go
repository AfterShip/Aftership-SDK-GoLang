package courier

import (
	"github.com/aftership/aftership-sdk-go/v2/response"
	"github.com/aftership/aftership-sdk-go/v2/tracking"
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

// List is the model describing an AfterShip courier list
type List struct {
	Total    int       `json:"total"`    // Total number of couriers supported by AfterShip.
	Couriers []Courier `json:"couriers"` // Array of Hash describes the couriers information.
}

// DetectList is the model describing an AfterShip couriers derect list
type DetectList struct {
	Total    int               `json:"total"`    // Total number of matched couriers
	Tracking tracking.Tracking `json:"tracking"` // Hash describes the tracking information.
	Couriers []Courier         `json:"couriers"` // A list of matched couriers based on tracking number format.
}

// Envelope is the message envelope for the courier API responses
type Envelope struct {
	Meta response.Meta `json:"meta"`
	Data List          `json:"data"`
}

// DetectEnvelope is the message envelope for the couriers detect API responses
type DetectEnvelope struct {
	Meta response.Meta `json:"meta"`
	Data DetectList    `json:"data"`
}

// GetMeta returns the response meta
func (e *Envelope) GetMeta() response.Meta {
	return e.Meta
}

// GetMeta returns the response meta
func (e *DetectEnvelope) GetMeta() response.Meta {
	return e.Meta
}

// DetectParam contains fields required and optional fields for
// courier detection
type DetectParam struct {

	// TrackingNumber of a shipment. Mandatory field.
	TrackingNumber string `json:"tracking_number"`

	// TrackingPostalCode is the postal code of receiver's address.
	// Required by some couriers, such as postnl
	TrackingPostalCode string `json:"tracking_postal_code,omitempty"`

	// TrackingShipDate in YYYYMMDD format. Required by some couriers, such as deutsch-post
	TrackingShipDate string `json:"tracking_ship_date,omitempty"`

	// TrackingAccountNumber of the shipper for a specific courier. Required by some couriers, such as dynamic-logistics
	TrackingAccountNumber string `json:"tracking_account_number,omitempty"`

	// TrackingKey of the shipment for a specific courier. Required by some couriers, such as sic-teliway
	TrackingKey string `json:"tracking_key,omitempty"`

	// TrackingDestinationCountry of the shipment for a specific courier. Required by some couriers, such as postnl-3s
	TrackingDestinationCountry string `json:"tracking_destination_country,omitempty"`

	// Slug If not specified, Aftership will automatically detect the courier based on the tracking number format and
	// your selected couriers.
	// Use array to input a list of couriers for auto detect.
	Slug []string `json:"slug,omitempty"`
}

// DetectCourierRequest is a model for detect courier API request
type DetectCourierRequest struct {
	Tracking DetectParam `json:"tracking"`
}
