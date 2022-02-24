package aftership

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateTracking(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/trackings", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		w.Write([]byte(`{
			"meta": {
					"code": 201
			},
			"data": {
					"tracking": {
							"id": "5b766a5cc7c33c0e007de3c9",
							"created_at": "2018-08-17T06:25:32+00:00",
							"updated_at": "2018-08-17T06:25:32+00:00",
							"last_updated_at": "2018-08-17T06:25:32+00:00",
							"tracking_number": "1111111111111",
							"slug": "fedex",
							"active": true,
							"android": [],
							"custom_fields": null,
							"customer_name": null,
							"delivery_time": 0,
							"destination_country_iso3": null,
							"courier_destination_country_iso3": null,
							"emails": [],
							"expected_delivery": null,
							"ios": [],
							"note": null,
							"order_id": null,
							"order_id_path": null,
                            "order_date": null,
                            "order_number": null,
							"origin_country_iso3": null,
							"shipment_package_count": 0,
							"shipment_pickup_date": null,
							"shipment_delivery_date": null,
							"shipment_type": null,
							"shipment_weight": null,
							"shipment_weight_unit": null,
							"signed_by": null,
							"smses": [],
							"source": "api",
							"tag": "Pending",
							"subtag": "Pending_001",
							"subtag_message": "Pending",
							"title": "1111111111111",
							"tracked_count": 0,
							"last_mile_tracking_supported": null,
							"language": null,
							"unique_token": "deprecated",
							"checkpoints": [],
							"subscribed_smses": [],
							"subscribed_emails": [],
							"return_to_sender": false,
							"tracking_account_number": null,
							"tracking_origin_country": null,
							"tracking_destination_country": null,
							"tracking_key": null,
							"tracking_postal_code": null,
							"tracking_ship_date": null,
							"tracking_state": null,
							"order_promised_delivery_date": "2019-05-20",
							"delivery_type": "pickup_at_store",
							"pickup_location": "Flagship Store",
							"pickup_note": "Contact shop keepers when you arrive our stores for shipment pickup",
							"courier_tracking_link": "https://www.fedex.com/fedextrack/?tracknumbers=1111111111111&cntry_code=us",
							"courier_redirect_link": "https://www.fedex.com/track?loc=en_US&tracknum=111111111111&requester=WT/trackdetails",
							"first_attempted_at": null,
                            "aftership_estimated_delivery_date": null
					}
			}
	}`))
	})

	params := CreateTrackingParams{
		Slug:           "dhl",
		TrackingNumber: "123456789",
		Title:          "Title Name",
		SMSes: []string{
			"+18555072509",
			"+18555072501",
		},
		Emails: []string{
			"email@yourdomain.com",
			"another_email@yourdomain.com",
		},
		OrderID:     "ID 1234",
		OrderIDPath: "http://www.aftership.com/order_id=1234",
		CustomFields: map[string]string{
			"product_name":  "iPhone Case",
			"product_price": "USD19.99",
		},
		Language:                  "en",
		OrderPromisedDeliveryDate: "2019-05-20",
		DeliveryType:              "pickup_at_store",
		PickupLocation:            "Flagship Store",
		PickupNote:                "Reach out to our staffs when you arrive our stores for shipment pickup",
	}

	createdAt, _ := time.Parse(time.RFC3339, "2018-08-17T06:25:32+00:00")
	updatedAt, _ := time.Parse(time.RFC3339, "2018-08-17T06:25:32+00:00")
	lastUpdatedAt, _ := time.Parse(time.RFC3339, "2018-08-17T06:25:32+00:00")
	tracking := Tracking{
		ID:                        "5b766a5cc7c33c0e007de3c9",
		CreatedAt:                 &createdAt,
		UpdatedAt:                 &updatedAt,
		LastUpdatedAt:             &lastUpdatedAt,
		TrackingNumber:            "1111111111111",
		Slug:                      "fedex",
		Active:                    true,
		Android:                   []string{},
		DeliveryTime:              0,
		Emails:                    []string{},
		IOS:                       []string{},
		SMSes:                     []string{},
		Source:                    "api",
		Tag:                       "Pending",
		Subtag:                    "Pending_001",
		SubtagMessage:             "Pending",
		Title:                     "1111111111111",
		UniqueToken:               "deprecated",
		Checkpoints:               []Checkpoint{},
		SubscribedSMSes:           []string{},
		SubscribedEmails:          []string{},
		ReturnToSender:            false,
		OrderPromisedDeliveryDate: "2019-05-20",
		DeliveryType:              "pickup_at_store",
		PickupLocation:            "Flagship Store",
		PickupNote:                "Contact shop keepers when you arrive our stores for shipment pickup",
		CourierTrackingLink:       "https://www.fedex.com/fedextrack/?tracknumbers=1111111111111&cntry_code=us",
		CourierRedirectLink:       "https://www.fedex.com/track?loc=en_US&tracknum=111111111111&requester=WT/trackdetails",
	}

	res, err := client.CreateTracking(context.Background(), params)
	assert.Equal(t, tracking, res)
	assert.Nil(t, err)
}

func TestCreateTrackingError(t *testing.T) {
	params := CreateTrackingParams{
		TrackingNumber: "",
	}

	_, err := client.CreateTracking(context.Background(), params)
	assert.NotNil(t, err)
	assert.Equal(t, errMissingTrackingNumber, err.Error())
}

func TestDeleteTracking(t *testing.T) {
	setup()
	defer teardown()

	p := SlugTrackingNumber{
		Slug:           "fedex",
		TrackingNumber: "772857780801111",
	}

	uri := fmt.Sprintf("/trackings/%s/%s", p.Slug, p.TrackingNumber)
	mux.HandleFunc(uri, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)
		w.Write([]byte(`{
			"meta": {
					"code": 200
			},
			"data": {
					"tracking": {
							"id": "5b7658cec7c33c0e007de3c5",
							"tracking_number": "772857780801111",
							"slug": "fedex",
							"tracking_account_number": null,
							"tracking_origin_country": null,
							"tracking_destination_country": null,
							"tracking_key": null,
							"tracking_postal_code": null,
							"tracking_ship_date": null,
							"tracking_state": null
					}
			}
	}`))
	})

	exp := Tracking{
		ID:             "5b7658cec7c33c0e007de3c5",
		TrackingNumber: "772857780801111",
		Slug:           "fedex",
	}

	res, _ := client.DeleteTracking(context.Background(), p)
	assert.Equal(t, exp, res)
}

func TestDeleteTrackingByID(t *testing.T) {
	setup()
	defer teardown()

	var id TrackingID = "5b7658cec7c33c0e007de3c5"

	uri := fmt.Sprintf("/trackings/%s", id)
	mux.HandleFunc(uri, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)
		w.Write([]byte(`{
			"meta": {
					"code": 200
			},
			"data": {
					"tracking": {
							"id": "5b7658cec7c33c0e007de3c5",
							"tracking_number": "772857780801111",
							"slug": "fedex"
					}
			}
	}`))
	})

	exp := Tracking{
		ID:             "5b7658cec7c33c0e007de3c5",
		TrackingNumber: "772857780801111",
		Slug:           "fedex",
	}

	res, _ := client.DeleteTracking(context.Background(), id)
	assert.Equal(t, exp, res)
}

func TestDeleteTrackingError(t *testing.T) {
	// empty slug or tracking_number
	p := SlugTrackingNumber{
		Slug:           "fedex",
		TrackingNumber: "",
	}

	_, err := client.DeleteTracking(context.Background(), p)
	assert.NotNil(t, err)
}

func TestGetTrackings(t *testing.T) {
	setup()
	defer teardown()

	p := GetTrackingsParams{
		Page:  1,
		Limit: 10,
	}

	mux.HandleFunc("/trackings", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		w.Write([]byte(`{
			"meta": {
					"code": 200
			},
			"data": {
					"page": 1,
					"limit": 100,
					"count": 3,
					"keyword": "",
					"slug": "",
					"origin": [],
					"destination": [],
					"tag": "",
					"fields": "",
					"created_at_min": "2018-05-19T06:23:00+00:00",
					"created_at_max": "2018-08-17T06:23:59+00:00",
					"last_updated_at": null,
					"return_to_sender": [],
					"courier_destination_country_iso3": [],
					"trackings": [
							{
									"id": "5b74f4958776db0e00b6f5ed",
									"created_at": "2018-08-16T03:50:45+00:00",
									"updated_at": "2018-08-16T03:50:54+00:00",
									"last_updated_at": "2018-08-16T03:50:53+00:00",
									"tracking_number": "1111111111111",
									"slug": "fedex",
									"active": false,
									"android": [],
									"custom_fields": null,
									"customer_name": null,
									"delivery_time": 2,
									"destination_country_iso3": null,
									"courier_destination_country_iso3": null,
									"emails": [],
									"expected_delivery": null,
									"ios": [],
									"note": null,
									"order_id": null,
									"order_id_path": null,
                                    "order_date": null,
                                    "order_number": null,
									"origin_country_iso3": "USA",
									"shipment_package_count": 1,
									"shipment_pickup_date": "2018-07-31T06:00:00",
									"shipment_delivery_date": "2018-08-01T17:19:47",
									"shipment_type": "FedEx Home Delivery",
									"shipment_weight": null,
									"shipment_weight_unit": "kg",
									"signed_by": "Signature not required",
									"smses": [],
									"source": "web",
									"tag": "Delivered",
									"subtag": "Delivered_001",
									"subtag_message": "Delivered",
									"title": "1111111111111",
									"tracked_count": 1,
									"last_mile_tracking_supported": null,
									"language": null,
									"unique_token": "deprecated",
									"checkpoints": [
											{
													"slug": "fedex",
													"city": null,
													"created_at": "2018-08-16T03:50:47+00:00",
													"location": null,
													"country_name": null,
													"message": "Shipment information sent to FedEx",
													"country_iso3": null,
													"tag": "InfoReceived",
													"subtag": "InfoReceived_001",
													"subtag_message": "Info Received",
													"checkpoint_time": "2018-07-31T10:33:00-04:00",
													"coordinates": [],
													"state": null,
													"zip": null,
													"raw_tag": "FPX_L_RPIF"
											}
									],
									"subscribed_smses": [],
									"subscribed_emails": [],
									"return_to_sender": false,
									"tracking_account_number": null,
									"tracking_origin_country": null,
									"tracking_destination_country": null,
									"tracking_key": null,
									"tracking_postal_code": null,
									"tracking_ship_date": null,
									"tracking_state": null,
									"order_promised_delivery_date": null,
									"delivery_type": null,
									"pickup_location": null,
									"pickup_note": null,
									"courier_tracking_link": "https://www.fedex.com/fedextrack/?tracknumbers=111111111111&cntry_code=us",
									"courier_redirect_link": "https://www.fedex.com/track?loc=en_US&tracknum=111111111111&requester=WT/trackdetails",
									"first_attempted_at": "2018-08-01T13:19:47-04:00",
                                    "aftership_estimated_delivery_date": {
                                         "estimated_delivery_date": "2018-08-01",   
                                         "confidence_score": 0.8,   
                                         "estimated_delivery_date_min": "2018-08-01",   
                                         "estimated_delivery_date_max": "2018-08-05"   
                                    }
							},
							{
									"id": "5b0516676a810a1400eb5c1c",
									"created_at": "2018-05-23T07:21:11+00:00",
									"updated_at": "2018-06-22T07:21:57+00:00",
									"last_updated_at": "2018-06-22T07:21:57+00:00",
									"tracking_number": "2222222222222",
									"slug": "ups",
									"active": false,
									"android": [],
									"custom_fields": null,
									"customer_name": null,
									"delivery_time": 0,
									"destination_country_iso3": null,
									"courier_destination_country_iso3": null,
									"emails": [
											"asdfasdf@asdf.com"
									],
									"expected_delivery": null,
									"ios": [],
									"note": null,
									"order_id": null,
									"order_id_path": null,
                                    "order_date": null,
                                    "order_number": null,
									"origin_country_iso3": null,
									"shipment_package_count": 0,
									"shipment_pickup_date": null,
									"shipment_delivery_date": null,
									"shipment_type": null,
									"shipment_weight": null,
									"shipment_weight_unit": null,
									"signed_by": null,
									"smses": [
											"+85261234567",
											"+85291234567"
									],
									"source": "web",
									"tag": "Expired",
									"subtag": "Expired_001",
									"subtag_message": "Expired",
									"title": "12ASDF121312",
									"tracked_count": 42,
									"last_mile_tracking_supported": null,
									"language": null,
									"unique_token": "deprecated",
									"checkpoints": [],
									"subscribed_smses": [
											"+85222222222",
											"+8533333333"
									],
									"subscribed_emails": [
											"yoyo@yoyo.com",
											"yoyo2@yoyo.com"
									],
									"return_to_sender": false,
									"tracking_account_number": null,
									"tracking_origin_country": null,
									"tracking_destination_country": null,
									"tracking_key": null,
									"tracking_postal_code": null,
									"tracking_ship_date": null,
									"tracking_state": null,
									"order_promised_delivery_date": null,
									"delivery_type": null,
									"pickup_location": null,
									"pickup_note": null,
									"courier_tracking_link": "https://www.fedex.com/fedextrack/?tracknumbers=2222222222222&cntry_code=us",
									"courier_redirect_link": "https://www.fedex.com/track?loc=en_US&tracknum=2222222222222&requester=WT/trackdetails",
									"first_attempted_at": null,
                                    "aftership_estimated_delivery_date": null
							}
					]
			}
	}`))
	})

	createdAt, _ := time.Parse(time.RFC3339, "2018-08-16T03:50:45+00:00")
	updatedAt, _ := time.Parse(time.RFC3339, "2018-08-16T03:50:54+00:00")
	lastUpdatedAt, _ := time.Parse(time.RFC3339, "2018-08-16T03:50:53+00:00")
	checkpointCreatedAt, _ := time.Parse(time.RFC3339, "2018-08-16T03:50:47+00:00")

	checkpoint := Checkpoint{
		Slug:           "fedex",
		CreatedAt:      &checkpointCreatedAt,
		Message:        "Shipment information sent to FedEx",
		Tag:            "InfoReceived",
		Subtag:         "InfoReceived_001",
		SubtagMessage:  "Info Received",
		CheckpointTime: "2018-07-31T10:33:00-04:00",
		Coordinates:    []string{},
		RawTag:         "FPX_L_RPIF",
	}

	tracking1 := Tracking{
		ID:                   "5b74f4958776db0e00b6f5ed",
		CreatedAt:            &createdAt,
		UpdatedAt:            &updatedAt,
		LastUpdatedAt:        &lastUpdatedAt,
		TrackingNumber:       "1111111111111",
		Slug:                 "fedex",
		Active:               false,
		Android:              []string{},
		DeliveryTime:         2,
		Emails:               []string{},
		IOS:                  []string{},
		OriginCountryISO3:    "USA",
		ShipmentPackageCount: 1,
		ShipmentPickupDate:   "2018-07-31T06:00:00",
		ShipmentDeliveryDate: "2018-08-01T17:19:47",
		ShipmentType:         "FedEx Home Delivery",
		ShipmentWeightUnit:   "kg",
		SignedBy:             "Signature not required",
		SMSes:                []string{},
		Source:               "web",
		Tag:                  "Delivered",
		Subtag:               "Delivered_001",
		SubtagMessage:        "Delivered",
		Title:                "1111111111111",
		TrackedCount:         1,
		UniqueToken:          "deprecated",
		Checkpoints: []Checkpoint{
			checkpoint,
		},
		SubscribedSMSes:     []string{},
		SubscribedEmails:    []string{},
		ReturnToSender:      false,
		CourierTrackingLink: "https://www.fedex.com/fedextrack/?tracknumbers=111111111111&cntry_code=us",
		CourierRedirectLink: "https://www.fedex.com/track?loc=en_US&tracknum=111111111111&requester=WT/trackdetails",
		FirstAttemptedAt:    "2018-08-01T13:19:47-04:00",
		EstimatedDeliveryDate: EstimatedDeliveryDate{
			EstimatedDeliveryDate:    "2018-08-01",
			ConfidenceScore:          0.8,
			EstimatedDeliveryDateMin: "2018-08-01",
			EstimatedDeliveryDateMax: "2018-08-05",
		},
	}

	t2CreatedAt, _ := time.Parse(time.RFC3339, "2018-05-23T07:21:11+00:00")
	t2UpdatedAt, _ := time.Parse(time.RFC3339, "2018-06-22T07:21:57+00:00")
	t2LastUpdatedAt, _ := time.Parse(time.RFC3339, "2018-06-22T07:21:57+00:00")

	tracking2 := Tracking{
		ID:             "5b0516676a810a1400eb5c1c",
		CreatedAt:      &t2CreatedAt,
		UpdatedAt:      &t2UpdatedAt,
		LastUpdatedAt:  &t2LastUpdatedAt,
		TrackingNumber: "2222222222222",
		Slug:           "ups",
		Active:         false,
		Android:        []string{},
		DeliveryTime:   0,
		Emails: []string{
			"asdfasdf@asdf.com",
		},
		IOS:                  []string{},
		ShipmentPackageCount: 0,
		SMSes: []string{
			"+85261234567",
			"+85291234567",
		},
		Source:        "web",
		Tag:           "Expired",
		Subtag:        "Expired_001",
		SubtagMessage: "Expired",
		Title:         "12ASDF121312",
		TrackedCount:  42,
		UniqueToken:   "deprecated",
		Checkpoints:   []Checkpoint{},
		SubscribedSMSes: []string{
			"+85222222222",
			"+8533333333",
		},
		SubscribedEmails: []string{
			"yoyo@yoyo.com",
			"yoyo2@yoyo.com",
		},
		ReturnToSender:      false,
		CourierTrackingLink: "https://www.fedex.com/fedextrack/?tracknumbers=2222222222222&cntry_code=us",
		CourierRedirectLink: "https://www.fedex.com/track?loc=en_US&tracknum=2222222222222&requester=WT/trackdetails",
	}

	exp := PagedTrackings{
		Page:  1,
		Limit: 100,
		Count: 3,
		Trackings: []Tracking{
			tracking1,
			tracking2,
		},
	}

	res, err := client.GetTrackings(context.Background(), p)
	assert.Equal(t, exp, res)
	assert.Nil(t, err)
}

func TestGetTracking(t *testing.T) {
	setup()
	defer teardown()

	p := SlugTrackingNumber{
		Slug:           "fedex",
		TrackingNumber: "111111111111",
	}

	uri := fmt.Sprintf("/trackings/%s/%s", p.Slug, p.TrackingNumber)
	mux.HandleFunc(uri, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		w.Write([]byte(`{
		"meta": {
			"code": 200
		},
		"data": {
				"tracking": {
						"id": "5b7658cec7c33c0e007de3c5",
						"created_at": "2018-08-17T05:10:38+00:00",
						"updated_at": "2018-08-17T05:10:46+00:00",
						"last_updated_at": "2018-08-17T05:10:46+00:00",
						"tracking_number": "111111111111",
						"slug": "fedex",
						"active": false,
						"android": [],
						"custom_fields": null,
						"customer_name": null,
						"delivery_time": 2,
						"destination_country_iso3": "JPN",
						"courier_destination_country_iso3": "JPN",
						"emails": [],
						"expected_delivery": null,
						"ios": [],
						"note": null,
						"order_id": null,
						"order_id_path": null,
                        "order_date": null,
                        "order_number": null,
						"origin_country_iso3": "CHN",
						"shipment_package_count": 1,
						"shipment_pickup_date": "2018-07-23T08:58:00",
						"shipment_delivery_date": "2018-07-25T01:10:00",
						"shipment_type": "FedEx International Economy",
						"shipment_weight": 4.1,
						"shipment_weight_unit": "kg",
						"signed_by": "..KOSUTOKO",
						"smses": [],
						"source": "api",
						"tag": "Delivered",
						"subtag": "Delivered_001",
						"subtag_message": "Delivered",
						"title": "Title Name",
						"tracked_count": 1,
						"last_mile_tracking_supported": null,
						"language": null,
						"unique_token": "deprecated",
						"checkpoints": [
								{
										"slug": "fedex",
										"city": null,
										"created_at": "2018-08-17T05:10:41+00:00",
										"location": null,
										"country_name": null,
										"message": "Shipment information sent to FedEx",
										"country_iso3": null,
										"tag": "InfoReceived",
										"subtag": "InfoReceived_001",
										"subtag_message": "Info Received",
										"checkpoint_time": "2018-07-23T01:21:39-05:00",
										"coordinates": [],
										"state": null,
										"zip": null,
										"raw_tag": "FPX_L_RPIF"
								}
						],
						"subscribed_smses": [],
						"subscribed_emails": [],
						"return_to_sender": false,
						"tracking_account_number": null,
						"tracking_origin_country": null,
						"tracking_destination_country": null,
						"tracking_key": null,
						"tracking_postal_code": null,
						"tracking_ship_date": null,
						"tracking_state": null,
						"order_promised_delivery_date": "2019-05-02",
						"delivery_type": "pickup_at_store",
						"pickup_location": "Flagship Store",
						"pickup_note": null,
						"courier_tracking_link": "https://www.fedex.com/fedextrack/?tracknumbers=111111111111&cntry_code=us",
						"courier_redirect_link": "https://www.fedex.com/track?loc=en_US&tracknum=111111111111&requester=WT/trackdetails",
						"first_attempted_at": "2018-07-25T10:10:00+09:00",
                        "aftership_estimated_delivery_date": null
				}
		}
	}`))
	})

	createdAt, _ := time.Parse(time.RFC3339, "2018-08-17T05:10:38+00:00")
	updatedAt, _ := time.Parse(time.RFC3339, "2018-08-17T05:10:46+00:00")
	lastUpdatedAt, _ := time.Parse(time.RFC3339, "2018-08-17T05:10:46+00:00")
	checkpointCreatedAt, _ := time.Parse(time.RFC3339, "2018-08-17T05:10:41+00:00")

	checkpoint := Checkpoint{
		Slug:           "fedex",
		CreatedAt:      &checkpointCreatedAt,
		Message:        "Shipment information sent to FedEx",
		Tag:            "InfoReceived",
		Subtag:         "InfoReceived_001",
		SubtagMessage:  "Info Received",
		CheckpointTime: "2018-07-23T01:21:39-05:00",
		Coordinates:    []string{},
		RawTag:         "FPX_L_RPIF",
	}

	exp := Tracking{
		ID:                            "5b7658cec7c33c0e007de3c5",
		CreatedAt:                     &createdAt,
		UpdatedAt:                     &updatedAt,
		LastUpdatedAt:                 &lastUpdatedAt,
		TrackingNumber:                "111111111111",
		Slug:                          "fedex",
		Active:                        false,
		Android:                       []string{},
		DeliveryTime:                  2,
		DestinationCountryISO3:        "JPN",
		CourierDestinationCountryISO3: "JPN",
		Emails:                        []string{},
		IOS:                           []string{},
		OriginCountryISO3:             "CHN",
		ShipmentPackageCount:          1,
		ShipmentPickupDate:            "2018-07-23T08:58:00",
		ShipmentDeliveryDate:          "2018-07-25T01:10:00",
		ShipmentType:                  "FedEx International Economy",
		ShipmentWeight:                4.1,
		ShipmentWeightUnit:            "kg",
		SignedBy:                      "..KOSUTOKO",
		SMSes:                         []string{},
		Source:                        "api",
		Tag:                           "Delivered",
		Subtag:                        "Delivered_001",
		SubtagMessage:                 "Delivered",
		Title:                         "Title Name",
		TrackedCount:                  1,
		UniqueToken:                   "deprecated",
		Checkpoints: []Checkpoint{
			checkpoint,
		},
		SubscribedSMSes:           []string{},
		SubscribedEmails:          []string{},
		ReturnToSender:            false,
		OrderPromisedDeliveryDate: "2019-05-02",
		DeliveryType:              "pickup_at_store",
		PickupLocation:            "Flagship Store",
		CourierTrackingLink:       "https://www.fedex.com/fedextrack/?tracknumbers=111111111111&cntry_code=us",
		CourierRedirectLink:       "https://www.fedex.com/track?loc=en_US&tracknum=111111111111&requester=WT/trackdetails",
		FirstAttemptedAt:          "2018-07-25T10:10:00+09:00",
	}

	res, err := client.GetTracking(context.Background(), p, GetTrackingParams{})
	assert.Equal(t, exp, res)
	assert.Nil(t, err)
}

func TestGetTrackingError(t *testing.T) {
	// empty slug or tracking_number
	p := SlugTrackingNumber{
		Slug:           "",
		TrackingNumber: "",
	}

	_, err := client.GetTracking(context.Background(), p, GetTrackingParams{})
	assert.NotNil(t, err)

	// empty tracking id
	var id TrackingID = ""

	_, err = client.GetTracking(context.Background(), id, GetTrackingParams{})
	assert.NotNil(t, err)
}

func TestUpdateTracking(t *testing.T) {
	setup()
	defer teardown()

	p := SlugTrackingNumber{
		Slug:           "fedex",
		TrackingNumber: "111111111111",
	}

	uri := fmt.Sprintf("/trackings/%s/%s", p.Slug, p.TrackingNumber)
	mux.HandleFunc(uri, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method)
		w.Write([]byte(`{
				"meta": {
						"code": 200
				},
				"data": {
						"tracking": {
								"id": "5b74f4958776db0e00b6f5ed",
								"created_at": "2018-08-16T03:50:45+00:00",
								"updated_at": "2018-08-17T06:25:10+00:00",
								"last_updated_at": "2018-08-17T06:25:10+00:00",
								"tracking_number": "1111111111111",
								"slug": "fedex",
								"active": false,
								"android": [],
								"custom_fields": null,
								"customer_name": null,
								"delivery_time": 2,
								"destination_country_iso3": null,
								"courier_destination_country_iso3": null,
								"emails": [],
								"expected_delivery": null,
								"ios": [],
								"note": null,
								"order_id": null,
								"order_id_path": null,
                                "order_date": null,
                                "order_number": null,
								"origin_country_iso3": "USA",
								"shipment_package_count": 1,
								"shipment_pickup_date": "2018-07-31T06:00:00",
								"shipment_delivery_date": "2018-08-01T17:19:47",
								"shipment_type": "FedEx Home Delivery",
								"shipment_weight": null,
								"shipment_weight_unit": "kg",
								"signed_by": "Signature not required",
								"smses": [],
								"source": "web",
								"tag": "Delivered",
								"subtag": "Delivered_001",
								"subtag_message": null,
								"title": "1111111111111",
								"tracked_count": 1,
								"last_mile_tracking_supported": null,
								"language": null,
								"unique_token": "deprecated",
								"checkpoints": [
										{
												"slug": null,
												"city": "BROOKLYN",
												"created_at": "2018-08-16T03:50:47+00:00",
												"location": "BROOKLYN, NY",
												"country_name": null,
												"message": "Picked up",
												"country_iso3": null,
												"tag": "InTransit",
												"subtag": "InTransit_002",
												"subtag_message": null,
												"checkpoint_time": "2018-07-31T20:47:00",
												"coordinates": [],
												"state": "NY",
												"zip": null,
												"raw_tag": "FPX_L_RPIF"
										}
								],
								"subscribed_smses": [],
								"subscribed_emails": [
										"sample@aftership.com"
								],
								"return_to_sender": false,
								"tracking_account_number": null,
								"tracking_origin_country": null,
								"tracking_destination_country": null,
								"tracking_key": null,
								"tracking_postal_code": null,
								"tracking_ship_date": null,
								"tracking_state": null,
								"order_promised_delivery_date": "2019-05-20",
								"delivery_type": "pickup_at_store",
								"pickup_location": "Flagship Store",
								"pickup_note": "Contact shop keepers when you arrive our stores for shipment pickup",
								"courier_tracking_link": "https://www.fedex.com/fedextrack/?tracknumbers=1111111111111&cntry_code=us",
								"courier_redirect_link": "https://www.fedex.com/track?loc=en_US&tracknum=1111111111111&requester=WT/trackdetails",
								"first_attempted_at": "2018-08-01T17:19:47",
                                "aftership_estimated_delivery_date": null
						}
				}
		}`))
	})

	createdAt, _ := time.Parse(time.RFC3339, "2018-08-16T03:50:45+00:00")
	updatedAt, _ := time.Parse(time.RFC3339, "2018-08-17T06:25:10+00:00")
	lastUpdatedAt, _ := time.Parse(time.RFC3339, "2018-08-17T06:25:10+00:00")
	checkpointCreatedAt, _ := time.Parse(time.RFC3339, "2018-08-16T03:50:47+00:00")

	checkpoint := Checkpoint{
		City:           "BROOKLYN",
		CreatedAt:      &checkpointCreatedAt,
		Location:       "BROOKLYN, NY",
		Message:        "Picked up",
		Tag:            "InTransit",
		Subtag:         "InTransit_002",
		CheckpointTime: "2018-07-31T20:47:00",
		Coordinates:    []string{},
		State:          "NY",
		RawTag:         "FPX_L_RPIF",
	}

	exp := Tracking{
		ID:                   "5b74f4958776db0e00b6f5ed",
		CreatedAt:            &createdAt,
		UpdatedAt:            &updatedAt,
		LastUpdatedAt:        &lastUpdatedAt,
		TrackingNumber:       "1111111111111",
		Slug:                 "fedex",
		Active:               false,
		Android:              []string{},
		DeliveryTime:         2,
		Emails:               []string{},
		IOS:                  []string{},
		OriginCountryISO3:    "USA",
		ShipmentPackageCount: 1,
		ShipmentPickupDate:   "2018-07-31T06:00:00",
		ShipmentDeliveryDate: "2018-08-01T17:19:47",
		ShipmentType:         "FedEx Home Delivery",
		ShipmentWeightUnit:   "kg",
		SignedBy:             "Signature not required",
		SMSes:                []string{},
		Source:               "web",
		Tag:                  "Delivered",
		Subtag:               "Delivered_001",
		Title:                "1111111111111",
		TrackedCount:         1,
		UniqueToken:          "deprecated",
		Checkpoints: []Checkpoint{
			checkpoint,
		},
		SubscribedSMSes: []string{},
		SubscribedEmails: []string{
			"sample@aftership.com",
		},
		ReturnToSender:            false,
		OrderPromisedDeliveryDate: "2019-05-20",
		DeliveryType:              "pickup_at_store",
		PickupLocation:            "Flagship Store",
		PickupNote:                "Contact shop keepers when you arrive our stores for shipment pickup",
		CourierTrackingLink:       "https://www.fedex.com/fedextrack/?tracknumbers=1111111111111&cntry_code=us",
		CourierRedirectLink:       "https://www.fedex.com/track?loc=en_US&tracknum=1111111111111&requester=WT/trackdetails",
		FirstAttemptedAt:          "2018-08-01T17:19:47",
	}

	data := UpdateTrackingParams{
		Title: "New Title",
	}

	res, _ := client.UpdateTracking(context.Background(), p, data)
	assert.Equal(t, exp, res)
}

func TestUpdateTrackingError(t *testing.T) {
	// empty slug or tracking_number
	p := SlugTrackingNumber{
		Slug:           "",
		TrackingNumber: "",
	}

	data := UpdateTrackingParams{
		Title: "New Title",
	}

	_, err := client.UpdateTracking(context.Background(), p, data)
	assert.NotNil(t, err)
}

func TestRetrackTracking(t *testing.T) {
	setup()
	defer teardown()

	p := SlugTrackingNumber{
		Slug:           "russian-post",
		TrackingNumber: "RA223547577RU",
	}

	uri := fmt.Sprintf("/trackings/%s/%s/retrack", p.Slug, p.TrackingNumber)
	mux.HandleFunc(uri, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		w.Write([]byte(`{
			"meta": {
					"code": 200
			},
			"data": {
					"tracking": {
							"active": true,
							"id": "52e0dfd21246ff7488066941",
							"slug": "russian-post",
							"tracking_number": "RA223547577RU"
					}
			}
	}`))
	})

	exp := Tracking{
		ID:             "52e0dfd21246ff7488066941",
		Active:         true,
		Slug:           "russian-post",
		TrackingNumber: "RA223547577RU",
	}

	res, _ := client.RetrackTracking(context.Background(), p)
	assert.Equal(t, exp, res)
}

func TestRetrackTrackingError(t *testing.T) {
	// empty slug or tracking_number
	p := SlugTrackingNumber{
		Slug:           "",
		TrackingNumber: "",
	}

	_, err := client.RetrackTracking(context.Background(), p)
	assert.NotNil(t, err)
}

func TestMarkTrackingAsCompleted(t *testing.T) {
	setup()
	defer teardown()

	p := SlugTrackingNumber{
		Slug:           "fedex",
		TrackingNumber: "111111111111",
	}

	uri := fmt.Sprintf("/trackings/%s/%s/mark-as-completed", p.Slug, p.TrackingNumber)
	mux.HandleFunc(uri, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		w.Write([]byte(`{
			"meta": {
					"code": 200
			},
			"data": {
					"tracking": {
							"id": "5b7658cec7c33c0e007de3c5",
							"created_at": "2018-08-17T05:10:38+00:00",
							"updated_at": "2018-08-17T05:10:46+00:00",
							"last_updated_at": "2018-08-17T05:10:46+00:00",
							"tracking_number": "111111111111",
							"slug": "fedex",
							"active": false,
							"android": [],
							"custom_fields": null,
							"customer_name": null,
							"delivery_time": 2,
							"destination_country_iso3": "JPN",
							"courier_destination_country_iso3": "JPN",
							"emails": [],
							"expected_delivery": null,
							"ios": [],
							"note": null,
							"order_id": null,
							"order_id_path": null,
                            "order_date": null,
                            "order_number": null,
							"origin_country_iso3": "CHN",
							"shipment_package_count": 1,
							"shipment_pickup_date": "2018-07-23T08:58:00",
							"shipment_delivery_date": "2018-07-25T01:10:00",
							"shipment_type": "FedEx International Economy",
							"shipment_weight": 4,
							"shipment_weight_unit": "kg",
							"signed_by": "..KOSUTOKO",
							"smses": [],
							"source": "api",
							"tag": "Exception",
							"subtag": "Exception_013",
							"subtag_message": "Shipment lost",
							"title": "Title Name",
							"tracked_count": 1,
							"last_mile_tracking_supported": null,
							"language": null,
							"unique_token": "deprecated",
							"checkpoints": [
									{
											"slug": "fedex",
											"city": null,
											"created_at": "2018-08-17T05:10:41+00:00",
											"location": null,
											"country_name": null,
											"message": "Shipment information sent to FedEx",
											"country_iso3": null,
											"tag": "InfoReceived",
											"subtag": "InfoReceived_001",
											"subtag_message": "Info Received",
											"checkpoint_time": "2018-07-23T01:21:39-05:00",
											"coordinates": [],
											"state": null,
											"zip": null,
											"raw_tag": "FPX_L_RPIF"
									}
							],
							"subscribed_smses": [],
							"subscribed_emails": [],
							"return_to_sender": false,
							"tracking_account_number": null,
							"tracking_origin_country": null,
							"tracking_destination_country": null,
							"tracking_key": null,
							"tracking_postal_code": null,
							"tracking_ship_date": null,
							"tracking_state": null,
							"order_promised_delivery_date": "2019-05-02",
							"delivery_type": "pickup_at_store",
							"pickup_location": "Flagship Store",
							"pickup_note": null,
							"courier_tracking_link": "https://www.fedex.com/fedextrack/?tracknumbers=111111111111&cntry_code=us",
							"courier_redirect_link": "https://www.fedex.com/track?loc=en_US&tracknum=111111111111&requester=WT/trackdetails",
							"first_attempted_at": "2018-07-25T10:10:00+09:00",
                            "aftership_estimated_delivery_date": null
					}
			}
	}`))
	})

	createdAt, _ := time.Parse(time.RFC3339, "2018-08-17T05:10:38+00:00")
	updatedAt, _ := time.Parse(time.RFC3339, "2018-08-17T05:10:46+00:00")
	lastUpdatedAt, _ := time.Parse(time.RFC3339, "2018-08-17T05:10:46+00:00")
	checkpointCreatedAt, _ := time.Parse(time.RFC3339, "2018-08-17T05:10:41+00:00")

	checkpoint := Checkpoint{
		Slug:           "fedex",
		CreatedAt:      &checkpointCreatedAt,
		Message:        "Shipment information sent to FedEx",
		Tag:            "InfoReceived",
		Subtag:         "InfoReceived_001",
		SubtagMessage:  "Info Received",
		CheckpointTime: "2018-07-23T01:21:39-05:00",
		Coordinates:    []string{},
		RawTag:         "FPX_L_RPIF",
	}

	exp := Tracking{
		ID:                            "5b7658cec7c33c0e007de3c5",
		CreatedAt:                     &createdAt,
		UpdatedAt:                     &updatedAt,
		LastUpdatedAt:                 &lastUpdatedAt,
		TrackingNumber:                "111111111111",
		Slug:                          "fedex",
		Active:                        false,
		Android:                       []string{},
		DeliveryTime:                  2,
		DestinationCountryISO3:        "JPN",
		CourierDestinationCountryISO3: "JPN",
		Emails:                        []string{},
		IOS:                           []string{},
		OriginCountryISO3:             "CHN",
		ShipmentPackageCount:          1,
		ShipmentPickupDate:            "2018-07-23T08:58:00",
		ShipmentDeliveryDate:          "2018-07-25T01:10:00",
		ShipmentType:                  "FedEx International Economy",
		ShipmentWeight:                4,
		ShipmentWeightUnit:            "kg",
		SignedBy:                      "..KOSUTOKO",
		SMSes:                         []string{},
		Source:                        "api",
		Tag:                           "Exception",
		Subtag:                        "Exception_013",
		SubtagMessage:                 "Shipment lost",
		Title:                         "Title Name",
		TrackedCount:                  1,
		UniqueToken:                   "deprecated",
		Checkpoints: []Checkpoint{
			checkpoint,
		},
		SubscribedSMSes:           []string{},
		SubscribedEmails:          []string{},
		ReturnToSender:            false,
		OrderPromisedDeliveryDate: "2019-05-02",
		DeliveryType:              "pickup_at_store",
		PickupLocation:            "Flagship Store",
		CourierTrackingLink:       "https://www.fedex.com/fedextrack/?tracknumbers=111111111111&cntry_code=us",
		CourierRedirectLink:       "https://www.fedex.com/track?loc=en_US&tracknum=111111111111&requester=WT/trackdetails",
		FirstAttemptedAt:          "2018-07-25T10:10:00+09:00",
	}

	res, _ := client.MarkTrackingAsCompleted(context.Background(), p, TrackingCompletedStatusLost)
	assert.Equal(t, exp, res)
}

func TestMarkTrackingAsCompletedError(t *testing.T) {
	p := SlugTrackingNumber{
		Slug:           "",
		TrackingNumber: "",
	}

	_, err := client.MarkTrackingAsCompleted(context.Background(), p, TrackingCompletedStatusLost)
	assert.NotNil(t, err)
}
