package aftership

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// CreateTrackingParams provides parameters for new Tracking API request
type CreateTrackingParams struct {
	TrackingNumber             string            `json:"tracking_number"`                        // Tracking number of a shipment.
	Slug                       []string          `json:"slug,omitempty"`                         // Unique code of each courier. If you do not specify a slug, AfterShip will automatically detect the courier based on the tracking number format and your selected couriers.
	TrackingPostalCode         string            `json:"tracking_postal_code,omitempty"`         // The postal code of receiver's address. Required by some couriers, such as deutsch-post
	TrackingShipDate           string            `json:"tracking_ship_date,omitempty"`           // Shipping date in YYYYMMDD format. Required by some couriers, such as deutsch-post
	TrackingAccountNumber      string            `json:"tracking_account_number,omitempty"`      // Account number of the shipper for a specific courier. Required by some couriers, such as dynamic-logistics
	TrackingKey                string            `json:"tracking_key,omitempty"`                 // Key of the shipment for a specific courier. Required by some couriers, such as sic-teliway
	TrackingOriginCountry      string            `json:"tracking_origin_country,omitempty"`      // Origin Country of the shipment for a specific courier. Required by some couriers, such as dhl
	TrackingDestinationCountry string            `json:"tracking_destination_country,omitempty"` // Destination Country of the shipment for a specific courier. Required by some couriers, such as postnl-3s
	TrackingState              string            `json:"tracking_state,omitempty"`               // Located state of the shipment for a specific courier. Required by some couriers, such as star-track-courier
	Android                    []string          `json:"android,omitempty"`                      // Google cloud message registration IDs to receive the push notifications.
	IOS                        []string          `json:"ios,omitempty"`                          // Apple iOS device IDs to receive the push notifications.
	Emails                     []string          `json:"emails,omitempty"`                       // Email address(es) to receive email notifications.
	SMSes                      []string          `json:"smses,omitempty"`                        // Phone number(s) to receive sms notifications. Enter+ and area code before phone number.
	Title                      string            `json:"title,omitempty"`                        // Title of the tracking. Default value as tracking_number
	CustomerName               string            `json:"customer_name,omitempty"`                // Customer name of the tracking.
	OriginCountryISO3          string            `json:"origin_country_iso3,omitempty"`          // Enter ISO Alpha-3 (three letters) to specify the origin of the shipment (e.g. USA for United States).
	DestinationCountryISO3     string            `json:"destination_country_iso3,omitempty"`     // Enter ISO Alpha-3 (three letters) to specify the destination of the shipment (e.g. USA for United States). If you use postal service to send international shipments, AfterShip will automatically get tracking results at destination courier as well.
	OrderID                    string            `json:"order_id,omitempty"`                     // Text field for order ID
	OrderIDPath                string            `json:"order_id_path,omitempty"`                // Text field for order path
	CustomFields               map[string]string `json:"custom_fields,omitempty"`                // Custom fields that accept a hash with string, boolean or number fields
	Language                   string            `json:"language,omitempty"`                     // Enter ISO 639-1 Language Code to specify the store, customer or order language.
	OrderPromisedDeliveryDate  string            `json:"order_promised_delivery_date,omitempty"` // Promised delivery date of an order inYYYY-MM-DD format.
	DeliveryType               string            `json:"delivery_type,omitempty"`                // Shipment delivery type: pickup_at_store, pickup_at_courier, door_to_door
	PickupLocation             string            `json:"pickup_location,omitempty"`              // Shipment pickup location for receiver
	PickupNote                 string            `json:"pickup_note,omitempty"`                  // Shipment pickup note for receiver
}

type TrackingIdentifier interface {
	URIPath() string
}

type TrackingID string

func (id TrackingID) URIPath() string {
	return "/" + url.PathEscape(string(id))
}

type SlugTrackingNumber struct {
	Slug           string
	TrackingNumber string
}

func (stn SlugTrackingNumber) URIPath() string {
	return fmt.Sprintf("/%s/%s", url.PathEscape(stn.Slug), url.PathEscape(stn.TrackingNumber))
}

