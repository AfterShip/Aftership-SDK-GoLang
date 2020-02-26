![Picture](https://www.aftership.com/assets/common/img/logo-aftership-premium-bright.png)

Aftership-SDK-GoLang 
============

### Introduction:

[Aftership](https://aftership.com) provides an API to Track & Notify of shipments from hundreds of couriers worldwide. Aftership-SDK-GoLang is a SDK to develop Apps using Aftership API in go-lang. All endpoints including couriers, tracking and notification are supported.

Contact: <support@aftership.com>

You will need to create an account at [Aftership](https://aftership.com) and obtain an API key first to access Aftership APIs using aftership-go SDK..

### Installation
If you already have go installed locally then do,
````
go get github.com/vimukthi-git/aftership-go/apiV4
go get github.com/vimukthi-git/aftership-go/impl
````
### Example:

- Getting couriers already added to your account,

```go
package main

import (
        "fmt"
        "github.com/vimukthi-git/aftership-go/apiV4"
        "github.com/vimukthi-git/aftership-go/impl"
)

func main() {
        var api apiV4.CourierHandler = &impl.AfterShipApiV4Impl{
                "<your-api-key>",
                nil,
                nil,
        }
        res, meta := api.GetCouriers()
        if (meta.Code == 200) {
            fmt.Print(res)
        }
        
}

```

- Posting a tracking to the API,

````go
package main

import (
        "fmt"
        "github.com/vimukthi-git/aftership-go/apiV4"
        "github.com/vimukthi-git/aftership-go/impl"
)

func main() {
        var api apiV4.TrackingsHandler = &impl.AfterShipApiV4Impl{
                "<your-api-key>",
                nil,
                nil,
        }
        res, meta := api.CreateTracking(apiV4.NewTracking{
                "1Z9999999999999998",
                nil,
                "",
                "",
                "",
                "",
                "",
                nil,
                nil,
                nil,
                nil,
                "",
                "",
                "",
                "",
                "",
                nil,
        })
        
        if (meta.Code == 200) {
            fmt.Print(res)
        }
}

````

- Getting Checkpoints and other tracking information for a tracking number,

````go
package main

import (
        "fmt"
        "github.com/vimukthi-git/aftership-go/apiV4"
        "github.com/vimukthi-git/aftership-go/impl"
)

func main() {
        var api apiV4.TrackingsHandler = &impl.AfterShipApiV4Impl{
                "<your-api-key>",
                nil,
                nil,
        }
        res, meta := api.GetTracking(
            apiV4.TrackingId{
                "",
                "xq-express",
                "LS404494276CN",
            }, 
            "",
            "",
        )
        
        if (meta.Code == 200) {
            fmt.Print(res)
        }
}

````

- Setting a proxy through which the SDK will make requests is easy since we expose `http.Client` to the user. Eg: 

````
proxyUrl, err := url.Parse("http://proxyIp:proxyPort")
myClient := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}

// Instantiate sdk with your own client.
var api apiV4.TrackingsHandler = &impl.AfterShipApiV4Impl{
                "<your-api-key>",
                nil,
                myClient,
        }
````

Check `./impl/impl_test.go` for examples on using all endpoints.

### Documentation

- [API Version 4](https://godoc.org/github.com/vimukthi-git/Aftership-SDK-GoLang/apiV4)
- [API Version 4 Implementation](https://godoc.org/github.com/vimukthi-git/Aftership-SDK-GoLang/impl)
