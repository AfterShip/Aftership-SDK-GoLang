package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/aftership/aftership-sdk-go/v2/common"
	"github.com/aftership/aftership-sdk-go/v2/error"
	"github.com/aftership/aftership-sdk-go/v2/response"
	"github.com/google/uuid"
)

// APIRequest API request interface
type APIRequest interface {
	// MakeRequest Make a AfterShip API calls
	MakeRequest(method string, uri string, data interface{}, result response.AftershipResponse) *error.AfterShipError
}

// APIRequestImpl is the implementation of
type APIRequestImpl struct {
	Client           *http.Client
	RateLimit        *response.RateLimit
	APIKey           string
	Endpoint         string
	UserAagentPrefix string
}

// NewRequest returns the instance of API Request
func NewRequest(cfg *common.AfterShipConf, limit *response.RateLimit) APIRequest {
	return &APIRequestImpl{
		APIKey:           cfg.APIKey,
		Endpoint:         cfg.Endpoint,
		UserAagentPrefix: cfg.UserAagentPrefix,
		RateLimit:        limit,
	}
}

// MakeRequest Make a AfterShip API calls
func (impl *APIRequestImpl) MakeRequest(method string, uri string, data interface{}, result response.AftershipResponse) *error.AfterShipError {
	if impl.Client == nil {
		impl.Client = &http.Client{}
	}

	var body io.Reader
	if data != nil {
		jsonData, err := json.Marshal(data)
		if err != nil {
			return error.MakeRequestError("JsonError", err, data)
		}

		body = bytes.NewBuffer(jsonData)
	}

	req, err := http.NewRequest(method, impl.Endpoint+uri, body)
	if err != nil {
		return error.MakeRequestError("RequestError", err, data)
	}

	// Add headers
	req.Header.Add("aftership-api-key", impl.APIKey)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("request-id", uuid.New().String())
	req.Header.Add("User-Agent", fmt.Sprintf("%s/%s", impl.UserAagentPrefix, common.VERSION))
	req.Header.Add("aftership-agent", fmt.Sprintf("go-sdk-%s", common.VERSION))

	// Send request
	resp, err := impl.Client.Do(req)
	if err != nil {
		return error.MakeRequestError("RequestError", err, data)
	}

	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)

	// Rate Limit
	setRateLimit(impl.RateLimit, resp)

	// Unmarshal response object
	err = json.Unmarshal(contents, &result)
	if err != nil {
		return error.MakeRequestError("RequestError", err, string(contents))
	}

	if resp.StatusCode >= http.StatusOK && resp.StatusCode < http.StatusMultipleChoices {
		// The 2xx range indicate success
		return nil
	}

	return error.MakeAPIError(result)
}

func setRateLimit(rateLimit *response.RateLimit, resp *http.Response) {
	if rateLimit != nil && resp != nil && resp.Header != nil {
		reset := resp.Header.Get("x-ratelimit-reset")
		n, err := strconv.ParseInt(reset, 10, 64)
		if err == nil {
			rateLimit.Reset = n
		}

		// limit
		limit := resp.Header.Get("x-ratelimit-limit")
		i, err := strconv.Atoi(limit)
		if err == nil {
			rateLimit.Limit = i
		}

		// remaining
		remaining := resp.Header.Get("x-ratelimit-remaining")
		i, err = strconv.Atoi(remaining)
		if err == nil {
			rateLimit.Remaining = i
		}
	}
}