// Tracking represents a Tracking returned by the AfterShip API
type Tracking struct {
	ID                            string            `json:"id"`                                         // A unique identifier generated by AfterShip for the tracking.
	CreatedAt                     string            `json:"created_at"`                                 // Date and time of the tracking created.
	UpdatedAt                     string            `json:"updated_at"`                                 // Date and time of the tracking last updated.
	TrackingNumber                string            `json:"tracking_number"`                            // Tracking number of a shipment.
	TrackingPostalCode            string            `json:"tracking_postal_code,omitempty"`             // The postal code of receiver's address. Required by some couriers, such as deutsch-post
	TrackingShipDate              string            `json:"tracking_ship_date,omitempty"`               // Shipping date in YYYYMMDD format. Required by some couriers, such as deutsch-post
	TrackingAccountNumber         string            `json:"tracking_account_number,omitempty"`          // Account number of the shipper for a specific courier. Required by some couriers, such as dynamic-logistics
	TrackingOriginCountry         string            `json:"tracking_origin_country,omitempty"`          // Origin Country of the shipment for a specific courier. Required by some couriers, such as dhl
	TrackingDestinationCountry    string            `json:"tracking_destination_country,omitempty"`     // Destination Country of the shipment for a specific courier. Required by some couriers, such as postnl-3s
	TrackingState                 string            `json:"tracking_state,omitempty"`                   // Located state of the shipment for a specific courier. Required by some couriers, such as star-track-courier
	TrackingKey                   string            `json:"tracking_key,omitempty"`                     // Key of the shipment for a specific courier. Required by some couriers, such as sic-teliway
	Slug                          string            `json:"slug,omitempty"`                             // Unique code of each courier.
	Active                        bool              `json:"active,omitempty"`                           // Whether or not AfterShip will continue tracking the shipments. Value is false when tag (status) is Delivered, Expired, or further updates for 30 days since last update.
	Android                       []string          `json:"android,omitempty"`                          // Google cloud message registration IDs to receive the push notifications.
	CustomFields                  map[string]string `json:"custom_fields,omitempty"`                    // Custom fields that accept a hash with string, boolean or number fields
	DeliveryTime                  int               `json:"delivery_time,omitempty"`                    // Total delivery time in days.
	DestinationCountryISO3        string            `json:"destination_country_iso3,omitempty"`         // Destination country of the tracking. ISO Alpha-3 (three letters). If you use postal service to send international shipments, AfterShip will automatically get tracking results from destination postal service based on destination country.
	CourierDestinationCountryISO3 string            `json:"courier_destination_country_iso3,omitempty"` // Destination country of the tracking detected from the courier. ISO Alpha-3 (three letters). Value will be null if the courier doesn't provide the destination country.
	Emails                        []string          `json:"emails,omitempty"`                           // Email address(es) to receive email notifications.
	ExpectedDelivery              string            `json:"expected_delivery,omitempty"`                // Expected delivery date (nullable). Available format: YYYY-MM-DD, YYYY-MM-DDTHH:MM:SS, or YYYY-MM-DDTHH:MM:SS+TIMEZONE
	IOS                           []string          `json:"ios,omitempty"`                              // Apple iOS device IDs to receive the push notifications.
	OrderID                       string            `json:"order_id,omitempty"`                         // Text field for order ID
	OrderIDPath                   string            `json:"order_id_path,omitempty"`                    // Text field for order path
	OriginCountryISO3             string            `json:"origin_country_iso3,omitempty"`              // Origin country of the tracking. ISO Alpha-3 (three letters).
	UniqueToken                   string            `json:"unique_token,omitempty"`                     // The token to generate the direct tracking link: https://yourusername.aftership.com/unique_token or https://www.aftership.com/unique_token
	ShipmentPackageCount          int               `json:"shipment_package_count,omitempty"`           // Number of packages under the tracking (if any).
	ShipmentType                  string            `json:"shipment_type,omitempty"`                    // Shipment type provided by carrier (if any).
	ShipmentWeight                int               `json:"shipment_weight,omitempty"`                  // Shipment weight provided by carrier (if any)
	ShipmentWeightUnit            string            `json:"shipment_weight_unit,omitempty"`             // Weight unit provided by carrier, either in kg or lb (if any)
	LastUpdatedAt                 *time.Time        `json:"last_updated_at,omitempty"`                  // Date and time the tracking was last updated
	ShipmentPickupDate            *time.Time        `json:"shipment_pickup_date,omitempty"`             // Date and time the tracking was picked up
	ShipmentDeliveryDate          *time.Time        `json:"shipment_delivery_date,omitempty"`           // Date and time the tracking was delivered
	SubscribedSMSes               []string          `json:"subscribed_smses,omitempty"`                 // Phone number(s) subscribed to receive sms notifications.
	SubscribedEmails              []string          `json:"subscribed_emails,omitempty"`                // Email address(es) subscribed to receive email notifications. Comma separated for multiple values
	SignedBy                      string            `json:"signed_by,omitempty"`                        // Signed by information for delivered shipment (if any).
	SMSes                         []string          `json:"smses,omitempty"`                            // Phone number(s) to receive sms notifications. The phone number(s) to receive sms notifications. Phone number should begin with `+` and `Area Code` before phone number. Comma separated for multiple values.
	Source                        string            `json:"source,omitempty"`                           // Source of how this tracking is added.
	Tag                           string            `json:"tag,omitempty"`                              // Current status of tracking.
	Subtag                        string            `json:"subtag,omitempty"`                           // Current subtag of tracking. (See subtag definition)
	SubtagMessage                 string            `json:"subtag_message,omitempty"`                   // Current status of tracking.
	Title                         string            `json:"title,omitempty"`                            // Title of the tracking.
	TrackCount                    int               `json:"tracked_count,omitempty"`                    // Number of attempts AfterShip tracks at courier's system.
	LastMileTrackingSupported     bool              `json:"last_mile_tracking_supported,omitempty"`     // Indicates if the shipment is trackable till the final destination.
	Language                      string            `json:"language,omitempty"`                         // Store, customer, or order language of the tracking.
	ReturnToSender                bool              `json:"return_to_sender,omitempty"`                 // Whether or not the shipment is returned to sender. Value is true when any of its checkpoints has subtagException_010(returning to sender) orException_011(returned to sender). Otherwise value is false
	OrderPromisedDeliveryDate     string            `json:"order_promised_delivery_date,omitempty"`     // Promised delivery date of an order inYYYY-MM-DD format.
	DeliveryType                  string            `json:"delivery_type,omitempty"`                    // Shipment delivery type: pickup_at_store, pickup_at_courier, door_to_door
	PickupLocation                string            `json:"pickup_location,omitempty"`                  // Shipment pickup location for receiver
	PickupNote                    string            `json:"pickup_note,omitempty"`                      // Shipment pickup note for receiver
	CourierTrackingLink           string            `json:"courier_tracking_link,omitempty"`            // Official tracking URL of the courier (if any)
	FirstAttemptedAt              string            `json:"first_attempted_at,omitempty"`               // date and time of the first attempt by the carrier to deliver the package to the addressee. Available format: YYYY-MM-DDTHH:MM:SS, or YYYY-MM-DDTHH:MM:SS+TIMEZONE
	Checkpoints                   []Checkpoint      `json:"checkpoints,omitempty"`                      // Array of Hash describes the checkpoint information.
}

