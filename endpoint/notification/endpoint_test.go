package notification

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/aftership/aftership-sdk-go/v2/common"
	"github.com/aftership/aftership-sdk-go/v2/endpoint/tracking"
	"github.com/aftership/aftership-sdk-go/v2/request"
	"github.com/aftership/aftership-sdk-go/v2/response"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestAddNotification(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	p := tracking.SingleTrackingParam{
		ID:             "",
		Slug:           "xq-express",
		TrackingNumber: "LS404494276CN",
		OptionalParams: nil,
	}

	exp := Data{
		Notification{
			[]string{"vimukthi@aftership.net"},
			[]string{"+85254469627"},
		},
	}
	mockhttp("POST", fmt.Sprintf("/notifications/%s/%s/add", p.Slug, p.TrackingNumber), 200, Envelope{
		response.Meta{
			Code:    200,
			Message: "",
			Type:    "",
		},
		exp,
	}, nil)

	req := request.NewRequest(&common.AfterShipConf{
		APIKey: "YOUR_API_KEY",
	}, nil)
	endpoint := NewEndpoint(req)
	res, _ := endpoint.AddNotification(context.Background(), p, exp)
	assert.Equal(t, exp, res)
}

func TestAddNotificationError(t *testing.T) {
	req := request.NewRequest(&common.AfterShipConf{
		APIKey: "YOUR_API_KEY",
	}, nil)
	endpoint := NewEndpoint(req)

	// empty id, slug and tracking_number
	p := tracking.SingleTrackingParam{
		ID:             "",
		Slug:           "",
		TrackingNumber: "",
		OptionalParams: nil,
	}

	_, err := endpoint.AddNotification(context.Background(), p, Data{})
	assert.NotNil(t, err)
	assert.Equal(t, "HandlerError", err.Type)

	// Response with error
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	p = tracking.SingleTrackingParam{
		ID:             "",
		Slug:           "xq-express",
		TrackingNumber: "LS404494276CN",
		OptionalParams: nil,
	}

	mockhttp("POST", fmt.Sprintf("/notifications/%s/%s/add", p.Slug, p.TrackingNumber), 401, Envelope{
		response.Meta{
			Code:    401,
			Message: "Invalid API key.",
			Type:    "Unauthorized",
		},
		Data{},
	}, nil)

	_, err = endpoint.AddNotification(context.Background(), p, Data{})
	assert.NotNil(t, err)
	assert.Equal(t, "Unauthorized", err.Type)
}

func TestRemoveNotification(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	p := tracking.SingleTrackingParam{
		ID:             "",
		Slug:           "xq-express",
		TrackingNumber: "LS404494276CN",
		OptionalParams: nil,
	}

	exp := Data{
		Notification{
			[]string{"vimukthi@aftership.net"},
			[]string{"+85254469627"},
		},
	}
	mockhttp("POST", fmt.Sprintf("/notifications/%s/%s/remove", p.Slug, p.TrackingNumber), 200, Envelope{
		response.Meta{
			Code:    200,
			Message: "",
			Type:    "",
		},
		exp,
	}, nil)

	req := request.NewRequest(&common.AfterShipConf{
		APIKey: "YOUR_API_KEY",
	}, nil)
	endpoint := NewEndpoint(req)
	res, _ := endpoint.RemoveNotification(context.Background(), p, exp)
	assert.Equal(t, exp, res)
}

func TestRemoveNotificationError(t *testing.T) {
	req := request.NewRequest(&common.AfterShipConf{
		APIKey: "YOUR_API_KEY",
	}, nil)
	endpoint := NewEndpoint(req)

	// empty id, slug and tracking_number
	p := tracking.SingleTrackingParam{
		ID:             "",
		Slug:           "",
		TrackingNumber: "",
		OptionalParams: nil,
	}

	_, err := endpoint.RemoveNotification(context.Background(), p, Data{})
	assert.NotNil(t, err)
	assert.Equal(t, "HandlerError", err.Type)

	// Response with error
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	p = tracking.SingleTrackingParam{
		ID:             "",
		Slug:           "xq-express",
		TrackingNumber: "LS404494276CN",
		OptionalParams: nil,
	}

	mockhttp("POST", fmt.Sprintf("/notifications/%s/%s/remove", p.Slug, p.TrackingNumber), 401, Envelope{
		response.Meta{
			Code:    401,
			Message: "Invalid API key.",
			Type:    "Unauthorized",
		},
		Data{},
	}, nil)

	_, err = endpoint.RemoveNotification(context.Background(), p, Data{})
	assert.NotNil(t, err)
	assert.Equal(t, "Unauthorized", err.Type)
}

func TestGetNotificationSetting(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	p := tracking.SingleTrackingParam{
		ID:             "",
		Slug:           "xq-express",
		TrackingNumber: "LS404494276CN",
		OptionalParams: nil,
	}

	exp := Data{
		Notification{
			[]string{"vimukthi@aftership.net"},
			[]string{"+85254469627"},
		},
	}
	mockhttp("GET", fmt.Sprintf("/notifications/%s/%s", p.Slug, p.TrackingNumber), 200, Envelope{
		response.Meta{
			Code:    200,
			Message: "",
			Type:    "",
		},
		exp,
	}, nil)

	req := request.NewRequest(&common.AfterShipConf{
		APIKey: "YOUR_API_KEY",
	}, nil)
	endpoint := NewEndpoint(req)
	res, _ := endpoint.GetNotification(context.Background(), p)
	assert.Equal(t, exp, res)
}

func TestGetotificationError(t *testing.T) {
	req := request.NewRequest(&common.AfterShipConf{
		APIKey: "YOUR_API_KEY",
	}, nil)
	endpoint := NewEndpoint(req)

	// empty id, slug and tracking_number
	p := tracking.SingleTrackingParam{
		ID:             "",
		Slug:           "",
		TrackingNumber: "",
		OptionalParams: nil,
	}

	_, err := endpoint.GetNotification(context.Background(), p)
	assert.NotNil(t, err)
	assert.Equal(t, "HandlerError", err.Type)

	// Response with error
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	p = tracking.SingleTrackingParam{
		ID:             "",
		Slug:           "xq-express",
		TrackingNumber: "LS404494276CN",
		OptionalParams: nil,
	}

	mockhttp("GET", fmt.Sprintf("/notifications/%s/%s", p.Slug, p.TrackingNumber), 401, Envelope{
		response.Meta{
			Code:    401,
			Message: "Invalid API key.",
			Type:    "Unauthorized",
		},
		Data{},
	}, nil)

	_, err = endpoint.GetNotification(context.Background(), p)
	assert.NotNil(t, err)
	assert.Equal(t, "Unauthorized", err.Type)
}

func mockhttp(method string, url string, status int, resp interface{}, headers map[string]string) {
	httpmock.RegisterResponder(method, url,
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(status, resp)
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
