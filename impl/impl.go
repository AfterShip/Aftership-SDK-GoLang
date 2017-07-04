package impl

import (
	"net/http"

	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"time"

	"github.com/AfterShip/aftership-sdk-go/apiV4"
	"github.com/google/go-querystring/query"
)

type AfterShipApiV4Impl struct {
	ApiKey      string
	RetryPolicy *apiV4.RetryPolicy
	Client      *http.Client
}

// GetCouriers returns a list of couriers activated at your AfterShip account.
func (api *AfterShipApiV4Impl) GetCouriers() ([]apiV4.Courier, apiV4.AfterShipApiError) {
	var courierEnvelope apiV4.CourierEnvelope
	err := api.request("GET", apiV4.COURIERS_ENDPOINT, &courierEnvelope, nil)
	return courierEnvelope.Data.Couriers, err
}

// GetAll returns a list of all couriers.
func (api *AfterShipApiV4Impl) GetAllCouriers() ([]apiV4.Courier, apiV4.AfterShipApiError) {
	var courierEnvelope apiV4.CourierEnvelope
	err := api.request("GET", apiV4.COURIERS_ALL_ENDPOINT, &courierEnvelope, nil)
	return courierEnvelope.Data.Couriers, err
}

// DetectCouriers returns a list of matched couriers based on tracking number format
// and selected couriers or a list of couriers.
func (api *AfterShipApiV4Impl) DetectCouriers(params apiV4.CourierDetectParam) ([]apiV4.Courier, apiV4.AfterShipApiError) {
	var courierEnvelope apiV4.CourierEnvelope
	err := api.request("POST", apiV4.COURIERS_DETECT_ENDPOINT, &courierEnvelope, &apiV4.CourierDetectParamReqBody{params})
	return courierEnvelope.Data.Couriers, err
}

// CreateTracking creates a new tracking
func (api *AfterShipApiV4Impl) CreateTracking(newTracking apiV4.NewTracking) (apiV4.Tracking, apiV4.AfterShipApiError) {
	var trackingEnvelope apiV4.TrackingEnvelope
	err := api.request("POST", apiV4.TRACKINGS_ENDPOINT, &trackingEnvelope, &apiV4.NewTrackingReqBody{newTracking})
	return trackingEnvelope.Data.Tracking, err
}

// DeleteTracking Deletes a tracking.
func (api *AfterShipApiV4Impl) DeleteTracking(id apiV4.TrackingId) (apiV4.DeletedTracking, apiV4.AfterShipApiError) {
	var deletedtrackingEnvelope apiV4.DeleteTrackingEnvelope
	var url string
	if id.Id != "" {
		url = apiV4.TRACKINGS_ENDPOINT + "/" + id.Id
	} else if id.Slug != "" && id.TrackingNumber != "" {
		url = apiV4.TRACKINGS_ENDPOINT + "/" + id.Slug + "/" + id.TrackingNumber
	}
	err := api.request("DELETE", url, &deletedtrackingEnvelope, nil)
	return deletedtrackingEnvelope.Data.Tracking, err
}

// GetTrackings Gets tracking results of multiple trackings.
func (api *AfterShipApiV4Impl) GetTrackings(params apiV4.GetTrackingsParams) (apiV4.TrackingsData, apiV4.AfterShipApiError) {
	var trackingsEnvelope apiV4.TrackingsEnvelope
	url := apiV4.TRACKINGS_ENDPOINT
	queryStringObj, err := query.Values(params)
	if err != nil {
		log.Fatal(err)
	}
	queryString := queryStringObj.Encode()
	if queryString != "" {
		url += "?" + queryString
	}
	error := api.request("GET", url, &trackingsEnvelope, nil)
	return trackingsEnvelope.Data, error
}

// GetTrackingsExport Gets all trackings results (for backup or analytics purpose)
func (api *AfterShipApiV4Impl) GetTrackingsExport(params apiV4.GetTrackingsParams) (apiV4.TrackingsData, apiV4.AfterShipApiError) {
	var trackingsEnvelope apiV4.TrackingsEnvelope
	url := apiV4.TRACKINGS_EXPORTS_ENDPOINT
	queryStringObj, err := query.Values(params)
	if err != nil {
		log.Fatal(err)
	}
	queryString := queryStringObj.Encode()
	if queryString != "" {
		url += "?" + queryString
	}
	error := api.request("GET", url, &trackingsEnvelope, nil)
	return trackingsEnvelope.Data, error
}

