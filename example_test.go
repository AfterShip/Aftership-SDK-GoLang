package aftership_test

import (
	"fmt"

	"github.com/aftership/aftership-sdk-go/v2"
	"github.com/aftership/aftership-sdk-go/v2/common"
)

func ExampleNewClient() {
	client, err := aftership.NewClient(&common.AfterShipConf{
		APIKey: "YOUR_API_KEY",
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(client)
}
