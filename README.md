# Deprecation notice: New Version Available

This version of the SDK has been deprecated and replaced with the newly reconstructed SDK.

For the latest API features and improved integration, please visit our updated repository at https://github.com/AfterShip/tracking-sdk-go and follow the provided instructions.

# aftership-sdk-go

[![GoDoc](https://godoc.org/github.com/AfterShip/aftership-sdk-go?status.svg)](https://godoc.org/github.com/AfterShip/aftership-sdk-go)
[![AfterShip SDKs channel](https://aftership-sdk-slackin.herokuapp.com/badge.svg)](https://aftership-sdk-slackin.herokuapp.com/)

## Introduction

[AfterShip](https://aftership.com) provides an API to Track & Notify of shipments from hundreds of couriers worldwide. aftership-sdk-go is a SDK to develop Apps using [AfterShip API 2023-10](https://www.aftership.com/docs/tracking/quickstart/api-quick-start) in golang. All endpoints including couriers, tracking, last checkpoint and notification are supported.

You will need to create an account at [AfterShip](https://aftership.com) and obtain an API key first to access AfterShip APIs using aftership-go SDK.

## Installation

aftership-sdk-go requires a Go version with [Modules](https://github.com/golang/go/wiki/Modules) support and uses import versioning. So please make sure to initialize a Go module before installing aftership-sdk-go:

``` shell
go mod init github.com/my/repo
go get github.com/aftership/aftership-sdk-go/v3
```

Import:

``` go
import "github.com/aftership/aftership-sdk-go/v3"
```

## Quick Start

```go
package main

import (
        "context"
        "fmt"

        "github.com/aftership/aftership-sdk-go/v3"
)

func main() {

        client, err := aftership.NewClient(aftership.Config{
                APIKey: "YOUR_API_KEY",
        })

        if err != nil {
                fmt.Println(err)
                return
        }

        // Get couriers
        result, err := client.GetCouriers(context.Background())
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
  - `AuthenticationType` - `APIKey`  / `AES`
  - `APISecret` - if AuthenticationType is AES, use aes api secret
  - `Endpoint` - *string*, AfterShip endpoint, default 'https://api.aftership.com/tracking/2023-10'
  - `UserAagentPrefix` - *string*, prefix of User-Agent in headers, default "aftership-sdk-go"

Example:

AuthenticationType `APIKey`
```go
client, err := aftership.NewClient(aftership.Config{
    APIKey: "YOUR_API_KEY",
    UserAagentPrefix: "aftership-sdk-go",
})
```
AuthenticationType `AES` signature
```go
client, err := aftership.NewClient(aftership.Config{
    APIKey:             "YOUR_API_KEY",
    AuthenticationType: aftership.AES, 
    APISecret:          "YOUR_API_SECRET",
})
```

## Rate Limiter

To understand AfterShip rate limit policy, please see `Limit` section in https://www.aftership.com/docs/tracking/quickstart/rate-limit

You can get the recent rate limit by `client.GetRateLimit()`. Initially all value are `0`.

```go
import (
    "context"
    "fmt"

    "github.com/aftership/aftership-sdk-go/v3"
)

func main() {
    client, err := aftership.NewClient(aftership.Config{
        APIKey: "YOUR_API_KEY",
    })

    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(client.GetRateLimit())

    // terminal output
    /*
    {
        "reset": 0,
        "limit": 0,
        "remaining": 0,
    }
    */

    // Get couriers
    result, err := client.GetCouriers(context.Background())
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(result)
    }

    // Rate Limit
    fmt.Println(client.GetRateLimit())

    // terminal output
    /*
    {
        "reset": 1588249242,
        "limit": 10,
        "remaining": 9,
    }
    */
}
```

In case you exceeded the rate limit, you will receive the `429 Too Many Requests` error with the following error message:

```json
{
  "code": 429,
  "type": "TooManyRequests",
  "message": "You have exceeded the API call rate limit. Default limit is 10 requests per second.",
  "path": "/couriers",
  "rate_limit": {
    "rest": 1458463600,
    "limit": 10,
    "remaining": 0
  }
}
```

## Error Handling

There are 3 kinds of error

- SDK Error
- Request Error
- API Error

### SDK Error

**Throw** by the new SDK client

```go
client, err := aftership.NewClient(aftership.Config{
    APIKey: "",
})

if err != nil {
    fmt.Println(err)
    return
}

/*
invalid credentials: API Key must not be empty
*/
```

**Throw** by the parameter validation in function

```go
client, err := aftership.NewClient(aftership.Config{
    APIKey: "YOUR_API_KEY",
})

// Get notification
param := aftership.SlugTrackingNumber{
    Slug: "dhl",
}

result, err := client.GetNotification(context.Background(), param)
if err != nil {
    fmt.Println(err)
    return
}

fmt.Println(result)

/*
slug or tracking number is empty, both of them must be provided
*/
```

### Request Error

```go
client, err := aftership.NewClient(aftership.Config{
    APIKey: "YOUR_API_KEY",
})

// Get couriers
result, err := client.GetCouriers(context.Background())
if err != nil {
    fmt.Println(err)
    return
}

fmt.Println(result)
/*
HTTP request failed: Get https://api.aftership.com/v4/couriers: dial tcp: lookup api.aftership.com: no such host
*/
```

### API Error

Error return by the AfterShip API https://www.aftership.com/docs/tracking/quickstart/request-errors

API Error struct of this SDK contain fields:

- `Code` - error code for API Error
- `Type` - type of the error
- `Message` - detail message of the error
- `Path` - URI path when making request
- `RateLimit` - **Optional** - When the API gets `429 Too Many Requests` error, the error struct will return the `RateLimit` information as well.

```go
client, err := aftership.NewClient(aftership.Config{
    APIKey: "INVALID_API_KEY",
})

// Get couriers
result, err := client.GetCouriers(context.Background())
if err != nil {
    fmt.Println(err)
    return
}

fmt.Println(result)
/*
{
  "code": 401,
  "type": "Unauthorized",
  "message": "Invalid API key.",
  "path": "/couriers"
}
*/
```

## Examples

### /couriers

> Get a list of our supported couriers.

**GET** /couriers
> Return a list of couriers activated at your AfterShip account.

```go
result, err := client.GetCouriers(context.Background())
if err != nil {
    fmt.Println(err)
    return
}

fmt.Println(result)
```

**GET** /couriers/all
> Return a list of all couriers.

```go
result, err := client.GetAllCouriers(context.Background())
if err != nil {
    fmt.Println(err)
    return
}

fmt.Println(result)
```

**POST** /couriers/detect
> Return a list of matched couriers based on tracking number format and selected couriers or a list of couriers.

```go
params := aftership.CourierDetectionParams{
    TrackingNumber: "906587618687",
}

result, err := client.DetectCouriers(context.Background(), params)
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
newTracking := aftership.CreateTrackingParams{
    TrackingNumber: "1234567890",
    Slug:           "dhl",
    Title:          "Title Name",
    SMSes: []string{
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
}

result, err := client.CreateTracking(context.Background(), newTracking)
if err != nil {
    fmt.Println(err)
    return
}

fmt.Println(result)
```

**DELETE** /trackings/:slug/:tracking_number
> Delete a tracking.

```go
param := aftership.SlugTrackingNumber{
    Slug:           "dhl",
    TrackingNumber: "1234567890",
}

result, err := client.DeleteTracking(context.Background(), param)
if err != nil {
    fmt.Println(err)
    return
}

fmt.Println(result)
```

**GET** /trackings
> Get tracking results of multiple trackings.

```go
multiParams := aftership.GetTrackingsParams{
    Page:  1,
    Limit: 10,
}

result, err := client.GetTrackings(context.Background(), multiParams)
if err != nil {
    fmt.Println(err)
    return
}

fmt.Println(result)
```

**GET** /trackings/:slug/:tracking_number
> Get tracking results of a single tracking.

```go
param := aftership.SlugTrackingNumber{
    Slug:           "dhl",
    TrackingNumber: "1588226550",
}

result, err := client.GetTracking(context.Background(), param, aftership.GetTrackingParams{})
if err != nil {
    fmt.Println(err)
    return
}

fmt.Println(result)
```

> Pro Tip: You can always use /:id to replace /:slug/:tracking_number.

```go
// GET /trackings/:id
var id TrackingID = "5b7658cec7c33c0e007de3c5"

result, err := client.GetTracking(context.Background(), id, aftership.GetTrackingParams{})
if err != nil {
    fmt.Println(err)
    return
}

fmt.Println(result)
```

**PUT** /trackings/:slug/:tracking_number
> Update a tracking.

```go
param := aftership.SlugTrackingNumber{
    Slug:           "dhl",
    TrackingNumber: "1588226550",
}

updateReq := aftership.UpdateTrackingParams{
    Title: "New Title",
}

result, err := client.UpdateTracking(context.Background(), param, updateReq)
if err != nil {
    fmt.Println(err)
    return
}

fmt.Println(result)
```

**POST** /trackings/:slug/:tracking_number/retrack
> Retrack an expired tracking. Max 3 times per tracking.

```go
param := aftership.SlugTrackingNumber{
    Slug:           "dhl",
    TrackingNumber: "1588226550",
}

result, err := client.RetrackTracking(context.Background(), param)
if err != nil {
    fmt.Println(err)
    return
}

fmt.Println(result)
```

**POST** /trackings/:slug/:tracking_number/mark-as-completed
> Mark a tracking as completed. The tracking won't auto update until retrack it.

```go
param := aftership.SlugTrackingNumber{
    Slug:           "dhl",
    TrackingNumber: "1588226550",
}

result, err := client.MarkTrackingAsCompleted(context.Background(), param, aftership.TrackingCompletedStatusDelivered)
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
param := aftership.SlugTrackingNumber{
    Slug:           "ups",
    TrackingNumber: "1234567890",
}

result, err := client.GetLastCheckpoint(context.Background(), param, aftership.GetCheckpointParams{})
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
param := aftership.SlugTrackingNumber{
    Slug:           "dhl",
    TrackingNumber: "1588226550",
}

result, err := client.GetNotification(context.Background(), param)
if err != nil {
    fmt.Println(err)
    return
}

fmt.Println(result)
```

**POST** /notifications/:slug/:tracking_number/add
> Add notification receivers to a tracking number.

```go
param := aftership.SlugTrackingNumber{
    Slug:           "dhl",
    TrackingNumber: "1588226550",
}

data := notification.Data{
    Notification: notification.Notification{
        Emails: []string{"user1@gmail.com", "user2@gmail.com", "invalid EMail @ Gmail. com"},
        Smses:  []string{"+85291239123", "+85261236123", "Invalid Mobile Phone Number"},
    },
}

result, err := client.AddNotification(context.Background(), param, data)
if err != nil {
    fmt.Println(err)
    return
}

fmt.Println(result)
```

**POST** /notifications/:slug/:tracking_number/remove
> Remove notification receivers from a tracking number.

```go
param := aftership.SlugTrackingNumber{
    Slug:           "dhl",
    TrackingNumber: "1588226550",
}

data := notification.Data{
    Notification: notification.Notification{
        Emails: []string{"user1@gmail.com"},
        Smses:  []string{"+85291239123"},
    },
}

result, err := client.RemoveNotification(context.Background(), param, data)
if err != nil {
    fmt.Println(err)
    return
}

fmt.Println(result)
```

## Migrations

- `Checkpoint.Coordinates` change type from `[]string` into `[]float32`
- `Tracking` struct add fields
- remove `android` field from `Tracking` struct

## Help

If you get stuck, we're here to help. The following are the best ways to get assistance working through your issue:

- [Issue Tracker](https://github.com/AfterShip/aftership-sdk-go/issues) for questions, feature requests, bug reports and general discussion related to this package. Try searching before you create a new issue.
- [Slack AfterShip SDKs](https://aftership-sdk-slackin.herokuapp.com/): a Slack community, you can find the maintainers and users of this package in #aftership-sdks.
- [Email us](support@aftership.com) in AfterShip support: `support@aftership.com`

## Contributing

For details on contributing to this repository, see the [contributing guide](https://github.com/AfterShip/aftership-sdk-go/blob/master/CONTRIBUTING.md).
