package aftership

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/pkg/errors"
)

// CreateTrackingParams provides parameters for new Tracking API request
type CreateTrackingParams struct {
	/**
	 * Tracking number of a shipment.
	 */
	TrackingNumber string `json:"tracking_number"`

	/**
	 * Unique code of each courier. If you do not specify a slug,
	 * AfterShip will automatically detect the courier based on the tracking number format and your selected couriers.
	 */
	Slug string `json:"slug,omitempty"`

	/**
	 * Title of the tracking. Default value as tracking_number
	 */
	Title string `json:"title,omitempty"`

	/**
	 * Text field for order ID
	 */
	OrderID string `json:"order_id,omitempty"`

	/**
	 * Text field for order path
	 */
	OrderIDPath string `json:"order_id_path,omitempty"`

	/**
	 * Custom fields that accept a hash with string, boolean or number fields
	 */
	CustomFields map[string]string `json:"custom_fields,omitempty"`

	/**
	 * Enter ISO 639-1 Language Code to specify the store, customer or order language.
	 */
	Language string `json:"language,omitempty"`

	/**
	 * Promised delivery date of an order inYYYY-MM-DD format.
	 */
	OrderPromisedDeliveryDate string `json:"order_promised_delivery_date,omitempty"`

	/**
	 * Shipment delivery type: pickup_at_store, pickup_at_courier, door_to_door
	 */
	DeliveryType string `json:"delivery_type,omitempty"`

	/**
	 * Shipment pickup location for receiver
	 */
	PickupLocation string `json:"pickup_location,omitempty"`

	/**
	 * Shipment pickup note for receiver
	 */
	PickupNote string `json:"pickup_note,omitempty"`

	/**
	 * Account number of the shipper for a specific courier. Required by some couriers, such as dynamic-logistics
	 */
	TrackingAccountNumber string `json:"tracking_account_number,omitempty"`

	/**
	 * Origin Country of the shipment for a specific courier. Required by some couriers, such as dhl
	 */
	TrackingOriginCountry string `json:"tracking_origin_country,omitempty"`

	/**
	 * Destination Country of the shipment for a specific courier. Required by some couriers, such as postnl-3s
	 */
	TrackingDestinationCountry string `json:"tracking_destination_country,omitempty"`

	/**
	 * Key of the shipment for a specific courier. Required by some couriers, such as sic-teliway
	 */
	TrackingKey string `json:"tracking_key,omitempty"`

	/**
	 * The postal code of receiver's address. Required by some couriers, such as deutsch-post
	 */
	TrackingPostalCode string `json:"tracking_postal_code,omitempty"`

	/**
	 * Shipping date in YYYYMMDD format. Required by some couriers, such as deutsch-post
	 */
	TrackingShipDate string `json:"tracking_ship_date,omitempty"`

	/**
	 * Located state of the shipment for a specific courier. Required by some couriers, such as star-track-courier
	 */
	TrackingState string `json:"tracking_state,omitempty"`

	/**
	 * Apple iOS device IDs to receive the push notifications.
	 */
	IOS []string `json:"ios,omitempty"`

	/**
	 * Google cloud message registration IDs to receive the push notifications.
	 */
	Android []string `json:"android,omitempty"`

	/**
	 * Email address(es) to receive email notifications.
	 */
	Emails []string `json:"emails,omitempty"`

	/**
	 * Phone number(s) to receive sms notifications. Enter+ and area code before phone number.
	 */
	SMSes []string `json:"smses,omitempty"`

	/**
	 * Customer name of the tracking.
	 */
	CustomerName string `json:"customer_name,omitempty"`

	/**
	 * Enter ISO Alpha-3 (three letters) to specify the origin of the shipment (e.g. USA for United States).
	 */
	OriginCountryISO3 string `json:"origin_country_iso3,omitempty"`

	/**
	 * Enter ISO Alpha-3 (three letters) to specify the destination of the shipment (e.g. USA for United States). If you use postal service to send international shipments, AfterShip will automatically get tracking results at destination courier as well.
	 */
	DestinationCountryISO3 string `json:"destination_country_iso3,omitempty"`

	/**
	 * Text field for the note
	 */
	Note string `json:"note,omitempty"`

	/**
	 * Slug group is a group of slugs which belong to same courier. For example, when you inpit "fedex-group" as slug_group, AfterShip will detect the tracking with "fedex-uk", "fedex-fims", and other slugs which belong to "fedex". It cannot be used with slug at the same time.
	 */
	SlugGroup string `json:"slug_group,omitempty"`

	/**
	 * Date and time of the order created
	 */
	OrderDate string `json:"order_date,omitempty"`

	/**
	 * Text field for order number
	 */
	OrderNumber string `json:"order_number,omitempty"`
}

