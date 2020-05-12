package notification_test

import (
	"fmt"

	"github.com/aftership/aftership-sdk-go/v2"
	"github.com/aftership/aftership-sdk-go/v2/common"
	"github.com/aftership/aftership-sdk-go/v2/endpoint/notification"
)

func ExampleEndpoint_GetNotification() {
	client, err := aftership.NewClient(&common.AfterShipConf{
		APIKey: "YOUR_API_KEY",
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	// Get the notification
	param := common.SingleTrackingParam{
		Slug:           "dhl",
		TrackingNumber: "1588226550",
	}

	result, err := client.Notification.GetNotification(param)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}

func ExampleEndpoint_AddNotification() {
	client, err := aftership.NewClient(&common.AfterShipConf{
		APIKey: "YOUR_API_KEY",
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	// Add notification receivers to a tracking number.
	param := common.SingleTrackingParam{
		Slug:           "dhl",
		TrackingNumber: "1588226550",
	}

	data := notification.Data{
		Notification: notification.Notification{
			Emails: []string{"user1@gmail.com", "user2@gmail.com", "invalid EMail @ Gmail. com"},
			Smses:  []string{"+85291239123", "+85261236123", "Invalid Mobile Phone Number"},
		},
	}
	result, err := client.Notification.AddNotification(param, data)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}

func ExampleEndpoint_RemoveNotification() {
	client, err := aftership.NewClient(&common.AfterShipConf{
		APIKey: "YOUR_API_KEY",
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	// Remove notification receivers from a tracking number.
	param := common.SingleTrackingParam{
		Slug:           "dhl",
		TrackingNumber: "1588226550",
	}

	data := notification.Data{
		Notification: notification.Notification{
			Emails: []string{"user2@gmail.com"},
			Smses:  []string{"+85261236123"},
		},
	}
	result, err := client.Notification.RemoveNotification(param, data)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}