// Checkpoint represents a checkpoint returned by the Aftership API
type Checkpoint struct {
	Slug           string   `json:"slug,omitempty"`
	CreatedAt      string   `json:"created_at,omitempty"`
	CheckpointTime string   `json:"checkpoint_time,omitempty"`
	City           string   `json:"city,omitempty"`
	Coordinates    []string `json:"coordinates,omitempty"`
	CountryISO3    string   `json:"country_iso3,omitempty"`
	CountryName    string   `json:"country_name,omitempty"`
	Message        string   `json:"message,omitempty"`
	State          string   `json:"state,omitempty"`
	Location       string   `json:"location,omitempty"`
	Tag            string   `json:"tag,omitempty"`
	Zip            string   `json:"zip,omitempty"`
}

// GetTrackingParams is the additional parameters in single tracking query
type GetTrackingParams struct {
	// List of fields to include in the response.
	// Use comma for multiple values. Fields to include:
	// tracking_postal_code,tracking_ship_date,tracking_account_number,tracking_key,
	// tracking_origin_country,tracking_destination_country,tracking_state,title,order_id,
	// tag,checkpoints,checkpoint_time, message, country_name
	// Defaults: none, Example: title,order_id
	Fields string `url:"fields,omitempty" json:"fields,omitempty"`

	// Support Chinese to English translation for china-ems  and  china-post  only (Example: en)
	Lang string `url:"lang,omitempty" json:"lang,omitempty"`
}

