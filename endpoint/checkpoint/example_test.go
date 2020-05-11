package checkpoint_test

import (
	"fmt"

	"github.com/aftership/aftership-sdk-go/v2"
	"github.com/aftership/aftership-sdk-go/v2/common"
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
	param := common.SingleTrackingParam{
		Slug:           "ups",
		TrackingNumber: "1234567890",
	}

	result, err := client.LastCheckpoint.GetLastCheckpoint(param, nil)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}
