package courier_test

import (
	"context"
	"fmt"

	"github.com/aftership/aftership-sdk-go/v2"
	"github.com/aftership/aftership-sdk-go/v2/common"
	"github.com/aftership/aftership-sdk-go/v2/endpoint/courier"
)

func ExampleEndpoint_GetCouriers() {
	client, err := aftership.NewClient(&common.AfterShipConf{
		APIKey: "YOUR_API_KEY",
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	// Get couriers
	result, err := client.Courier.GetCouriers(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}

func ExampleEndpoint_GetAllCouriers() {
	client, err := aftership.NewClient(&common.AfterShipConf{
		APIKey: "YOUR_API_KEY",
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	// Get all couriers
	result, err := client.Courier.GetAllCouriers(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}

func ExampleEndpoint_DetectCouriers() {
	client, err := aftership.NewClient(&common.AfterShipConf{
		APIKey: "YOUR_API_KEY",
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	// Detect courier
	req := courier.DetectCourierRequest{
		Tracking: courier.DetectParam{
			TrackingNumber: "906587618687",
		},
	}

	list, err := client.Courier.DetectCouriers(context.Background(), req)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(list)
}