// UpdateTrackingParams represents an update to Tracking details
type UpdateTrackingParams struct {
	Emails                 []string          `json:"emails,omitempty"`
	SMSes                  []string          `json:"smses,omitempty"`
	Title                  string            `json:"title,omitempty"`
	CustomerName           string            `json:"customer_name,omitempty"`
	DestinationCountryISO3 string            `json:"destination_country_iso3,omitempty"`
	OrderID                string            `json:"order_id,omitempty"`
	OrderIDPath            string            `json:"order_id_path,omitempty"`
	CustomFields           map[string]string `json:"custom_fields,omitempty"`
}

// GetTrackingsParams represents the set of params for get Trackings API
type GetTrackingsParams struct {
	Page         int    `url:"page,omitempty" json:"page,omitempty"`                     // Page to show. (Default: 1)
	Limit        int    `url:"limit,omitempty" json:"limit,omitempty"`                   // Number of trackings each page contain. (Default: 100, Max: 200)
	Keyword      string `url:"keyword,omitempty" json:"keyword,omitempty"`               // Search the content of the tracking record fields:tracking_number, title, order_id, customer_name, custom_fields, order_id, emails, smses
	Slug         string `url:"slug,omitempty" json:"slug,omitempty"`                     // Unique courier code Use comma for multiple values. (Example: dhl,ups,usps)
	DeliveryTime int    `url:"delivery_time,omitempty" json:"delivery_time,omitempty"`   // Total delivery time in days.
	Origin       string `url:"origin,omitempty" json:"origin,omitempty"`                 // Origin country of trackings. Use ISO Alpha-3 (three letters). Use comma for multiple values. (Example: USA,HKG)
	Destination  string `url:"destination,omitempty" json:"destination,omitempty"`       // Destination country of trackings. Use ISO Alpha-3 (three letters). Use comma for multiple values. (Example: USA,HKG)
	Tag          string `url:"tag,omitempty" json:"tag,omitempty"`                       // Current status of tracking. Values include Pending, InfoReceived, InTransit, OutForDelivery, AttemptFail, Delivered, Exception, Expired(See status definition)
	CreatedAtMin string `url:"created_at_min,omitempty" json:"created_at_min,omitempty"` // Start date and time of trackings created. AfterShip only stores data of 90 days. (Defaults: 30 days ago, Example: 2013-03-15T16:41:56+08:00)"
	CreatedAtMax string `url:"created_at_max,omitempty" json:"created_at_max,omitempty"` // End date and time of trackings created. (Defaults: now, Example: 2013-04-15T16:41:56+08:00)"
	Fields       string `url:"fields,omitempty" json:"fields,omitempty"`                 // "List of fields to include in the http. Use comma for multiple values. Fields to include: title, order_id, tag, checkpoints, checkpoint_time, message, country_name. Defaults: none, Example: title,order_id"
	Lang         string `url:"lang,omitempty" json:"lang,omitempty"`                     // "Default: '' / Example: 'en'. Support Chinese to English translation for china-ems and china-post only"
}

