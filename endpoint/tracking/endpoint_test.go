package tracking

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/aftership/aftership-sdk-go/v2/common"
	"github.com/aftership/aftership-sdk-go/v2/request"
	"github.com/aftership/aftership-sdk-go/v2/response"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestCreateTracking(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	data := NewTrackingRequest{
		Tracking: NewTracking{
			Slug:           []string{"dhl"},
			TrackingNumber: "tracking_number",
			Title:          "Title Name",
			Smses: []string{
				"+18555072509",
				"+18555072501",
			},
		},
	}

	exp := SingleTrackingData{
		Tracking: Tracking{
			Slug:  "dhl",
			Title: "Title Name",
			Smses: []string{
				"+18555072509",
				"+18555072501",
			},
		},
	}

	mockhttp("POST", "/trackings", 200, SingleTrackingEnvelope{
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
	res, _ := endpoint.CreateTracking(context.Background(), data)
	assert.Equal(t, exp, res)
}

func TestCreateTrackingError(t *testing.T) {
	req := request.NewRequest(&common.AfterShipConf{
		APIKey: "YOUR_API_KEY",
	}, nil)
	endpoint := NewEndpoint(req)

	// Response with error
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	data := NewTrackingRequest{
		Tracking: NewTracking{
			TrackingNumber: "tracking_number",
		},
	}

	mockhttp("POST", "/trackings", 401, SingleTrackingEnvelope{
		response.Meta{
			Code:    401,
			Message: "Invalid API key.",
			Type:    "Unauthorized",
		},
		SingleTrackingData{},
	}, nil)

	_, err := endpoint.CreateTracking(context.Background(), data)
	assert.NotNil(t, err)
	assert.Equal(t, "Unauthorized", err.Type)
}

func TestDeleteTracking(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	p := SingleTrackingParam{
		Slug:           "ups",
		TrackingNumber: "1Z9999999999999998",
	}

	exp := SingleTrackingData{
		Tracking: Tracking{
			Slug:  "ups",
			Title: "Title Name",
		},
	}

	mockhttp("DELETE", fmt.Sprintf("/trackings/%s/%s", p.Slug, p.TrackingNumber), 200, SingleTrackingEnvelope{
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
	res, _ := endpoint.DeleteTracking(context.Background(), p)
	assert.Equal(t, exp, res)
}

func TestDeleteTrackingError(t *testing.T) {
	req := request.NewRequest(&common.AfterShipConf{
		APIKey: "YOUR_API_KEY",
	}, nil)
	endpoint := NewEndpoint(req)

	// empty id, slug and tracking_number
	p := SingleTrackingParam{
		ID:             "",
		Slug:           "",
		TrackingNumber: "",
		OptionalParams: nil,
	}

	_, err := endpoint.DeleteTracking(context.Background(), p)
	assert.NotNil(t, err)
	assert.Equal(t, "HandlerError", err.Type)

	// Response with error
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	p = SingleTrackingParam{
		ID:             "",
		Slug:           "xq-express",
		TrackingNumber: "LS404494276CN",
		OptionalParams: nil,
	}

	mockhttp("DELETE", fmt.Sprintf("/trackings/%s/%s", p.Slug, p.TrackingNumber), 401, SingleTrackingEnvelope{
		response.Meta{
			Code:    401,
			Message: "Invalid API key.",
			Type:    "Unauthorized",
		},
		SingleTrackingData{},
	}, nil)

	_, err = endpoint.DeleteTracking(context.Background(), p)
	assert.NotNil(t, err)
	assert.Equal(t, "Unauthorized", err.Type)
}

func TestGetTrackings(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	p := MultiTrackingsParams{
		Page:  1,
		Limit: 10,
	}

	exp := MultiTrackingsData{
		Page:  1,
		Count: 2,
		Limit: 10,
		Trackings: []Tracking{
			{
				Slug:  "dhl",
				Title: "Title Name",
			},
			{
				Slug:  "ups",
				Title: "Title Name",
			},
		},
	}

	mockhttp("GET", fmt.Sprintf("/trackings?limit=%d&page=%d", p.Limit, p.Page), 200, MultiTrackingsEnvelope{
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
	res, _ := endpoint.GetTrackings(context.Background(), p)
	assert.Equal(t, exp, res)
}

func TestGetTrackingsError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	p := MultiTrackingsParams{
		Page:  1,
		Limit: 10,
	}

	mockhttp("GET", fmt.Sprintf("/trackings?limit=%d&page=%d", p.Limit, p.Page), 401, MultiTrackingsEnvelope{
		response.Meta{
			Code:    401,
			Message: "Invalid API key.",
			Type:    "Unauthorized",
		},
		MultiTrackingsData{},
	}, nil)

	req := request.NewRequest(&common.AfterShipConf{
		APIKey: "YOUR_API_KEY",
	}, nil)
	endpoint := NewEndpoint(req)
	_, err := endpoint.GetTrackings(context.Background(), p)
	assert.NotNil(t, err)
	assert.Equal(t, "Unauthorized", err.Type)
}

func TestGetTracking(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	p := SingleTrackingParam{
		Slug:           "ups",
		TrackingNumber: "1Z9999999999999998",
	}

	exp := SingleTrackingData{
		Tracking: Tracking{
			Slug:  "ups",
			Title: "Title Name",
		},
	}

	mockhttp("GET", fmt.Sprintf("/trackings/%s/%s", p.Slug, p.TrackingNumber), 200, SingleTrackingEnvelope{
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
	res, _ := endpoint.GetTracking(context.Background(), p, nil)
	assert.Equal(t, exp, res)
}

func TestGetTrackingError(t *testing.T) {
	req := request.NewRequest(&common.AfterShipConf{
		APIKey: "YOUR_API_KEY",
	}, nil)
	endpoint := NewEndpoint(req)

	// empty id, slug and tracking_number
	p := SingleTrackingParam{
		ID:             "",
		Slug:           "",
		TrackingNumber: "",
		OptionalParams: nil,
	}

	_, err := endpoint.GetTracking(context.Background(), p, nil)
	assert.NotNil(t, err)
	assert.Equal(t, "HandlerError", err.Type)

	// Response with error
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	p = SingleTrackingParam{
		ID:             "",
		Slug:           "xq-express",
		TrackingNumber: "LS404494276CN",
		OptionalParams: nil,
	}

	mockhttp("GET", fmt.Sprintf("/trackings/%s/%s", p.Slug, p.TrackingNumber), 401, SingleTrackingEnvelope{
		response.Meta{
			Code:    401,
			Message: "Invalid API key.",
			Type:    "Unauthorized",
		},
		SingleTrackingData{},
	}, nil)

	_, err = endpoint.GetTracking(context.Background(), p, nil)
	assert.NotNil(t, err)
	assert.Equal(t, "Unauthorized", err.Type)
}

func TestUpdateTracking(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	p := SingleTrackingParam{
		Slug:           "ups",
		TrackingNumber: "1Z9999999999999998",
	}

	data := UpdateTrackingRequest{
		UpdateTracking{
			Title: "New Title",
		},
	}

	exp := SingleTrackingData{
		Tracking: Tracking{
			Slug:  "ups",
			Title: "New Title",
		},
	}

	mockhttp("PUT", fmt.Sprintf("/trackings/%s/%s", p.Slug, p.TrackingNumber), 200, SingleTrackingEnvelope{
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
	res, _ := endpoint.UpdateTracking(context.Background(), p, data)
	assert.Equal(t, exp, res)
}

func TestUpdateTrackingError(t *testing.T) {
	req := request.NewRequest(&common.AfterShipConf{
		APIKey: "YOUR_API_KEY",
	}, nil)
	endpoint := NewEndpoint(req)

	// empty id, slug and tracking_number
	p := SingleTrackingParam{
		ID:             "",
		Slug:           "",
		TrackingNumber: "",
		OptionalParams: nil,
	}

	data := UpdateTrackingRequest{
		UpdateTracking{
			Title: "New Title",
		},
	}

	_, err := endpoint.UpdateTracking(context.Background(), p, data)
	assert.NotNil(t, err)
	assert.Equal(t, "HandlerError", err.Type)

	// Response with error
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	p = SingleTrackingParam{
		ID:             "",
		Slug:           "xq-express",
		TrackingNumber: "LS404494276CN",
		OptionalParams: nil,
	}

	mockhttp("PUT", fmt.Sprintf("/trackings/%s/%s", p.Slug, p.TrackingNumber), 401, SingleTrackingEnvelope{
		response.Meta{
			Code:    401,
			Message: "Invalid API key.",
			Type:    "Unauthorized",
		},
		SingleTrackingData{},
	}, nil)

	_, err = endpoint.UpdateTracking(context.Background(), p, data)
	assert.NotNil(t, err)
	assert.Equal(t, "Unauthorized", err.Type)
}

func TestReTrack(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	p := SingleTrackingParam{
		Slug:           "ups",
		TrackingNumber: "1Z9999999999999998",
	}

	exp := SingleTrackingData{
		Tracking: Tracking{
			Slug:  "ups",
			Title: "Title Name",
		},
	}

	mockhttp("POST", fmt.Sprintf("/trackings/%s/%s/retrack", p.Slug, p.TrackingNumber), 200, SingleTrackingEnvelope{
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
	res, _ := endpoint.ReTrack(context.Background(), p)
	assert.Equal(t, exp, res)
}

func TestReTrackError(t *testing.T) {
	req := request.NewRequest(&common.AfterShipConf{
		APIKey: "YOUR_API_KEY",
	}, nil)
	endpoint := NewEndpoint(req)

	// empty id, slug and tracking_number
	p := SingleTrackingParam{
		ID:             "",
		Slug:           "",
		TrackingNumber: "",
		OptionalParams: nil,
	}

	_, err := endpoint.ReTrack(context.Background(), p)
	assert.NotNil(t, err)
	assert.Equal(t, "HandlerError", err.Type)

	// Response with error
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	p = SingleTrackingParam{
		ID:             "",
		Slug:           "xq-express",
		TrackingNumber: "LS404494276CN",
		OptionalParams: nil,
	}

	mockhttp("POST", fmt.Sprintf("/trackings/%s/%s/retrack", p.Slug, p.TrackingNumber), 401, SingleTrackingEnvelope{
		response.Meta{
			Code:    401,
			Message: "Invalid API key.",
			Type:    "Unauthorized",
		},
		SingleTrackingData{},
	}, nil)

	_, err = endpoint.ReTrack(context.Background(), p)
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