// GetTracking Gets tracking results of a single tracking.
func (api *AfterShipApiV4Impl) GetTracking(id apiV4.TrackingId, fields string, lang string) (apiV4.Tracking, apiV4.AfterShipApiError) {
	var trackingEnvelope apiV4.TrackingEnvelope
	var url string
	if id.Id != "" {
		url = apiV4.TRACKINGS_ENDPOINT + "/" + id.Id
	} else if id.Slug != "" && id.TrackingNumber != "" {
		url = apiV4.TRACKINGS_ENDPOINT + "/" + id.Slug + "/" + id.TrackingNumber
	}
	fieldsAdded := false
	if fields != "" {
		url += "?fields=" + fields
		fieldsAdded = true
	}
	if lang != "" {
		if fieldsAdded {
			url += "&"
		} else {
			url += "?"
		}
		url += "lang=" + lang
	}
	err := api.request("GET", url, &trackingEnvelope, nil)
	return trackingEnvelope.Data.Tracking, err
}

// UpdateTracking Updates a tracking.
func (api *AfterShipApiV4Impl) UpdateTracking(id apiV4.TrackingId, update apiV4.TrackingUpdate) (apiV4.Tracking, apiV4.AfterShipApiError) {
	var trackingEnvelope apiV4.TrackingEnvelope
	var url string
	if id.Id != "" {
		url = apiV4.TRACKINGS_ENDPOINT + "/" + id.Id
	} else if id.Slug != "" && id.TrackingNumber != "" {
		url = apiV4.TRACKINGS_ENDPOINT + "/" + id.Slug + "/" + id.TrackingNumber
	}
	err := api.request("PUT", url, &trackingEnvelope, &apiV4.TrackingUpdateReqBody{update})
	return trackingEnvelope.Data.Tracking, err
}

// ReTrack an expired tracking once. Max. 3 times per tracking.
func (api *AfterShipApiV4Impl) ReTrack(id apiV4.TrackingId) (apiV4.Tracking, apiV4.AfterShipApiError) {
	var trackingEnvelope apiV4.TrackingEnvelope
	var url string
	if id.Id != "" {
		url = apiV4.TRACKINGS_ENDPOINT + "/" + id.Id
	} else if id.Slug != "" && id.TrackingNumber != "" {
		url = apiV4.TRACKINGS_ENDPOINT + "/" + id.Slug + "/" + id.TrackingNumber
	}
	url += "/retrack"
	var body struct{}
	err := api.request("POST", url, &trackingEnvelope, body)
	return trackingEnvelope.Data.Tracking, err
}

// LastCheckPoint Return the tracking information of the last checkpoint of a single tracking.
func (api *AfterShipApiV4Impl) GetLastCheckPoint(id apiV4.TrackingId, fields string,
	lang string) (apiV4.LastCheckPoint, apiV4.AfterShipApiError) {
	var lastCheckPointEnvelope apiV4.LastCheckPointEnvelope
	var url string
	if id.Id != "" {
		url = apiV4.LAST_CHECKPOINT_ENDPOINT + "/" + id.Id
	} else if id.Slug != "" && id.TrackingNumber != "" {
		url = apiV4.LAST_CHECKPOINT_ENDPOINT + "/" + id.Slug + "/" + id.TrackingNumber
	}
	fieldsAdded := false
	if fields != "" {
		url += "?fields=" + fields
		fieldsAdded = true
	}
	if lang != "" {
		if fieldsAdded {
			url += "&"
		} else {
			url += "?"
		}
		url += "lang=" + lang
	}
	err := api.request("GET", url, &lastCheckPointEnvelope, nil)
	return lastCheckPointEnvelope.Data, err
}

// AddNotification Adds notifications to a tracking number.
func (api *AfterShipApiV4Impl) AddNotification(id apiV4.TrackingId,
	notification apiV4.NotificationSetting) (apiV4.NotificationSetting, apiV4.AfterShipApiError) {
	var notificationSettingEnvelope apiV4.NotificationSettingEnvelope
	var url string
	if id.Id != "" {
		url = apiV4.NOTIFICATIONS + "/" + id.Id
	} else if id.Slug != "" && id.TrackingNumber != "" {
		url = apiV4.NOTIFICATIONS + "/" + id.Slug + "/" + id.TrackingNumber
	}
	url += "/add"
	err := api.request("POST", url, &notificationSettingEnvelope, &apiV4.NotificationSettingWrapper{notification})
	return notificationSettingEnvelope.Data.Notification, err
}

