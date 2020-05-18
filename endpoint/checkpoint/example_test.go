package checkpoint_test

import (
	"context"
	"fmt"

	"github.com/aftership/aftership-sdk-go/v2"
	"github.com/aftership/aftership-sdk-go/v2/common"
	"github.com/aftership/aftership-sdk-go/v2/endpoint/tracking"
)

func Example_getLastCheckpoint() {
	client, err := aftership.NewClient(&common.AfterShipConf{
		APIKey: "YOUR_API_KEY",
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	// Get last checkpopint
	param := tracking.SingleTrackingParam{
		Slug:           "ups",
		TrackingNumber: "1234567890",
	}

	result, err := client.LastCheckpoint.GetLastCheckpoint(context.Background(), param, nil)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}
