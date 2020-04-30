package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/aftership/aftership-sdk-go/v2"
	"github.com/aftership/aftership-sdk-go/v2/conf"
	"github.com/aftership/aftership-sdk-go/v2/tracking"
)

func main() {
	aftership := aftership.NewAfterShip(conf.AfterShipConf{
		AppKey: "YOUR_API_KEY",
	})

	// Create a tracking
	trackingNumber := strconv.FormatInt(time.Now().Unix(), 10)
	newTracking := tracking.NewTrackingRequest{
		Tracking: tracking.NewTracking{
			TrackingNumber: trackingNumber,
			Slug:           []string{"dhl"},
			Title:          "Title Name",
			Smses: []string{
				"+18555072509",
				"+18555072501",
			},
			Emails: []string{
				"email@yourdomain.com",
				"another_email@yourdomain.com",
			},
			OrderID: "ID 1234",
			CustomFields: map[string]string{
				"product_name":  "iPhone Case",
				"product_price": "USD19.99",
			},
			Language:                  "en",
			OrderPromisedDeliveryDate: "2019-05-20",
			DeliveryType:              "pickup_at_store",
			PickupLocation:            "Flagship Store",
			PickupNote:                "Reach out to our staffs when you arrive our stores for shipment pickup",
		},
	}

	result, err := aftership.Tracking.CreateTracking(newTracking)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	// Delete a tracking
	param := tracking.SingleTrackingParam{
		Slug:           "dhl",
		TrackingNumber: trackingNumber,
	}

	result, err = aftership.Tracking.DeleteTracking(param)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	// Get tracking results of multiple trackings.
	multiParams := tracking.MultiTrackingsParams{
		Page:  1,
		Limit: 10,
	}

	multiResults, err := aftership.Tracking.GetTrackings(multiParams)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(multiResults)
	}

	// Get tracking results of a single tracking.
	param = tracking.SingleTrackingParam{
		Slug:           "dhl",
		TrackingNumber: "1588226550",
	}

	result, err = aftership.Tracking.GetTracking(param, tracking.GetTrackingParams{})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	// Get tracking results of a single tracking by id.
	param = tracking.SingleTrackingParam{
		ID: "rymq9l34ztbvvk9md2ync00r",
	}

	result, err = aftership.Tracking.GetTracking(param, tracking.GetTrackingParams{
		Fields: "tracking_postal_code,title,order_id",
	})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	// Update a tracking.
	param = tracking.SingleTrackingParam{
		Slug:           "dhl",
		TrackingNumber: "1588226550",
	}

	updateReq := tracking.UpdateTrackingRequest{
		Tracking: tracking.UpdateTracking{
			Title: "New Title",
		},
	}

	result, err = aftership.Tracking.UpdateTracking(param, updateReq)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	// Retrack an expired tracking.
	param = tracking.SingleTrackingParam{
		Slug:           "dhl",
		TrackingNumber: "1588226550",
	}

	result, err = aftership.Tracking.ReTrack(param)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