// TrackingIdentifier is an identifier for a single tracking
type TrackingIdentifier interface {
	URIPath() (string, error)
}

// TrackingID is a unique identifier generated by AfterShip for the tracking.
type TrackingID string

// URIPath returns the URL path of TrackingID
func (id TrackingID) URIPath() (string, error) {
	if id == "" {
		return "", errors.New(errMissingTrackingID)
	}
	return "/" + url.PathEscape(string(id)), nil
}

// SlugTrackingNumber is a unique identifier for a single tracking by slug and tracking number
type SlugTrackingNumber struct {
	Slug           string
	TrackingNumber string
}

// URIPath returns the URL path of SlugTrackingNumber
func (stn SlugTrackingNumber) URIPath() (string, error) {
	if stn.Slug == "" || stn.TrackingNumber == "" {
		return "", errors.New(errMissingSlugOrTrackingNumber)
	}
	return fmt.Sprintf("/%s/%s", url.PathEscape(stn.Slug), url.PathEscape(stn.TrackingNumber)), nil
}

// Tracking represents a Tracking returned by the AfterShip API
type Tracking struct {
	/**
	 * A unique identifier generated by AfterShip for the tracking.
	 */
	ID string `json:"id"`

	/**
	 * Date and time of the tracking created.
	 */
	CreatedAt *time.Time `json:"created_at"`

	/**
	 * Date and time of the tracking last updated.
	 */
	UpdatedAt *time.Time `json:"updated_at"`

	/**
	 * Date and time the tracking was last updated
	 */
	LastUpdatedAt *time.Time `json:"last_updated_at,omitempty"`

	/**
	 * Tracking number of a shipment.
	 */
	TrackingNumber string `json:"tracking_number"`

	/**
	 * Unique code of each courier.
	 */
	Slug string `json:"slug,omitempty"`

	/**
	 * Whether or not AfterShip will continue tracking the shipments. Value is false when tag (status) is Delivered, Expired, or further updates for 30 days since last update.
	 */
	Active bool `json:"active,omitempty"`

	/**
	 * Google cloud message registration IDs to receive the push notifications.
	 */
	Android []string `json:"android,omitempty"`

	/**
	 * Custom fields that accept a hash with string, boolean or number fields
	 */
	CustomFields map[string]string `json:"custom_fields,omitempty"`

	/**
	 * Customer name of the tracking.
	 */
	CustomerName string `json:"custom_name,omitempty"`

	/**
	 * Total delivery time in days.
	 */
	DeliveryTime int `json:"delivery_time,omitempty"`

	/**
	 * Destination country of the tracking. ISO Alpha-3 (three letters).
	 * If you use postal service to send international shipments, AfterShip will automatically get tracking results from destination postal service based on destination country.
	 */
	DestinationCountryISO3 string `json:"destination_country_iso3,omitempty"`

	/**
	 * Shipping address that the shipment is shipping to.
	 */
	DestinationRawLocation string `json:"destination_raw_location,omitempty"`

	/**
	 * Destination country of the tracking detected from the courier. ISO Alpha-3 (three letters). Value will be null if the courier doesn't provide the destination country.
	 */
	CourierDestinationCountryISO3 string `json:"courier_destination_country_iso3,omitempty"`

	/**
	 * Email address(es) to receive email notifications.
	 */
	Emails []string `json:"emails,omitempty"`

	/**
	 * Expected delivery date (nullable). Available format: YYYY-MM-DD, YYYY-MM-DDTHH:MM:SS, or YYYY-MM-DDTHH:MM:SS+TIMEZONE
	 */
	ExpectedDelivery string `json:"expected_delivery,omitempty"`

	/**
	 * Apple iOS device IDs to receive the push notifications.
	 */
	IOS []string `json:"ios,omitempty"`

	/**
	 * Text field for the note.
	 */
	Note string `json:"note,omitempty"`

	/**
	 * Text field for order ID
	 */
	OrderID string `json:"order_id,omitempty"`

	/**
	 * Text field for order path
	 */
	OrderIDPath string `json:"order_id_path,omitempty"`

	/**
	 * Date and time of the order created
	 */
	OrderDate string `json:"order_date,omitempty"`

	/**
	 * Origin country of the tracking. ISO Alpha-3 (three letters).
	 */
	OriginCountryISO3 string `json:"origin_country_iso3,omitempty"`

	/**
	 * Number of packages under the tracking (if any).
	 */
	ShipmentPackageCount int `json:"shipment_package_count,omitempty"`

	/**
	 * Date and time the tracking was picked up
	 */
	ShipmentPickupDate string `json:"shipment_pickup_date,omitempty"`

	/**
	 * Date and time the tracking was delivered
	 */
	ShipmentDeliveryDate string `json:"shipment_delivery_date,omitempty"`

	/**
	 * Shipment type provided by carrier (if any).
	 */
	ShipmentType string `json:"shipment_type,omitempty"`

	/**
	 * Shipment weight provided by carrier (if any)
	 */
	ShipmentWeight float64 `json:"shipment_weight,omitempty"`

	/**
	 * Weight unit provided by carrier, either in kg or lb (if any)
	 */
	ShipmentWeightUnit string `json:"shipment_weight_unit,omitempty"`

	/**
	 * Signed by information for delivered shipment (if any).
	 */
	SignedBy string `json:"signed_by,omitempty"`

	/**
	 * Phone number(s) to receive sms notifications. The phone number(s) to receive sms notifications. Phone number should begin with `+` and `Area Code` before phone number. Comma separated for multiple values.
	 */
	SMSes []string `json:"smses,omitempty"`

	/**
	 * Source of how this tracking is added.
	 */
	Source string `json:"source,omitempty"`

	/**
	 * Current status of tracking.
	 */
	Tag string `json:"tag,omitempty"`

	/**
	 * Current subtag of tracking. (See subtag definition)
	 */
	Subtag string `json:"subtag,omitempty"`

	/**
	 * Current status of tracking.
	 */
	SubtagMessage string `json:"subtag_message,omitempty"`

	/**
	 * Title of the tracking.
	 */
	Title string `json:"title,omitempty"`

	/**
	 * Number of attempts AfterShip tracks at courier's system.
	 */
	TrackedCount int `json:"tracked_count,omitempty"`

	/**
	 * Indicates if the shipment is trackable till the final destination.
	 */
	LastMileTrackingSupported bool `json:"last_mile_tracking_supported,omitempty"`

	/**
	 * Store, customer, or order language of the tracking.
	 */
	Language string `json:"language,omitempty"`

	/**
	 * The token to generate the direct tracking link: https://yourusername.aftership.com/unique_token or https://www.aftership.com/unique_token
	 */
	UniqueToken string `json:"unique_token,omitempty"`

	/**
	 * Array of Hash describes the checkpoint information.
	 */
	Checkpoints []Checkpoint `json:"checkpoints,omitempty"`

	/**
	 * Phone number(s) subscribed to receive sms notifications.
	 */
	SubscribedSMSes []string `json:"subscribed_smses,omitempty"`

	/**
	 * Email address(es) subscribed to receive email notifications. Comma separated for multiple values
	 */
	SubscribedEmails []string `json:"subscribed_emails,omitempty"`

	/**
	 * Whether or not the shipment is returned to sender. Value is true when any of its checkpoints has subtagException_010(returning to sender) orException_011(returned to sender). Otherwise value is false
	 */
	ReturnToSender bool `json:"return_to_sender,omitempty"`

	/**
	 * Promised delivery date of an order inYYYY-MM-DD format.
	 */
	OrderPromisedDeliveryDate string `json:"order_promised_delivery_date,omitempty"`

	/**
	 * Shipment delivery type: pickup_at_store, pickup_at_courier, door_to_door
	 */
	DeliveryType string `json:"delivery_type,omitempty"`

	/**
	 * Shipment pickup location for receiver
	 */
	PickupLocation string `json:"pickup_location,omitempty"`

	/**
	 * Shipment pickup note for receiver
	 */
	PickupNote string `json:"pickup_note,omitempty"`

	/**
	 * Official tracking URL of the courier (if any)
	 */
	CourierTrackingLink string `json:"courier_tracking_link,omitempty"`

	/**
	 * date and time of the first attempt by the carrier to deliver the package to the addressee. Available format: YYYY-MM-DDTHH:MM:SS, or YYYY-MM-DDTHH:MM:SS+TIMEZONE
	 */
	FirstAttemptedAt string `json:"first_attempted_at,omitempty"`

	/**
	 * Delivery instructions (delivery date or address) can be modified by visiting the link if supported by a carrier.
	 */
	CourierRedirectLink string `json:"courier_redirect_link,omitempty"`

	/**
	 * Account number of the shipper for a specific courier. Required by some couriers, such as dynamic-logistics
	 */
	TrackingAccountNumber string `json:"tracking_account_number,omitempty"`

	/**
	 * Origin Country of the shipment for a specific courier. Required by some couriers, such as dhl
	 */
	TrackingOriginCountry string `json:"tracking_origin_country,omitempty"`

	/**
	 * Destination Country of the shipment for a specific courier. Required by some couriers, such as postnl-3s
	 */
	TrackingDestinationCountry string `json:"tracking_destination_country,omitempty"`

	/**
	 * Key of the shipment for a specific courier. Required by some couriers, such as sic-teliway
	 */
	TrackingKey string `json:"tracking_key,omitempty"`

	/**
	 * The postal code of receiver's address. Required by some couriers, such as deutsch-post
	 */
	TrackingPostalCode string `json:"tracking_postal_code,omitempty"`

	/**
	 * Shipping date in YYYYMMDD format. Required by some couriers, such as deutsch-post
	 */
	TrackingShipDate string `json:"tracking_ship_date,omitempty"`

	/**
	 * Located state of the shipment for a specific courier. Required by some couriers, such as star-track-courier
	 */
	TrackingState string `json:"tracking_state,omitempty"`

	/**
	 * Whether the tracking is delivered on time or not.
	 */
	OnTimeStatus string `json:"on_time_status,omitempty"`

	/**
	 * The difference days of the on time.
	 */
	OnTimeDifference int `json:"on_time_difference,omitempty"`

	/**
	 * The tags of the order.
	 */
	OrderTags []string `json:"order_tags,omitempty"`

	/**
	 * Estimated delivery time of the shipment provided by AfterShip, indicate when the shipment should arrive.
	 */
	EstimatedDeliveryDate EstimatedDeliveryDate `json:"aftership_estimated_delivery_date,omitempty"`

	/**
	 * Text field for order number
	 */
	OrderNumber string `json:"order_number,omitempty"`

	/**
	 * The latest estimated delivery date.
	 * May come from the carrier, AfterShip AI, or based on your custom settings.
	 * This can appear in 1 of 3 formats based on the data received.
	 *  1. Date only: `YYYY-MM-DD`
	 *  2. Date and time: `YYYY-MM-DDTHH:mm:ss`
	 *  3. Date, time, and time zone: `YYYY-MM-DDTHH:mm:ssZ`
	 */
	LatestEstimatedDelivery LatestEstimatedDelivery `json:"latest_estimated_delivery,omitempty"`
}

