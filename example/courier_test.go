package example

import (
	"fmt"
	"testing"

	"github.com/aftership/aftership-sdk-go/v2"
	"github.com/aftership/aftership-sdk-go/v2/common"
	"github.com/aftership/aftership-sdk-go/v2/courier"
)

func TestCourierExample(t *testing.T) {
	client, err := aftership.NewClient(&common.AfterShipConf{
		APIKey: "YOUR_API_KEY",
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	// Get couriers
	result, err := aftership.Courier.GetCouriers()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	// Rate Limit
	fmt.Println(aftership.RateLimit)

	// Get all couriers
	result, err = aftership.Courier.GetAllCouriers()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	// Detect courier
	req := courier.DetectCourierRequest{
		Tracking: courier.DetectParam{
			TrackingNumber: "906587618687",
		},
	}

	list, err := aftership.Courier.DetectCouriers(req)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(list)
	}
}
