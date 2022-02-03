package aftership

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type mockData struct {
	Meta Meta   `json:"meta"`
	Data string `json:"data"`
}

func TestMakeGETRequest(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		w.Write([]byte(`{
			"meta": {
					"code": 200
			},
			"data": "test"
	}`))
	})

	exp := mockData{
		Data: "test",
	}

	var result string
	// GET with status 200
	err := client.makeRequest(context.Background(), http.MethodGet, "/test", nil, nil, &result)
	assert.Nil(t, err)
	assert.Equal(t, exp.Data, result)
}

func TestMakePOSTRequest(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		w.Write([]byte(`{
			"meta": {
					"code": 200
			},
			"data": "test"
	}`))
	})

	exp := mockData{
		Data: "test",
	}

	var result string
	// POST with status 201
	err := client.makeRequest(context.Background(), http.MethodPost, "/test", nil, nil, &result)
	assert.Nil(t, err)
	assert.Equal(t, exp.Data, result)
}

func TestMakePUTRequest(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method)
		w.Write([]byte(`{
			"meta": {
					"code": 200
			},
			"data": "test"
	}`))
	})

	exp := mockData{
		Data: "test",
	}

	var result string
	// PUT with status 200
	err := client.makeRequest(context.Background(), http.MethodPut, "/test", nil, nil, &result)
	assert.Nil(t, err)
	assert.Equal(t, exp.Data, result)
}

func TestMakeRequestError(t *testing.T) {
	setup()

	var result string
	// Invalid data
	err := client.makeRequest(context.Background(), http.MethodGet, "/test", nil, make(chan int), &result)
	assert.NotNil(t, err)

	// Bad method
	err = client.makeRequest(context.Background(), "bad method", "/test", nil, nil, &result)
	assert.NotNil(t, err)

	// Invalid query params
	err = client.makeRequest(context.Background(), http.MethodGet, "/test", 1, nil, &result)
	assert.NotNil(t, err)

	// 500 err
	mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		w.WriteHeader(500)
		w.Write([]byte(`{
			"meta": {
					"code": 500,
					"type": "InternalError",
					"message": "Something went wrong on AfterShip's end."
			},
			"data": "test"
		}`))
	})

	apiErr := APIError{
		Code:    500,
		Type:    "InternalError",
		Message: "Something went wrong on AfterShip's end.",
		Path:    "/test",
	}
	exp, _ := json.Marshal(apiErr)

	err = client.makeRequest(context.Background(), http.MethodGet, "/test", nil, nil, &result)
	assert.NotNil(t, err)
	assert.Equal(t, string(exp), err.Error())
	teardown()

	// String response
	setup()
	mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		w.Write([]byte(`test string`))
	})

	err = client.makeRequest(context.Background(), http.MethodGet, "/test", nil, nil, &result)
	assert.NotNil(t, err)
	teardown()

	// Err read body
	setup()
	mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		w.Header().Set("Content-Length", "1")
		w.Write([]byte(`test string`))
	})

	err = client.makeRequest(context.Background(), http.MethodGet, "/test", nil, nil, &result)
	assert.NotNil(t, err)
	teardown()

	// Invalid URI
	client.Config.BaseURL = "/this/field/is/illegal/and/should/error/"

	err = client.makeRequest(context.Background(), http.MethodGet, "/test", nil, nil, &result)
	assert.NotNil(t, err)
}

func TestRateLimit(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		w.Header().Set("x-ratelimit-reset", "1458463600")
		w.Header().Set("x-ratelimit-limit", "10")
		w.Header().Set("x-ratelimit-remaining", "9")
		w.WriteHeader(http.StatusTooManyRequests)
		w.Write([]byte(`{
			"meta": {
					"code": 429,
					"type": "TooManyRequests",
					"message": "You have exceeded the API call rate limit. Default limit is 10 requests per second."
			},
			"data": {}
		}`))
	})

	var result mockData
	err := client.makeRequest(context.Background(), http.MethodGet, "/test", nil, nil, &result)

	apiErr := TooManyRequestsError{
		APIError: APIError{
			Code:    429,
			Type:    "TooManyRequests",
			Message: "You have exceeded the API call rate limit. Default limit is 10 requests per second.",
			Path:    "/test",
		},
		RateLimit: client.rateLimit,
	}
	exp, _ := json.Marshal(apiErr)

	assert.NotNil(t, err)
	assert.Equal(t, string(exp), err.Error())
	assert.Equal(t, int64(1458463600), client.rateLimit.Reset)
	assert.Equal(t, 10, client.rateLimit.Limit)
	assert.Equal(t, 9, client.rateLimit.Remaining)
}

func TestBlockRequestWhenReachLimit(t *testing.T) {
	setup()
	defer teardown()

	reset := time.Now().Add(5000).Unix()
	mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		w.Header().Set("x-ratelimit-reset", strconv.FormatInt(reset, 10))
		w.Header().Set("x-ratelimit-limit", "10")
		w.Header().Set("x-ratelimit-remaining", "0")
		w.WriteHeader(http.StatusTooManyRequests)
		w.Write([]byte(`{
			"meta": {
					"code": 429,
					"type": "TooManyRequests",
					"message": "You have exceeded the API call rate limit. Default limit is 10 requests per second."
			},
			"data": {}
		}`))
	})

	var result mockData
	client.makeRequest(context.Background(), http.MethodGet, "/test", nil, nil, &result)
	assert.Equal(t, reset, client.rateLimit.Reset)

	// Another request after exceeded limits
	exp := fmt.Sprintf(errExceedRateLimt, time.Unix(reset, 0))
	err := client.makeRequest(context.Background(), http.MethodGet, "/test", nil, nil, &result)
	assert.NotNil(t, err)
	assert.Equal(t, exp, err.Error())
}

func TestFillingOriginalStatusCode(t *testing.T) {
	setup()
	defer teardown()

	statusTeapot := http.StatusTeapot
	mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPut, r.Method)
		w.WriteHeader(statusTeapot)
		_, _ = w.Write([]byte(`{ "data": "test"}`))
	})

	var result string
	err := client.makeRequest(context.Background(), http.MethodPut, "/test", nil, nil, &result)
	require.Error(t, err)

	var targetErr *APIError
	require.True(t, errors.As(err, &targetErr))

	require.Equal(t, statusTeapot, targetErr.HTTPStatusCode)
}