// LatestEstimatedDelivery represents a latest_estimated_delivery returned by the Aftership API
type LatestEstimatedDelivery struct {
	Type        string `json:"type,omitempty"`         // The format of the EDD. Either a single date or a date range.
	Source      string `json:"source,omitempty"`       // The source of the EDD. Either the carrier, AfterShip AI, or based on your custom EDD settings.
	Datetime    string `json:"datetime,omitempty"`     // The latest EDD time.
	DatetimeMin string `json:"datetime_min,omitempty"` // For a date range EDD format, the date and time for the lower end of the range.
	DatetimeMax string `json:"datetime_max,omitempty"` // For a date range EDD format, the date and time for the upper end of the range.
}

// Checkpoint represents a checkpoint returned by the Aftership API
type Checkpoint struct {
	Slug           string     `json:"slug,omitempty"`
	CreatedAt      *time.Time `json:"created_at,omitempty"`
	CheckpointTime string     `json:"checkpoint_time,omitempty"`
	City           string     `json:"city,omitempty"`
	Coordinates    []string   `json:"coordinates,omitempty"`
	CountryISO3    string     `json:"country_iso3,omitempty"`
	CountryName    string     `json:"country_name,omitempty"`
	Message        string     `json:"message,omitempty"`
	State          string     `json:"state,omitempty"`
	Location       string     `json:"location,omitempty"`
	Tag            string     `json:"tag,omitempty"`
	Subtag         string     `json:"subtag,omitempty"`
	SubtagMessage  string     `json:"subtag_message,omitempty"`
	Zip            string     `json:"zip,omitempty"`
	RawTag         string     `json:"raw_tag,omitempty"`
}

