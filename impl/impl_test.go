package impl

import (
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/vimukthi-git/aftership-go/apiV4"
	"net/http"
	"testing"
	//"time"
)

func TestMain(m *testing.M) {
	httpmock.Activate()
	m.Run()
	httpmock.DeactivateAndReset()
}

func TestGetCouriers(t *testing.T) {
	exp := []apiV4.Courier{
		apiV4.Courier{
			"ups",
			"ups",
			"ups",
			"ups",
			"ups",
			nil,
			"ups",
			nil,
			nil,
		},
	}
	mockhttp("GET", "https://api.aftership.com/v4/couriers", apiV4.CourierEnvelope{
		apiV4.ResponseMeta{200, "", ""},
		apiV4.CourierResponseData{exp},
	}, nil)
	var api apiV4.CourierHandler = &AfterShipApiV4Impl{
		"1eb4d44d-228c-4d2e-aa8d-53eeebe86c61",
		nil,
		nil,
	}

	res, _ := api.GetCouriers()
	assert.Equal(t, exp, res)
}

func TestGetAllCouriers(t *testing.T) {
	exp := []apiV4.Courier{
		apiV4.Courier{
			"ups",
			"ups",
			"ups",
			"ups",
			"ups",
			nil,
			"ups",
			nil,
			nil,
		},
	}
	mockhttp("GET", "https://api.aftership.com/v4/couriers/all", apiV4.CourierEnvelope{
		apiV4.ResponseMeta{200, "", ""},
		apiV4.CourierResponseData{exp},
	}, map[string]string{
		"X-RateLimit-Reset":"1458463600",
		"X-RateLimit-Limit": "",
		"X-RateLimit-Remaining": "",
	})

	var api apiV4.CourierHandler = &AfterShipApiV4Impl{
		"XXXX",
		nil,
		nil,
	}

	res, _ := api.GetAllCouriers()
	assert.Equal(t, exp, res)
}

func TestDetectCouriers(t *testing.T) {
	exp := []apiV4.Courier{
		apiV4.Courier{
			"ups",
			"ups",
			"ups",
			"ups",
			"ups",
			nil,
			"ups",
			nil,
			nil,
		},
	}
	mockhttp("GET", "https://api.aftership.com/v4/couriers/all", apiV4.CourierEnvelope{
		apiV4.ResponseMeta{200, "", ""},
		apiV4.CourierResponseData{exp},
	}, nil)
	var api apiV4.CourierHandler = &AfterShipApiV4Impl{
		"xxxx",
		nil,
		nil,
	}

	fmt.Println(api.DetectCouriers(apiV4.CourierDetectParam{
		"1Z1W824V0345129836",
		"",
		"",
		"",
		"",
		"",
		nil,
	}))

	fmt.Println(api.DetectCouriers(apiV4.CourierDetectParam{
		"906587618687",
		"DA15BU",
		"20131231",
		"1234567890",
		"",
		"",
		[]string{"dhl", "ups", "fedex"},
	}))
	assert.Equal(t, "", "")
}

func TestCreateTracking(t *testing.T) {
	var api apiV4.TrackingsHandler = &AfterShipApiV4Impl{
		"xxxx",
		nil,
		nil,
	}

	fmt.Println(api.CreateTracking(apiV4.NewTracking{
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
	}))
	assert.Equal(t, "", "")
}

func TestDeleteTracking(t *testing.T) {
	var api apiV4.TrackingsHandler = &AfterShipApiV4Impl{
		"xxxx",
		nil,
		nil,
	}

	api.DeleteTracking(apiV4.TrackingId{
		"",
		"ups",
		"1Z9999999999999998",
	})

	api.DeleteTracking(apiV4.TrackingId{
		"56ddabb0ccacfe6c0d76e40e",
		"",
		"",
	})
	assert.Equal(t, "", "")
}

func TestGetTrackings(t *testing.T) {
	p := apiV4.GetTrackingsParams{
		1,
		5,
		"",
		"ups",
		0,
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
	}
	var api apiV4.TrackingsHandler = &AfterShipApiV4Impl{
		"xxxx",
		nil,
		nil,
	}
	fmt.Println(api.GetTrackings(p))

	p = apiV4.GetTrackingsParams{
		0,
		0,
		"",
		"",
		0,
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
	}
	fmt.Println(api.GetTrackings(p))
	assert.Equal(t, "", "")
}

func TestGetTrackingsExport(t *testing.T) {
	p := apiV4.GetTrackingsParams{
		1,
		5,
		"",
		"ups",
		0,
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
	}
	var api apiV4.TrackingsHandler = &AfterShipApiV4Impl{
		"xxxx",
		nil,
		nil,
	}
	fmt.Println(api.GetTrackingsExport(p))
	assert.Equal(t, "", "")
}

func TestGetTracking(t *testing.T) {
	p := apiV4.TrackingId{
		"",
		"xq-express",
		"LS404494276CN",
	}
	var api apiV4.TrackingsHandler = &AfterShipApiV4Impl{
		"xxxx",
		nil,
		nil,
	}
	fmt.Println(api.GetTracking(p, "", ""))

	p = apiV4.TrackingId{
		"",
		"ceska-posta",
		"RR8480201064M",
	}
	fmt.Println(api.GetTracking(p, "checkpoints", ""))
	assert.Equal(t, "", "")
}