// RemoveNotification Removes notifications from a tracking number.
func (api *AfterShipApiV4Impl) RemoveNotification(id apiV4.TrackingId,
	notification apiV4.NotificationSetting) (apiV4.NotificationSetting, apiV4.AfterShipApiError) {
	var notificationSettingEnvelope apiV4.NotificationSettingEnvelope
	var url string
	if id.Id != "" {
		url = apiV4.NOTIFICATIONS + "/" + id.Id
	} else if id.Slug != "" && id.TrackingNumber != "" {
		url = apiV4.NOTIFICATIONS + "/" + id.Slug + "/" + id.TrackingNumber
	}
	url += "/remove"
	err := api.request("POST", url, &notificationSettingEnvelope, &apiV4.NotificationSettingWrapper{notification})
	return notificationSettingEnvelope.Data.Notification, err
}

// GetNotificationSetting Gets notifications value from a tracking number.
func (api *AfterShipApiV4Impl) GetNotificationSetting(id apiV4.TrackingId, fields string) (apiV4.NotificationSetting,
	apiV4.AfterShipApiError) {
	var notificationSettingEnvelope apiV4.NotificationSettingEnvelope
	var url string
	if id.Id != "" {
		url = apiV4.NOTIFICATIONS + "/" + id.Id
	} else if id.Slug != "" && id.TrackingNumber != "" {
		url = apiV4.NOTIFICATIONS + "/" + id.Slug + "/" + id.TrackingNumber
	}
	if fields != "" {
		url += "?fields=" + fields
	}
	err := api.request("GET", url, &notificationSettingEnvelope, nil)
	return notificationSettingEnvelope.Data.Notification, err
}

// request is generic method to communicate all REST requests
func (api *AfterShipApiV4Impl) request(method string, endpoint string,
	result apiV4.Response, body interface{}) apiV4.AfterShipApiError {

	if api.Client == nil {
		api.Client = &http.Client{}
	}

	bodyStr, err := json.Marshal(body)
	if err != nil {
		return apiV4.AfterShipApiError{
			apiV4.ResponseMeta{
				apiV4.SDK_ERROR_CODE,
				fmt.Sprint(err),
				"JSON Error",
			},
		}
	}

	req, _ := http.NewRequest(method, apiV4.URL+endpoint, bytes.NewBuffer(bodyStr))
	// req, _ := http.NewRequest(method, "http://localhost:8080/post", bytes.NewBuffer(bodyStr))
	req.Header.Add(apiV4.API_KEY_HEADER_FIELD, api.ApiKey)
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("aftership-sdk-go", "v0.1")
	if body != nil {
		req.Header.Add("content-type", "application/json")
	}

	resp, err := api.Client.Do(req)
	if err != nil {
		return apiV4.AfterShipApiError{
			apiV4.ResponseMeta{
				apiV4.SDK_ERROR_CODE,
				fmt.Sprint(err),
				"IO Error",
			},
		}
	}
	//log.Print("X-RateLimit-Reset", resp.Header.Get("X-RateLimit-Reset"))
	//log.Print("X-RateLimit-Limit", resp.Header.Get("X-RateLimit-Limit"))
	//log.Print("X-RateLimit-Remaining", resp.Header.Get("X-RateLimit-Remaining"))
	rateLimitReset, _ := strconv.Atoi(resp.Header.Get("X-RateLimit-Reset"))

	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return apiV4.AfterShipApiError{
			apiV4.ResponseMeta{
				apiV4.SDK_ERROR_CODE,
				fmt.Sprint(err),
				"IO Error",
			},
		}
	}
	err = json.Unmarshal(contents, result)
	if err != nil {
		return apiV4.AfterShipApiError{
			apiV4.ResponseMeta{
				apiV4.SDK_ERROR_CODE,
				fmt.Sprint(err),
				"JSON Error",
			},
		}
	}
	code := result.ResponseCode().Code

	// handling rate limit error by sleeping and retrying after reset
	if code == 429 && api.RetryPolicy != nil && api.RetryPolicy.RetryOnHittingRateLimit {
		timeNow := time.Now().Unix()
		dur := time.Duration(int64(rateLimitReset)-timeNow) * time.Second
		log.Println("Hit rate limit, Auto retry after Dur : ", dur)
		c := time.After(dur)
		for {
			log.Println("Retrying start ", <-c)
			return api.request(method, endpoint, result, body)
		}
	}

	return apiV4.AfterShipApiError{
		result.ResponseCode(),
	}
}
