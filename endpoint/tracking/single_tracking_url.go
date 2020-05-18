package tracking

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/aftership/aftership-sdk-go/v2/error"
	"github.com/google/go-querystring/query"
)

// SingleTrackingParam identifies a single tracking
// its mandatory to provide either Id or Slug and TrackingNumber both
type SingleTrackingParam struct {
	ID             string
	Slug           string
	TrackingNumber string
	OptionalParams *SingleTrackingOptionalParams
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

// BuildTrackingURL returns the tracking URL
func (param *SingleTrackingParam) BuildTrackingURL(path string, subPath string) (string, *error.AfterShipError) {
	if path == "" {
		path = "trackings"
	}

	var trackingURL string
	if param.ID != "" {
		trackingURL = fmt.Sprintf("/%s/%s", path, url.QueryEscape(param.ID))
	} else if param.Slug != "" && param.TrackingNumber != "" {
		trackingURL = fmt.Sprintf("/%s/%s/%s", path, url.QueryEscape(param.Slug), url.QueryEscape(param.TrackingNumber))
	} else {
		return "", error.MakeSdkError(error.ErrorTypeHandlerError, "You must specify the id or slug and tracking number", param)
	}

	if subPath != "" {
		trackingURL += fmt.Sprintf("/%s", subPath)
	}

	if param.OptionalParams != nil {
		url, err := BuildURLWithQueryString(trackingURL, param.OptionalParams)
		if err != nil {
			return "", err
		}
		return url, nil
	}

	return trackingURL, nil
}

// BuildURLWithQueryString returns the url with query string
func BuildURLWithQueryString(uri string, params interface{}) (string, *error.AfterShipError) {
	if params == nil {
		return uri, nil
	}

	queryStringObj, err := query.Values(params)
	if err != nil {
		return "", error.MakeSdkError(error.ErrorTypeHandlerError, err.Error(), params)
	}

	queryString := queryStringObj.Encode()
	if queryString != "" {
		conn := "?"
		if strings.Contains(uri, "?") {
			conn = "&"
		}

		uri += conn + queryString
	}

	return uri, nil
}
