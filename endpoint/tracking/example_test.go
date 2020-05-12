package tracking_test

import (
	"fmt"
	"strconv"
	"time"

	"github.com/aftership/aftership-sdk-go/v2"
	"github.com/aftership/aftership-sdk-go/v2/common"
	"github.com/aftership/aftership-sdk-go/v2/endpoint/tracking"
)

func ExampleEndpoint_CreateTracking() {
	client, err := aftership.NewClient(&common.AfterShipConf{
		APIKey: "YOUR_API_KEY",
	})

	if err != nil {
		fmt.Println(err)
		return
	}

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

	result, err := client.Tracking.CreateTracking(newTracking)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}

func ExampleEndpoint_DeleteTracking() {
	client, err := aftership.NewClient(&common.AfterShipConf{
		APIKey: "YOUR_API_KEY",
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	// Delete a tracking
	param := common.SingleTrackingParam{
		Slug:           "dhl",
		TrackingNumber: "1234567890",
	}

	result, err := client.Tracking.DeleteTracking(param)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}

func ExampleEndpoint_GetTrackings() {
	client, err := aftership.NewClient(&common.AfterShipConf{
		APIKey: "YOUR_API_KEY",
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	// Get tracking results of multiple trackings.
	multiParams := tracking.MultiTrackingsParams{
		Page:  1,
		Limit: 10,
	}

	multiResults, err := client.Tracking.GetTrackings(multiParams)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(multiResults)
}

func ExampleEndpoint_GetTracking() {
	client, err := aftership.NewClient(&common.AfterShipConf{
		APIKey: "YOUR_API_KEY",
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	// Get tracking results of a single tracking.
	param := common.SingleTrackingParam{
		Slug:           "dhl",
		TrackingNumber: "1588226550",
	}

	result, err := client.Tracking.GetTracking(param, nil)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	// Get tracking results of a single tracking by id.
	param = common.SingleTrackingParam{
		ID: "rymq9l34ztbvvk9md2ync00r",
	}

	result, err = client.Tracking.GetTracking(param, &tracking.GetTrackingParams{
		Fields: "tracking_postal_code,title,order_id",
	})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

func ExampleEndpoint_UpdateTracking() {
	client, err := aftership.NewClient(&common.AfterShipConf{
		APIKey: "YOUR_API_KEY",
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	// Update a tracking.
	param := common.SingleTrackingParam{
		Slug:           "dhl",
		TrackingNumber: "1588226550",
	}

	updateReq := tracking.UpdateTrackingRequest{
		Tracking: tracking.UpdateTracking{
			Title: "New Title",
		},
	}

	result, err := client.Tracking.UpdateTracking(param, updateReq)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}

func ExampleEndpoint_ReTrack() {
	client, err := aftership.NewClient(&common.AfterShipConf{
		APIKey: "YOUR_API_KEY",
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	// Retrack an expired tracking.
	param := common.SingleTrackingParam{
		Slug:           "dhl",
		TrackingNumber: "1588226550",
	}

	result, err := client.Tracking.ReTrack(param)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
