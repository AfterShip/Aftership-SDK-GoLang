package courier

import (
	"net/http"
	"testing"

	"github.com/aftership/aftership-sdk-go/v2/common"
	"github.com/aftership/aftership-sdk-go/v2/request"
	"github.com/aftership/aftership-sdk-go/v2/response"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestGetCouriers(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	exp := List{
		Total: 1,
		Couriers: []Courier{
			{
				Slug: "ups",
				Name: "UPS",
			},
		},
	}
	mockhttp("GET", "/couriers", 200, Envelope{
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
	endpoint := NewEnpoint(req)
	res, err := endpoint.GetCouriers()
	assert.Equal(t, exp, res)
	assert.Nil(t, err)
}

func TestGetCouriersError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mockhttp("GET", "/couriers", 429, Envelope{
		response.Meta{
			Code:    429,
			Message: "You have exceeded the API call rate limit. Default limit is 10 requests per second.",
			Type:    "TooManyRequests",
		},
		List{},
	}, nil)

	req := request.NewRequest(&common.AfterShipConf{
		APIKey: "YOUR_API_KEY",
	}, nil)
	endpoint := NewEnpoint(req)
	_, err := endpoint.GetCouriers()
	assert.NotNil(t, err)
	assert.Equal(t, "TooManyRequests", err.Type)
}

func TestGetAllCouriers(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	exp := List{
		Total: 1,
		Couriers: []Courier{
			{
				Slug: "ups",
				Name: "ups",
			},
			{
				Slug: "fedex",
				Name: "FeDex",
			},
		},
	}
	mockhttp("GET", "/couriers/all", 200, Envelope{
		response.Meta{
			Code:    200,
			Message: "",
			Type:    "",
		},
		exp,
	}, map[string]string{
		"X-RateLimit-Reset":     "1458463600",
		"X-RateLimit-Limit":     "",
		"X-RateLimit-Remaining": "",
	})

	req := request.NewRequest(&common.AfterShipConf{
		APIKey: "YOUR_API_KEY",
	}, nil)
	endpoint := NewEnpoint(req)
	res, _ := endpoint.GetAllCouriers()
	assert.Equal(t, exp, res)
}

func TestGetAllCouriersError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mockhttp("GET", "/couriers/all", 429, Envelope{
		response.Meta{
			Code:    429,
			Message: "You have exceeded the API call rate limit. Default limit is 10 requests per second.",
			Type:    "TooManyRequests",
		},
		List{},
	}, nil)

	req := request.NewRequest(&common.AfterShipConf{
		APIKey: "YOUR_API_KEY",
	}, nil)
	endpoint := NewEnpoint(req)
	_, err := endpoint.GetAllCouriers()
	assert.NotNil(t, err)
	assert.Equal(t, "TooManyRequests", err.Type)
}

func TestDetectCouriers(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	exp := DetectList{
		Total: 2,
		Couriers: []Courier{
			{
				Slug: "ups",
				Name: "ups",
			},
		},
	}
	mockhttp("POST", "/couriers/detect", 200, DetectEnvelope{
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
	endpoint := NewEnpoint(req)
	result, _ := endpoint.DetectCouriers(DetectCourierRequest{
		Tracking: DetectParam{
			"906587618687",
			"DA15BU",
			"20131231",
			"1234567890",
			"",
			"",
			[]string{"dhl", "ups", "fedex"},
		},
	})
	assert.Equal(t, exp, result)
}

func TestInvalidDetectCouriers(t *testing.T) {
	req := request.NewRequest(&common.AfterShipConf{
		APIKey: "YOUR_API_KEY",
	}, nil)
	endpoint := NewEnpoint(req)
	_, err := endpoint.DetectCouriers(DetectCourierRequest{})

	assert.NotNil(t, err)
	assert.Equal(t, "HandlerError: Invalid TrackingNumber", err.Message)
}

func TestDetectCouriersError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mockhttp("POST", "/couriers/detect", 401, Envelope{
		response.Meta{
			Code:    402,
			Message: "Invalid API key.",
			Type:    "Unauthorized",
		},
		List{},
	}, nil)

	req := request.NewRequest(&common.AfterShipConf{}, nil)
	endpoint := NewEnpoint(req)
	_, err := endpoint.DetectCouriers(DetectCourierRequest{
		Tracking: DetectParam{
			TrackingNumber: "906587618687",
		},
	})
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
