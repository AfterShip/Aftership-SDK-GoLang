package aftership

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/google/go-querystring/query"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

// makeRequest makes a AfterShip API calls
func (client *Client) makeRequest(ctx context.Context, method string, path string,
	queryParams interface{}, inputData interface{}, resultData interface{}) error {

	// Check if rate limit is exceeded
	if client.rateLimit != nil && client.rateLimit.isExceeded() {
		return fmt.Errorf(errExceedRateLimt, time.Unix(client.rateLimit.Reset, 0))
	}

	// Read input data
	var body io.Reader
	var bodyStr string
	if inputData != nil {
		jsonData, err := json.Marshal(inputData)
		if err != nil {
			return errors.Wrap(err, "error marshalling params to JSON")
		}

		bodyStr = string(jsonData)
		body = bytes.NewBuffer(jsonData)
	}

	req, err := http.NewRequestWithContext(ctx, method, client.Config.BaseURL+path, body)
	if err != nil {
		return errors.Wrap(err, "HTTP request creation failed")
	}

	// Add headers
	contentType := "application/json"
	req.Header.Add("Content-Type", contentType)
	req.Header.Add("request-id", uuid.New().String())
	req.Header.Add("User-Agent", fmt.Sprintf("%s/%s", client.Config.UserAgentPrefix, VERSION))
	req.Header.Add("aftership-agent", fmt.Sprintf("go-sdk-%s", VERSION))

	if queryParams != nil {
		queryStringObj, err := query.Values(queryParams)
		if err != nil {
			return errors.Wrap(err, "error parsing query params")
		}
		req.URL.RawQuery = queryStringObj.Encode()
	}

	authenticationType := client.Config.AuthenticationType
	apiKey := client.Config.APIKey

	// set signature
	if authenticationType == AES {
		req.Header.Add("as-api-key", apiKey)

		asHeaders := make(map[string]string)
		for key, value := range req.Header {
			asHeaders[key] = value[0]
		}

		date := time.Now().UTC().Format(http.TimeFormat)
		signatureHeader, signature, err := GetSignature(
			authenticationType, []byte(client.Config.APISecret), asHeaders,
			contentType, req.URL.RequestURI(), req.Method, date, bodyStr)
		if err != nil {
			return errors.Wrap(err, "generate signature error")
		}

		req.Header.Add("date", date)
		req.Header.Add(signatureHeader, signature)
	} else {
		req.Header.Add("aftership-api-key", apiKey)
	}

	// Send request
	resp, err := client.httpClient.Do(req)
	if err != nil {
		return errors.Wrap(err, "HTTP request failed")
	}

	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrap(err, "could not read response body")
	}

	// Rate Limit
	setRateLimit(client.rateLimit, resp)

	result := &Response{
		Meta: Meta{},
		Data: resultData,
	}
	// Unmarshal response object
	err = json.Unmarshal(contents, result)
	if err != nil {
		return errors.Wrap(err, "error unmarshalling the JSON response")
	}

	if resp.StatusCode >= http.StatusOK && resp.StatusCode < http.StatusMultipleChoices {
		// The 2xx range indicate success
		return nil
	}

	apiError := APIError{
		Type:    result.Meta.Type,
		Code:    result.Meta.Code,
		Message: result.Meta.Message,
		Path:    path,
	}

	// Too many requests error
	if resp.StatusCode == http.StatusTooManyRequests {
		return &TooManyRequestsError{
			APIError:  apiError,
			RateLimit: client.rateLimit,
		}
	}

	// API error
	return &apiError
}

func setRateLimit(rateLimit *RateLimit, resp *http.Response) {
	if rateLimit != nil && resp != nil && resp.Header != nil {
		// reset timestamp
		if reset := resp.Header.Get("x-ratelimit-reset"); reset != "" {
			if n, err := strconv.ParseInt(reset, 10, 64); err == nil {
				rateLimit.Reset = n
			}
		}

		// limit
		if limit := resp.Header.Get("x-ratelimit-limit"); limit != "" {
			if i, err := strconv.Atoi(limit); err == nil {
				rateLimit.Limit = i
			}
		}

		// remaining
		if remaining := resp.Header.Get("x-ratelimit-remaining"); remaining != "" {
			if i, err := strconv.Atoi(remaining); err == nil {
				rateLimit.Remaining = i
			}
		}
	}
}
