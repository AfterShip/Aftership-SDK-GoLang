package error

import (
	"github.com/aftership/aftership-sdk-go/v2/response"
)

// AfterShipError is the error in AfterShip API calls
type AfterShipError struct {
	Type    string      `json:"type"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// NewSdkError Make SDK error
func NewSdkError(errType string, msg string, data interface{}) *AfterShipError {
	return &AfterShipError{
		Type:    errType,
		Message: msg,
		Data:    data,
	}
}

// NewRequestError Make request error
func NewRequestError(errType string, reqError error, data interface{}) *AfterShipError {
	return &AfterShipError{
		Type:    errType,
		Message: reqError.Error(),
		Data:    data,
	}
}

// NewAPIError Make API error
func NewAPIError(resp response.AftershipResponse) *AfterShipError {
	if resp == nil {
		return &AfterShipError{
			Type: "InternalError",
			Code: 500,
		}
	}

	meta := resp.GetMeta()
	return &AfterShipError{
		Type:    meta.Type,
		Code:    meta.Code,
		Message: meta.Message,
	}
}
