
# aftership-sdk-go 

[![Build Status](https://travis-ci.org/AfterShip/aftership-sdk-go.svg?branch=v2)](https://travis-ci.org/AfterShip/aftership-sdk-go)
[![codecov.io](https://codecov.io/github/AfterShip/aftership-sdk-go/coverage.svg?branch=v2)](https://codecov.io/github/AfterShip/aftership-sdk-go?branch=v2)
[![GoDoc](https://godoc.org/github.com/AfterShip/aftership-sdk-go?status.svg)](https://godoc.org/github.com/AfterShip/aftership-sdk-go)
[![AfterShip SDKs channel](https://aftership-sdk-slackin.herokuapp.com/badge.svg)](https://aftership-sdk-slackin.herokuapp.com/)

## Introduction:

[AfterShip](https://aftership.com) provides an API to Track & Notify of shipments from hundreds of couriers worldwide. aftership-sdk-go is a SDK to develop Apps using [AfterShip API v4](https://docs.aftership.com/api/4) in golang. All endpoints including couriers, tracking, last checkpoint and notification are supported.

You will need to create an account at [AfterShip](https://aftership.com) and obtain an API key first to access AfterShip APIs using aftership-go SDK.

## Installation

aftership-sdk-go requires a Go version with [Modules](https://github.com/golang/go/wiki/Modules) support and uses import versioning. So please make sure to initialize a Go module before installing aftership-sdk-go:

``` shell
go mod init github.com/my/repo
go get github.com/aftership/aftership-sdk-go/v2
```

Import:

``` go
import "github.com/aftership/aftership-sdk-go/v2"
```

## Quick Start

```go
package main

import (
        "context"
        "fmt"

        "github.com/aftership/aftership-sdk-go/v2"
        "github.com/aftership/aftership-sdk-go/v2/common"
)

func main() {

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

```

## Test

```shell
make test
```

## Table of contents

- [NewClient(config)](#newaftershipconfig)
- [Endpoints](#endpoints)
- [Rate Limiter](#rate-limiter)
- [Error Handling](#error-handling)
- [Examples](#examples)
  - [/couriers](#couriers)
  - [/trackings](#trackings)
  - [/last_checkpoint](#last_checkpoint)
  - [/notifications](#notifications)
- [Migrations](#migrations)
- [Help](#help)
- [Contributing](#contributing)

## NewClient(config)

Create AfterShip SDK instance with config

- `config` - object of request config
  - `APIKey` - **Required**, AfterShip API key
  - `Endpoint` - *string*, AfterShip endpoint, default "https://api.aftership.com/v4"
  - `UserAagentPrefix` - *string*, prefix of User-Agent in headers, default "aftership-sdk-go"

Example:

```go
client, err := aftership.NewClient(&common.AfterShipConf{
    APIKey: "YOUR_API_KEY",
    Endpoint: "https://api.aftership.com/OLDER_VERSIONOUR_API_KEY",
    UserAagentPrefix: "aftership-sdk-go",
})
```

## Endpoints

The AfterShip instance has the following properties which are exactly the same as the API endpoins:

- `Courier` - Get a list of our supported couriers.
- `Tracking` - Create trackings, update trackings, and get tracking results.
- `LastCheckpoint` - Get tracking information of the last checkpoint of a tracking.
- `Notification` - Get, add or remove contacts (sms or email) to be notified when the status of a tracking has changed.

Make request in a specific endpoint

```go
// GET /trackings/:slug/:tracking_number
param := tracking.SingleTrackingParam{
    Slug:           "dhl",
    TrackingNumber: "1588226550",
}

result, err := client.Tracking.GetTracking(context.Background(), param, nil)
if err != nil {
    fmt.Println(err)
    return
}

fmt.Println(result)
```

## Rate Limiter

To understand AfterShip rate limit policy, please see `Limit` section in https://www.aftership.com/docs/api/4

You can get the recent rate limit by `client.RateLimit`. Initially all value are `0`.

```go
import (
    "context"
    "fmt"

    "github.com/aftership/aftership-sdk-go/v2"
    "github.com/aftership/aftership-sdk-go/v2/common"
)

func main() {
    client, err := aftership.NewClient(&common.AfterShipConf{
        APIKey: "YOUR_API_KEY",
    })

    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(client.RateLimit)

    // terminal output
    /*
    {
        Reset: 0,
        Limit: 0,
        Remaining: 0,
    }
    */

    // Get couriers
    result, err := client.Courier.GetCouriers(context.Background())
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(result)
    }

    // Rate Limit
    fmt.Println(client.RateLimit)

    // terminal output
    /*
    {
        Reset: 1588249242,
        Limit: 10,
        Remaining: 9,
    }
    */
}
```

## Error Handling

There are 3 kinds of error

- SDK Error
- Request Error
- API Error

Error object of this SDK contain fields:

- `Type` - **Require** - type of the error, **please handle each error by this field**
- `Code` - **Optional** - error code for API Error
- `Message` - **Optional** - detail message of the error
- `Data` - **Optional** - data lead to the error

> Please handle each error by its `type`, since it is a require field

### SDK Error

Error return by the SDK instance, mostly invalid param type when calling `constructor` or `endpoint method`  
`error.Type` is one of [error_enum](https://github.com/AfterShip/aftership-sdk-go/blob/master/src/v2/error/error_enum.go)  
**Throw** by the SDK instance

```go
client, err := aftership.NewClient(&common.AfterShipConf{
    APIKey: "",
})

if err != nil {
    fmt.Println(err)
    return
}

/*
{
  Type: "ConstructorError",
  Code: 0,
  Message: "ConstructorError: Invalid API key",
  Data: { },
}
*/
```

**Throw** by endpoint method

```go
client, err := aftership.NewClient(&common.AfterShipConf{
    APIKey: "YOUR_API_KEY",
})

// Get notification
param := tracking.SingleTrackingParam{
    Slug: "dhl",
}

result, err := client.Notification.GetNotification(context.Background(), param)
if err != nil {
    fmt.Println(err)
    return
}

fmt.Println(result)

/*
{
  Type: "HandlerError",
  Code: 0,
  Message: "You must specify the id or slug and tracking number",
  Data: { dhl  <nil>},
}
*/
```

### Request Error

Error return by the `request` module  
`error.Type` could be `RequestError`, etc.  
**Catch** by promise

```go
client, err := aftership.NewClient(&common.AfterShipConf{
    APIKey: "YOUR_API_KEY",
})

// Get couriers
result, err := client.Courier.GetCouriers(context.Background())
if err != nil {
    fmt.Println(err)
    return
}

fmt.Println(result)
/*
{
    Type: "RequestError",
    Message: "Get https://api.aftership.com/v4/couriers: dial tcp: lookup api.aftership.com: no such host",
    .....
}
*/
```

### API Error

Error return by the AfterShip API  
`error.Type` should be the same as https://www.aftership.com/docs/api/4/errors

```go
client, err := aftership.NewClient(&common.AfterShipConf{
    APIKey: "YOUR_API_KEY",
})

// Get couriers
result, err := client.Courier.GetCouriers(context.Background())
if err != nil {
    fmt.Println(err)
    return
}

fmt.Println(result)
/*
{
  Type: 'Unauthorized',
  Code: 401,
  Message: 'Invalid API key.',
  Data: <nil>,
}
*/
```

## Examples

### /couriers

> Get a list of our supported couriers.

**GET** /couriers
> Return a list of couriers activated at your AfterShip account.

```go
result, err := client.Courier.GetCouriers(context.Background())
if err != nil {
    fmt.Println(err)
    return
}

fmt.Println(result)
```

**GET** /couriers/all
> Return a list of all couriers.

```go
result, err := client.Courier.GetAllCouriers(context.Background())
if err != nil {
    fmt.Println(err)
    return
}

fmt.Println(result)
```

**POST** /couriers/detect
> Return a list of matched couriers based on tracking number format and selected couriers or a list of couriers.

```go
req := courier.DetectCourierRequest{
    Tracking: courier.DetectParam{
        TrackingNumber: "906587618687",
    },
}

result, err := client.Courier.DetectCouriers(context.Background(), req)
if err != nil {
    fmt.Println(err)
    return
}

fmt.Println(result)
```

### /trackings

> Create trackings, update trackings, and get tracking results.

**POST** /trackings
> Create a tracking.

```go
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

result, err := client.Tracking.CreateTracking(context.Background(), newTracking)
if err != nil {
    fmt.Println(err)
    return
}

fmt.Println(result)
```

**DELETE** /trackings/:slug/:tracking_number
> Delete a tracking.

```go
param := tracking.SingleTrackingParam{
   Slug:           "dhl",
   TrackingNumber: "1234567890",
}

result, err := client.Tracking.DeleteTracking(context.Background(), param)
if err != nil {
    fmt.Println(err)
    return
}

fmt.Println(result)
```

**GET** /trackings
> Get tracking results of multiple trackings.

```go
multiParams := tracking.MultiTrackingsParams{
    Page:  1,
    Limit: 10,
}

result, err := client.Tracking.GetTrackings(context.Background(), multiParams)
if err != nil {
    fmt.Println(err)
    return
}

fmt.Println(result)
```

**GET** /trackings/:slug/:tracking_number
> Get tracking results of a single tracking.

```go
param := tracking.SingleTrackingParam{
    Slug:           "dhl",
    TrackingNumber: "1588226550",
}

result, err := client.Tracking.GetTracking(context.Background(), param, nil)
if err != nil {
    fmt.Println(err)
    return
}

fmt.Println(result)
```

Tip: You can also add `OptionalParams` to `/:slug/:tracking_number`

```go
// GET /trackings/:slug/:tracking_number?tracking_postal_code=:postal_code&tracking_ship_date=:ship_date
param := tracking.SingleTrackingParam{
    Slug:           "dhl",
    TrackingNumber: "1588226550",
    OptionalParams: &common.SingleTrackingOptionalParams{
       TrackingPostalCode: "1234",
       TrackingShipDate: "20200420",
    },
}
```

> Pro Tip: You can always use /:id to replace /:slug/:tracking_number.

```go
// GET /trackings/:id
param := tracking.SingleTrackingParam{
    ID: "1234567890",
}

result, err := client.Tracking.GetTracking(context.Background(), param, nil)
if err != nil {
    fmt.Println(err)
    return
}

fmt.Println(result)
```

**PUT** /trackings/:slug/:tracking_number
> Update a tracking.

```go
param := tracking.SingleTrackingParam{
    Slug:           "dhl",
    TrackingNumber: "1588226550",
}

updateReq := tracking.UpdateTrackingRequest{
    Tracking: tracking.UpdateTracking{
        Title: "New Title",
    },
}

result, err := client.Tracking.UpdateTracking(context.Background(), param, updateReq)
if err != nil {
    fmt.Println(err)
    return
}

fmt.Println(result)
```

**POST** /trackings/:slug/:tracking_number/retrack
> Retrack an expired tracking. Max 3 times per tracking.

```go
param := tracking.SingleTrackingParam{
    Slug:           "dhl",
    TrackingNumber: "1588226550",
}

result, err := client.Tracking.ReTrack(context.Background(), param)
if err != nil {
    fmt.Println(err)
    return
}

fmt.Println(result)
```

**POST** /trackings/:slug/:tracking_number/mark-as-completed
> Mark a tracking as completed. The tracking won't auto update until retrack it.

```go
param := tracking.SingleTrackingParam{
    Slug:           "dhl",
    TrackingNumber: "1588226550",
}

reason := tracking.MarkAsCompletedRequest{
    Reason: "DELIVERED",
}

result, err := client.Tracking.MarkAsCompleted(context.Background(), param, reason)
if err != nil {
    fmt.Println(err)
    return
}

fmt.Println(result)
```

### /last_checkpoint

> Get tracking information of the last checkpoint of a tracking.

**GET** /last_checkpoint/:slug/:tracking_number
> Return the tracking information of the last checkpoint of a single tracking.

```go
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
```

### /notifications

> Get, add or remove contacts (sms or email) to be notified when the status of a tracking has changed.

**GET** /notifications/:slug/:tracking_number
> Get contact information for the users to notify when the tracking changes. 

```go
param := tracking.SingleTrackingParam{
    Slug:           "dhl",
    TrackingNumber: "1588226550",
}

result, err := client.Notification.GetNotification(context.Background(), param)
if err != nil {
    fmt.Println(err)
    return
}

fmt.Println(result)
```

**POST** /notifications/:slug/:tracking_number/add
> Add notification receivers to a tracking number.

```go
param := tracking.SingleTrackingParam{
    Slug:           "dhl",
    TrackingNumber: "1588226550",
}

data := notification.Data{
    Notification: notification.Notification{
        Emails: []string{"user1@gmail.com", "user2@gmail.com", "invalid EMail @ Gmail. com"},
        Smses:  []string{"+85291239123", "+85261236123", "Invalid Mobile Phone Number"},
    },
}

result, err := client.Notification.AddNotification(context.Background(), param, data)
if err != nil {
    fmt.Println(err)
    return
}

fmt.Println(result)
```

**POST** /notifications/:slug/:tracking_number/remove
> Remove notification receivers from a tracking number.

```go
param := tracking.SingleTrackingParam{
    Slug:           "dhl",
    TrackingNumber: "1588226550",
}

data := notification.Data{
    Notification: notification.Notification{
        Emails: []string{"user1@gmail.com"},
        Smses:  []string{"+85291239123"},
    },
}

result, err := client.Notification.RemoveNotification(context.Background(), param, data)
if err != nil {
    fmt.Println(err)
    return
}

fmt.Println(result)
```

## Migrations

```go
// old version
var api apiV4.CourierHandler = &impl.AfterShipApiV4Impl{
    "<your-api-key>",
    nil,
    nil,
}
res, meta := api.GetCouriers()
if (meta.Code == 200) {
    fmt.Print(res)
}

// new version (v2)
client, err := aftership.NewClient(&common.AfterShipConf{
    APIKey: "YOUR_API_KEY",
})

result, err := client.Courier.GetCouriers(context.Background())
if err != nil {
    fmt.Println(err)
    return
}

fmt.Println(result)
```

## Help

If you get stuck, we're here to help. The following are the best ways to get assistance working through your issue:

- [Issue Tracker](https://github.com/AfterShip/aftership-sdk-go/issues) for questions, feature requests, bug reports and general discussion related to this package. Try searching before you create a new issue.
- [Slack AfterShip SDKs](https://aftership-sdk-slackin.herokuapp.com/): a Slack community, you can find the maintainers and users of this package in #aftership-sdks.
- [Email us](support@aftership.com) in AfterShip support: `support@aftership.com`

## Contributing

For details on contributing to this repository, see the [contributing guide](https://github.com/AfterShip/aftership-sdk-go/blob/master/CONTRIBUTING.md).