// PagedTrackings is a model for data part of the multiple trackings API responses
type PagedTrackings struct {
	Limit     int        `json:"limit"`     // Number of trackings each page contain. (Default: 100)
	Count     int        `json:"count"`     // Total number of matched trackings, max. number is 10,000
	Page      int        `json:"page"`      // Page to show. (Default: 1)
	Trackings []Tracking `json:"trackings"` // Array of Hash describes the tracking information.
}

// trackingWrapper is a model for data part of the single tracking API responses
type trackingWrapper struct {
	Tracking Tracking `json:"tracking"`
}

// SingleTrackingOptionalParams is the optional parameters in single tracking query
type SingleTrackingOptionalParams struct {
	TrackingPostalCode         string `url:"tracking_postal_code,omitempty" json:"tracking_postal_code,omitempty"`                 // The postal code of receiver's address. Required by some couriers, such asdeutsch-post
	TrackingShipDate           string `url:"tracking_ship_date,omitempty" json:"tracking_ship_date,omitempty"`                     // Shipping date in YYYYMMDD format. Required by some couriers, such asdeutsch-post
	TrackingDestinationCountry string `url:"tracking_destination_country,omitempty" json:"tracking_destination_country,omitempty"` // Destination Country of the shipment for a specific courier. Required by some couriers, such aspostnl-3s
	TrackingAccountNumber      string `url:"tracking_account_number,omitempty" json:"tracking_account_number,omitempty"`           // Account number of the shipper for a specific courier. Required by some couriers, such asdynamic-logistics
	TrackingKey                string `url:"tracking_key,omitempty" json:"tracking_key,omitempty"`                                 // Key of the shipment for a specific courier. Required by some couriers, such assic-teliway
	TrackingOriginCountry      string `url:"tracking_origin_country,omitempty" json:"tracking_origin_country,omitempty"`           // Origin Country of the shipment for a specific courier. Required by some couriers, such asdhl
	TrackingState              string `url:"tracking_state,omitempty" json:"tracking_state,omitempty"`                             // Located state of the shipment for a specific courier. Required by some couriers, such asstar-track-courier
}

type CompletedStatus string

// CompletedStatus DELIVERED
const CompletedStatusDelivered CompletedStatus = "DELIVERED"

// CompletedStatus LOST
const CompletedStatusLost CompletedStatus = "LOST"

// CompletedStatus RETURNED_TO_SENDER
const CompletedStatusReturnedToSender CompletedStatus = "RETURNED_TO_SENDER"

// TrackingsEndpoint provides the interface for all trackings API calls
type TrackingsEndpoint interface {
	// CreateTracking creates a tracking.
	CreateTracking(ctx context.Context, params CreateTrackingParams) (Tracking, error)

	// DeleteTracking deletes a tracking.
	DeleteTracking(ctx context.Context, id TrackingIdentifier) (Tracking, error)

	// GetTrackings gets tracking results of multiple trackings.
	GetTrackings(ctx context.Context, params GetTrackingsParams) (PagedTrackings, error)

	// GetTracking gets tracking results of a single tracking.
	GetTracking(ctx context.Context, id TrackingIdentifier, params GetTrackingParams) (Tracking, error)

	// UpdateTracking updates a tracking.
	UpdateTracking(ctx context.Context, id TrackingIdentifier, update UpdateTrackingParams) (Tracking, error)

	// ReTrack an expired tracking once. Max 3 times per tracking.
	ReTrack(ctx context.Context, id TrackingIdentifier) (Tracking, error)

	// MarkAsCompleted marks a tracking as completed. The tracking won't auto update until retrack it.
	MarkAsCompleted(ctx context.Context, id TrackingIdentifier, status CompletedStatus) (Tracking, error)
}

// trackingsEndpointImpl is the implementation of tracking endpoint
type trackingsEndpointImpl struct {
	request requestHelper
}

// newTrackingsEndpoint creates a instance of tracking endpoint
func newTrackingsEndpoint(req requestHelper) TrackingsEndpoint {
	return &trackingsEndpointImpl{
		request: req,
	}
}