// EstimatedDeliveryDate represents a aftership_estimated_delivery_date returned by the Aftership API
type EstimatedDeliveryDate struct {
	/**
	 * The estimated arrival date of the shipment.
	 */
	EstimatedDeliveryDate string `json:"estimated_delivery_date,omitempty"`

	/**
	 * The reliability of the estimated delivery date based on the trend of the transit time for the similar delivery route and the carrier's delivery performance range from 0.0 to 1.0 (Beta feature).
	 */
	ConfidenceScore float64 `json:"confidence_score,omitempty"`

	/**
	 * Earliest estimated delivery date of the shipment.
	 */
	EstimatedDeliveryDateMin string `json:"estimated_delivery_date_min,omitempty"`

	/**
	 * Latest estimated delivery date of the shipment.
	 */
	EstimatedDeliveryDateMax string `json:"estimated_delivery_date_max,omitempty"`
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
	OrderNumber            string            `json:"order_number,omitempty"`
	OrderDate              string            `json:"order_date,omitempty"`
	CustomFields           map[string]string `json:"custom_fields,omitempty"`
	ShipmentType           string            `json:"shipment_type,omitempty"` // The carrier’s shipment type. When you input this field, AfterShip will not get updates from the carrier.
}

// GetTrackingsParams represents the set of params for get Trackings API
type GetTrackingsParams struct {
	/**
	 * Page to show. (Default: 1)
	 */
	Page int `url:"page,omitempty" json:"page,omitempty"`

	/**
	 * Number of trackings each page contain. (Default: 100, Max: 200)
	 */
	Limit int `url:"limit,omitempty" json:"limit,omitempty"`

	/**
	 * Search the content of the tracking record fields:
	 * tracking_number,  title,  order_id,  customer_name,  custom_fields,  order_id,  emails,  smses
	 */
	Keyword string `url:"keyword,omitempty" json:"keyword,omitempty"`

	/**
	 * Tracking number of shipments. Use comma to separate multiple values
	 * (Example: RA123456789US,LE123456789US)
	 */
	TrackingNumbers string `url:"tracking_numbers,omitempty" json:"tracking_numbers,omitempty"`

	/**
	 * Unique courier code Use comma for multiple values. (Example: dhl,ups,usps)
	 */
	Slug string `url:"slug,omitempty" json:"slug,omitempty"`

	/**
	 * Total delivery time in days.
	 * - Difference of 1st checkpoint time and delivered time for delivered shipments
	 * - Difference of 1st checkpoint time and current time for non-delivered shipments
	 * Value as 0 for pending shipments or delivered shipment with only one checkpoint.
	 */
	DeliveryTime int `url:"delivery_time,omitempty" json:"delivery_time,omitempty"`

	/**
	 * Origin country of trackings. Use ISO Alpha-3 (three letters). Use comma for multiple values. (Example: USA,HKG)
	 */
	Origin string `url:"origin,omitempty" json:"origin,omitempty"`

	/**
	 * Destination country of trackings. Use ISO Alpha-3 (three letters).
	 * Use comma for multiple values. (Example: USA,HKG)
	 */
	Destination string `url:"destination,omitempty" json:"destination,omitempty"`

	/**
	 * Current status of tracking.
	 */
	Tag string `url:"tag,omitempty" json:"tag,omitempty"`

	/**
	 * Start date and time of trackings created. AfterShip only stores data of 90 days.
	 * (Defaults: 30 days ago, Example: 2013-03-15T16:41:56+08:00)
	 */
	CreatedAtMin string `url:"created_at_min,omitempty" json:"created_at_min,omitempty"`

	/**
	 * End date and time of trackings created.
	 * (Defaults: now, Example: 2013-04-15T16:41:56+08:00)
	 */
	CreatedAtMax string `url:"created_at_max,omitempty" json:"created_at_max,omitempty"`

	/**
	 * Start date and time of trackings updated.
	 * (Example: 2013-04-15T16:41:56+08:00)
	 */
	UpdatedAtMin string `url:"updated_at_min,omitempty" json:"updated_at_min,omitempty"`

	/**
	 * End date and time of trackings updated. (Example: 2013-04-15T16:41:56+08:00)
	 */
	UpdatedAtMax string `url:"updated_at_max,omitempty" json:"updated_at_max,omitempty"`

	/**
	 * List of fields to include in the response.
	 * Use comma for multiple values. Fields to include: title,  order_id,  tag,
	 * checkpoints,  checkpoint_time,  message,  country_name
	 * Defaults: none, Example: title,order_id
	 */
	Fields string `url:"fields,omitempty" json:"fields,omitempty"`

	/**
	 * Default: '' / Example: 'en'
	 * Support Chinese to English translation for  china-ems  and  china-post  only
	 */
	Lang string `url:"lang,omitempty" json:"lang,omitempty"`

	/**
	 * Tracking last updated at
	 * (Example: 2013-03-15T16:41:56+08:00)
	 */
	LastUpdatedAt string `url:"last_updated_at,omitempty" json:"last_updated_at,omitempty"`

	/**
	 * Select return to sender, the value should be true or false,
	 * with optional comma separated.
	 */
	ReturnToSender string `url:"return_to_sender,omitempty" json:"return_to_sender,omitempty"`

	/**
	 * Destination country of trackings returned by courier.
	 * Use ISO Alpha-3 (three letters).
	 * Use comma for multiple values. (Example: USA,HKG)
	 */
	CourierDestinationCountryIso3 string `url:"courier_destination_country_iso3,omitempty" json:"courier_destination_country_iso3,omitempty"`
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
	/**
	 * The postal code of receiver's address. Required by some couriers, such asdeutsch-post
	 */
	TrackingPostalCode string `url:"tracking_postal_code,omitempty" json:"tracking_postal_code,omitempty"`

	/**
	 * Shipping date in YYYYMMDD format. Required by some couriers, such asdeutsch-post
	 */
	TrackingShipDate string `url:"tracking_ship_date,omitempty" json:"tracking_ship_date,omitempty"`

	/**
	 * Destination Country of the shipment for a specific courier. Required by some couriers, such aspostnl-3s
	 */
	TrackingDestinationCountry string `url:"tracking_destination_country,omitempty" json:"tracking_destination_country,omitempty"`

	/**
	 * Account number of the shipper for a specific courier. Required by some couriers, such asdynamic-logistics
	 */
	TrackingAccountNumber string `url:"tracking_account_number,omitempty" json:"tracking_account_number,omitempty"`

	/**
	 * Key of the shipment for a specific courier. Required by some couriers, such assic-teliway
	 */
	TrackingKey string `url:"tracking_key,omitempty" json:"tracking_key,omitempty"`

	/**
	 * Origin Country of the shipment for a specific courier. Required by some couriers, such asdhl
	 */
	TrackingOriginCountry string `url:"tracking_origin_country,omitempty" json:"tracking_origin_country,omitempty"`

	/**
	 * Located state of the shipment for a specific courier. Required by some couriers, such asstar-track-courier
	 */
	TrackingState string `url:"tracking_state,omitempty" json:"tracking_state,omitempty"`
}

