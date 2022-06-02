package aftership

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCouriers(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/couriers", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		w.Write([]byte(`{
			"meta": {
					"code": 200
			},
			"data": {
					"total": 2,
					"couriers": [
							{
									"slug": "dhl",
									"name": "DHL",
									"phone": "+1 800 225 5345",
									"other_name": "DHL Express",
									"web_url": "http://www.dhl.com/",
									"required_fields": [],
									"optional_fields": []
							},
							{
									"slug": "deutsch-post",
									"name": "Deutsche Post Mail",
									"phone": "+49 (0) 180 2 000221",
									"other_name": "dpdhl",
									"web_url": "http://www.deutschepost.de/",
									"required_fields": [
											"tracking_ship_date"
									],
									"optional_fields": []
							}
					]
			}
	}`))
	})

	exp := CourierList{
		Total: 2,
		Couriers: []Courier{
			{
				Slug:           "dhl",
				Name:           "DHL",
				Phone:          "+1 800 225 5345",
				OtherName:      "DHL Express",
				WebURL:         "http://www.dhl.com/",
				RequiredFields: []string{},
				OptionalFields: []string{},
			},
			{
				Slug:           "deutsch-post",
				Name:           "Deutsche Post Mail",
				Phone:          "+49 (0) 180 2 000221",
				OtherName:      "dpdhl",
				WebURL:         "http://www.deutschepost.de/",
				RequiredFields: []string{"tracking_ship_date"},
				OptionalFields: []string{},
			},
		},
	}

	res, err := client.GetCouriers(context.Background())
	assert.Equal(t, exp, res)
	assert.Nil(t, err)
}

func TestGetAllCouriers(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/couriers/all", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		w.Write([]byte(`{
			"meta": {
					"code": 200
			},
			"data": {
					"total": 3,
					"couriers": [
							{
									"slug": "india-post-int",
									"name": "India Post International",
									"phone": "+91 1800 11 2011",
									"other_name": "भारतीय डाक, Speed Post & eMO, EMS, IPS Web",
									"web_url": "http://www.indiapost.gov.in/",
									"required_fields": [],
									"optional_fields": []
							},
							{
									"slug": "italy-sda",
									"name": "Italy SDA",
									"phone": "+39 199 113366",
									"other_name": "SDA Express Courier",
									"web_url": "http://www.sda.it/",
									"required_fields": [],
									"optional_fields": []
							},
							{
									"slug": "bpost",
									"name": "Belgium Post",
									"phone": "+32 2 276 22 74",
									"other_name": "bpost, Belgian Post",
									"web_url": "http://www.bpost.be/",
									"required_fields": [],
									"optional_fields": []
							}
					]
			}
	}`))
	})

	exp := CourierList{
		Total: 3,
		Couriers: []Courier{
			{
				Slug:           "india-post-int",
				Name:           "India Post International",
				Phone:          "+91 1800 11 2011",
				OtherName:      "भारतीय डाक, Speed Post & eMO, EMS, IPS Web",
				WebURL:         "http://www.indiapost.gov.in/",
				RequiredFields: []string{},
				OptionalFields: []string{},
			},
			{
				Slug:           "italy-sda",
				Name:           "Italy SDA",
				Phone:          "+39 199 113366",
				OtherName:      "SDA Express Courier",
				WebURL:         "http://www.sda.it/",
				RequiredFields: []string{},
				OptionalFields: []string{},
			},
			{
				Slug:           "bpost",
				Name:           "Belgium Post",
				Phone:          "+32 2 276 22 74",
				OtherName:      "bpost, Belgian Post",
				WebURL:         "http://www.bpost.be/",
				RequiredFields: []string{},
				OptionalFields: []string{},
			},
		},
	}

	res, err := client.GetAllCouriers(context.Background())
	assert.Equal(t, exp, res)
	assert.Nil(t, err)
}

func TestInvalidDetectCouriers(t *testing.T) {
	_, err := client.DetectCouriers(context.Background(), CourierDetectionParams{})

	assert.NotNil(t, err)
	assert.Equal(t, errMissingTrackingNumber, err.Error())
}

func TestDetectCouriers(t *testing.T) {

	setup()
	defer teardown()

	mux.HandleFunc("/couriers/detect", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		w.Write([]byte(`{
			"meta": {
					"code": 200
			},
			"data": {
					"total": 2,
					"couriers": [
							{
									"slug": "fedex",
									"name": "FedEx",
									"phone": "+1 800 247 4747",
									"other_name": "Federal Express",
									"web_url": "http://www.fedex.com/",
									"required_fields": [],
									"optional_fields": []
							},
							{
									"slug": "dx",
									"name": "DX",
									"phone": "+44 0844 826 1178",
									"other_name": "DX Freight",
									"web_url": "https://www.thedx.co.uk/",
									"required_fields": [
											"tracking_postal_code"
									],
									"optional_fields": []
							}
					]
			}
	}`))
	})

	exp := CourierList{
		Total: 2,
		Couriers: []Courier{
			{
				Slug:           "fedex",
				Name:           "FedEx",
				Phone:          "+1 800 247 4747",
				OtherName:      "Federal Express",
				WebURL:         "http://www.fedex.com/",
				RequiredFields: []string{},
				OptionalFields: []string{},
			},
			{
				Slug:           "dx",
				Name:           "DX",
				Phone:          "+44 0844 826 1178",
				OtherName:      "DX Freight",
				WebURL:         "https://www.thedx.co.uk/",
				RequiredFields: []string{"tracking_postal_code"},
				OptionalFields: []string{},
			},
		},
	}

	res, err := client.DetectCouriers(context.Background(), CourierDetectionParams{
		TrackingNumber: "906587618687",
	})

	assert.Equal(t, exp, res)
	assert.Nil(t, err)
}