// createTrackingRequest is a model for create tracking API request
type createTrackingRequest struct {
	Tracking CreateTrackingParams `json:"tracking"`
}

// CreateTracking creates a new tracking
func (impl *trackingsEndpointImpl) CreateTracking(ctx context.Context, params CreateTrackingParams) (Tracking, error) {
	var trackingWrapper trackingWrapper
	err := impl.request.makeRequest(ctx, http.MethodPost, "/trackings", nil,
		&createTrackingRequest{Tracking: params}, &trackingWrapper)
	return trackingWrapper.Tracking, err
}

// DeleteTracking deletes a tracking.
func (impl *trackingsEndpointImpl) DeleteTracking(ctx context.Context, id TrackingIdentifier) (Tracking, error) {
	uriPath := fmt.Sprintf("/trackings%s", id.URIPath())
	var trackingWrapper trackingWrapper
	err := impl.request.makeRequest(ctx, http.MethodDelete, uriPath, nil, nil, &trackingWrapper)
	return trackingWrapper.Tracking, err
}

// GetTrackings gets tracking results of multiple trackings.
func (impl *trackingsEndpointImpl) GetTrackings(ctx context.Context, params GetTrackingsParams) (PagedTrackings, error) {
	var pagedTrackings PagedTrackings
	err := impl.request.makeRequest(ctx, http.MethodGet, "/trackings", params, nil, &pagedTrackings)
	return pagedTrackings, err
}

// GetTracking gets tracking results of a single tracking.
func (impl *trackingsEndpointImpl) GetTracking(ctx context.Context, id TrackingIdentifier, optionalParams GetTrackingParams) (Tracking, error) {
	uriPath := fmt.Sprintf("/trackings%s", id.URIPath())
	var trackingWrapper trackingWrapper
	err := impl.request.makeRequest(ctx, http.MethodGet, uriPath, optionalParams, nil, &trackingWrapper)
	return trackingWrapper.Tracking, err
}

// UpdateTrackingRequest is a model for update tracking API request
type UpdateTrackingRequest struct {
	Tracking UpdateTrackingParams `json:"tracking"`
}

// UpdateTrackingParams updates a tracking.
func (impl *trackingsEndpointImpl) UpdateTracking(ctx context.Context, id TrackingIdentifier, update UpdateTrackingParams) (Tracking, error) {
	uriPath := fmt.Sprintf("/trackings%s", id.URIPath())
	var trackingWrapper trackingWrapper
	err := impl.request.makeRequest(ctx, http.MethodPut, uriPath, nil, update, &trackingWrapper)
	return trackingWrapper.Tracking, err
}

// ReTrack an expired tracking once. Max. 3 times per tracking.
func (impl *trackingsEndpointImpl) ReTrack(ctx context.Context, id TrackingIdentifier) (Tracking, error) {
	uriPath := fmt.Sprintf("/trackings%s/retrack", id.URIPath())
	var trackingWrapper trackingWrapper
	err := impl.request.makeRequest(ctx, http.MethodPost, uriPath, nil, nil, &trackingWrapper)
	return trackingWrapper.Tracking, err
}

// markAsCompletedRequest is a model for update tracking as completed API request
type markAsCompletedRequest struct {
	Reason string `json:"reason"` // One of "DELIVERED", "LOST" or "RETURNED_TO_SENDER".
}

// MarkAsCompleted marks a tracking as completed. The tracking won't auto update until retrack it.
func (impl *trackingsEndpointImpl) MarkAsCompleted(ctx context.Context, id TrackingIdentifier, status CompletedStatus) (Tracking, error) {
	uriPath := fmt.Sprintf("/trackings%s/mark-as-completed", id.URIPath())
	var trackingWrapper trackingWrapper
	err := impl.request.makeRequest(ctx, http.MethodPost, uriPath,
		nil, &markAsCompletedRequest{Reason: string(status)}, &trackingWrapper)
	return trackingWrapper.Tracking, err
}