// TrackingCompletedStatus is status to make the tracking as completed
type TrackingCompletedStatus string

// TrackingCompletedStatusDelivered is reason DELIVERED to make the tracking as completed
const TrackingCompletedStatusDelivered TrackingCompletedStatus = "DELIVERED"

// TrackingCompletedStatusLost is reason LOST to make the tracking as completed
const TrackingCompletedStatusLost TrackingCompletedStatus = "LOST"

// TrackingCompletedStatusReturnedToSender is reason RETURNED_TO_SENDER to make the tracking as completed
const TrackingCompletedStatusReturnedToSender TrackingCompletedStatus = "RETURNED_TO_SENDER"

// createTrackingRequest is a model for create tracking API request
type createTrackingRequest struct {
	Tracking CreateTrackingParams `json:"tracking"`
}

// CreateTracking creates a new tracking
func (client *Client) CreateTracking(ctx context.Context, params CreateTrackingParams) (Tracking, error) {
	if params.TrackingNumber == "" {
		return Tracking{}, errors.New(errMissingTrackingNumber)
	}

	var trackingWrapper trackingWrapper
	err := client.makeRequest(ctx, http.MethodPost, "/trackings", nil,
		&createTrackingRequest{Tracking: params}, &trackingWrapper)
	return trackingWrapper.Tracking, err
}

