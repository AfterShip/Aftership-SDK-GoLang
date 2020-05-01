package request

import (
	"net/http"
	"testing"

	"github.com/aftership/aftership-sdk-go/v2/common"
	"github.com/aftership/aftership-sdk-go/v2/response"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

type mockData struct {
	Meta response.Meta `json:"meta"`
	Data string        `json:"data"`
}

func (e *mockData) GetMeta() response.Meta {
	return e.Meta
}

func TestMakeRequest(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	exp := mockData{
		Data: "test",
	}

	// GET with status 200
	mockhttp("GET", "/test", 200, exp, nil)

	req := NewRequest(&common.AfterShipConf{
		APIKey: "YOUR_API_KEY",
	}, nil)

	var result mockData
	err := req.MakeRequest("GET", "/test", nil, &result)
	assert.Nil(t, err)
	assert.Equal(t, exp, result)

	// POST with status 201
	mockhttp("POST", "/test", 201, exp, nil)

	err = req.MakeRequest("POST", "/test", nil, &result)
	assert.Nil(t, err)
	assert.Equal(t, exp, result)

	// PUT with status 200
	mockhttp("POST", "/test", 200, exp, nil)

	err = req.MakeRequest("POST", "/test", exp, &result)
	assert.Nil(t, err)
	assert.Equal(t, exp, result)
}

func TestMakeRequestError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mockhttp("GET", "/test", 500, nil, nil)

	req := NewRequest(&common.AfterShipConf{
		APIKey: "YOUR_API_KEY",
	}, nil)

	var result mockData
	err := req.MakeRequest("GET", "/test", nil, &result)
	assert.NotNil(t, err)
	assert.Equal(t, "InternalError", err.Type)
}

func TestRateLimit(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	exp := mockData{
		Meta: response.Meta{
			Code:    429,
			Type:    "TooManyRequests",
			Message: "You have exceeded the API call rate limit. Default limit is 10 requests per second.",
		},
		Data: "test",
	}
	mockhttp("GET", "/test", 429, exp, map[string]string{
		"x-ratelimit-reset":     "1458463600",
		"x-ratelimit-limit":     "10",
		"x-ratelimit-remaining": "9",
	})

	rateLimit := &response.RateLimit{}
	req := NewRequest(&common.AfterShipConf{
		APIKey: "YOUR_API_KEY",
	}, rateLimit)

	var result mockData
	err := req.MakeRequest("GET", "/test", nil, &result)
	assert.NotNil(t, err)
	assert.Equal(t, "TooManyRequests", err.Type)
	assert.Equal(t, int64(1458463600), rateLimit.Reset)
	assert.Equal(t, 10, rateLimit.Limit)
	assert.Equal(t, 9, rateLimit.Remaining)
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