func TestUpdateTracking(t *testing.T) {
	p := apiV4.TrackingId{
		"",
		"ups",
		"1Z9999999999999998",
	}
	var api apiV4.TrackingsHandler = &AfterShipApiV4Impl{
		"xxxx",
		nil,
		nil,
	}
	u := apiV4.TrackingUpdate{
		nil,
		nil,
		"Bullshit",
		"",
		"",
		"",
		"",
		nil,
	}
	fmt.Println(api.UpdateTracking(p, u))
	assert.Equal(t, "", "")
}

func TestReTrack(t *testing.T) {
	p := apiV4.TrackingId{
		"",
		"ups",
		"1Z9999999999999998",
	}
	var api apiV4.TrackingsHandler = &AfterShipApiV4Impl{
		"xxxx",
		nil,
		nil,
	}
	fmt.Println(api.ReTrack(p))
	assert.Equal(t, "", "")
}

func TestGetLastCheckpoint(t *testing.T) {
	p := apiV4.TrackingId{
		"",
		"xq-express",
		"LS404494276CN",
	}
	var api apiV4.TrackingsHandler = &AfterShipApiV4Impl{
		"xxxx",
		nil,
		nil,
	}
	fmt.Println(api.GetLastCheckPoint(p, "", ""))

	p = apiV4.TrackingId{
		"",
		"ceska-posta",
		"RR8480201064M",
	}
	fmt.Println(api.GetLastCheckPoint(p, "checkpoint_time", ""))
	assert.Equal(t, "", "")
}

func TestAddNotification(t *testing.T) {
	p := apiV4.TrackingId{
		"",
		"xq-express",
		"LS404494276CN",
	}

	n := apiV4.NotificationSetting{
		nil,
		[]string{"vimukthi@aftership.net"},
		nil,
		[]string{"+85254469627"},
	}
	var api apiV4.NotificationsHandler = &AfterShipApiV4Impl{
		"xxxx",
		nil,
		nil,
	}
	fmt.Println(api.AddNotification(p, n))
	assert.Equal(t, "", "")
}

func TestRemoveNotification(t *testing.T) {
	p := apiV4.TrackingId{
		"",
		"xq-express",
		"LS404494276CN",
	}

	n := apiV4.NotificationSetting{
		nil,
		[]string{"vimukthi@aftership.net"},
		nil,
		[]string{"+85254469627"},
	}
	var api apiV4.NotificationsHandler = &AfterShipApiV4Impl{
		"xxxx",
		nil,
		nil,
	}
	// fmt.Println(api.RemoveNotification(p, n))
	res, _ := api.RemoveNotification(p, n)
	assert.Equal(t, "", res)
}

func TestGetNotificationSetting(t *testing.T) {
	expect := apiV4.NotificationSetting{Android: []string{}, Emails: []string{"vb"}, Ios: []string{"dewdewfe"}, Smses: []string{}}
	resp := apiV4.NotificationSettingEnvelope{
		apiV4.ResponseMeta{200, "", ""},
		apiV4.NotificationSettingWrapper{expect},
	}
	mockhttp("GET", "https://api.aftership.com/v4/notifications/xq-express/LS404494276CN", resp, nil)

	var api apiV4.NotificationsHandler = &AfterShipApiV4Impl{
		"XXXXXXX",
		nil,
		nil,
	}
	//fmt.Println(api.GetNotificationSetting(p, ""))
	res, _ := api.GetNotificationSetting(
		apiV4.TrackingId{
			"",
			"xq-express",
			"LS404494276CN",
		}, "")
	assert.Equal(t, expect, res)
}

func mockhttp(method string, url string, resp interface{}, headers map[string]string) {
	httpmock.RegisterResponder(method, url,
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(200, resp)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			for key, value := range headers {
				resp.Header.Set(key, value)
			}
			return resp, nil
		},
	)
}

func rateLimitCheck() {
	//	p := apiV4.GetTrackingsParams{
	//		1,
	//		5,
	//		"",
	//		"ups",
	//		0,
	//		"",
	//		"",
	//		"",
	//		"",
	//		"",
	//		"",
	//		"",
	//		"",
	//	}
	var api apiV4.CourierHandler = &AfterShipApiV4Impl{
		"1eb4d44d-228c-4d2e-aa8d-53eeebe86c61",
		&apiV4.RetryPolicy{
			false,
			0,
			false,
		},
		nil,
	}


	c := make(chan apiV4.AfterShipApiError)
	for i := 0; i < 100; i++ {

		go func() {
			//var meta apiV4.AfterShipApiError
			for i := 0; i < 10000; i++ {
				_, meta := api.GetCouriers();
				c <- meta
				// fmt.Println(res)
			}


		}()


	}
	for {

		<-c

	}


	//
	//	p = apiV4.GetTrackingsParams{
	//		0,
	//		0,
	//		"",
	//		"",
	//		0,
	//		"",
	//		"",
	//		"",
	//		"",
	//		"",
	//		"",
	//		"",
	//		"",
	//	}
	//	fmt.Println(api.GetTrackings(p))
	//	assert.Equal(t, "", "")
}