// DeleteTracking deletes a tracking.
func (client *Client) DeleteTracking(ctx context.Context, identifier TrackingIdentifier) (Tracking, error) {
	uriPath, err := identifier.URIPath()
	if err != nil {
		return Tracking{}, errors.Wrap(err, "error deleting tracking")
	}

	uriPath = fmt.Sprintf("/trackings%s", uriPath)
	var trackingWrapper trackingWrapper
	err = client.makeRequest(ctx, http.MethodDelete, uriPath, nil, nil, &trackingWrapper)
	return trackingWrapper.Tracking, err
}

// GetTrackings gets tracking results of multiple trackings.
func (client *Client) GetTrackings(ctx context.Context, params GetTrackingsParams) (PagedTrackings, error) {
	var pagedTrackings PagedTrackings
	err := client.makeRequest(ctx, http.MethodGet, "/trackings", params, nil, &pagedTrackings)
	return pagedTrackings, err
}

// GetTracking gets tracking results of a single tracking.
func (client *Client) GetTracking(ctx context.Context, identifier TrackingIdentifier, params GetTrackingParams) (Tracking, error) {
	uriPath, err := identifier.URIPath()
	if err != nil {
		return Tracking{}, errors.Wrap(err, "error getting tracking")
	}

	uriPath = fmt.Sprintf("/trackings%s", uriPath)
	var trackingWrapper trackingWrapper
	err = client.makeRequest(ctx, http.MethodGet, uriPath, params, nil, &trackingWrapper)
	return trackingWrapper.Tracking, err
}

// updateTrackingRequest is a model for update tracking API request
type updateTrackingRequest struct {
	Tracking UpdateTrackingParams `json:"tracking"`
}

// UpdateTracking updates a tracking.
func (client *Client) UpdateTracking(ctx context.Context, identifier TrackingIdentifier, params UpdateTrackingParams) (Tracking, error) {
	uriPath, err := identifier.URIPath()
	if err != nil {
		return Tracking{}, errors.Wrap(err, "error updating tracking")
	}

	uriPath = fmt.Sprintf("/trackings%s", uriPath)
	var trackingWrapper trackingWrapper
	err = client.makeRequest(ctx, http.MethodPut, uriPath, nil,
		&updateTrackingRequest{params}, &trackingWrapper)
	return trackingWrapper.Tracking, err
}

// RetrackTracking retracks an expired tracking. Max 3 times per tracking.
func (client *Client) RetrackTracking(ctx context.Context, identifier TrackingIdentifier) (Tracking, error) {
	uriPath, err := identifier.URIPath()
	if err != nil {
		return Tracking{}, errors.Wrap(err, "error retracking")
	}

	uriPath = fmt.Sprintf("/trackings%s/retrack", uriPath)
	var trackingWrapper trackingWrapper
	err = client.makeRequest(ctx, http.MethodPost, uriPath, nil, nil, &trackingWrapper)
	return trackingWrapper.Tracking, err
}

// markAsCompletedRequest is a model for update tracking as completed API request
type markAsCompletedRequest struct {
	Reason string `json:"reason"` // One of "DELIVERED", "LOST" or "RETURNED_TO_SENDER".
}

// MarkTrackingAsCompleted marks a tracking as completed. The tracking won't auto update until retrack it.
func (client *Client) MarkTrackingAsCompleted(ctx context.Context, identifier TrackingIdentifier, status TrackingCompletedStatus) (Tracking, error) {
	uriPath, err := identifier.URIPath()
	if err != nil {
		return Tracking{}, errors.Wrap(err, "error marking tracking as completed")
	}

	uriPath = fmt.Sprintf("/trackings%s/mark-as-completed", uriPath)
	var trackingWrapper trackingWrapper
	err = client.makeRequest(ctx, http.MethodPost, uriPath,
		nil, &markAsCompletedRequest{Reason: string(status)}, &trackingWrapper)
	return trackingWrapper.Tracking, err
}
