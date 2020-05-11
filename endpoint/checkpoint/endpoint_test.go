package checkpoint

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/aftership/aftership-sdk-go/v2/common"
	"github.com/aftership/aftership-sdk-go/v2/request"
	"github.com/aftership/aftership-sdk-go/v2/response"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestGetLastCheckpoint(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	p := common.SingleTrackingParam{
		Slug:           "xq-express",
		TrackingNumber: "LS404494276CN",
	}
	exp := LastCheckpoint{
		ID:             "5b74f4958776db0e00b6f5ed",
		TrackingNumber: "111111111111",
		Checkpoint: Checkpoint{
			Slug: "slug",
		},
	}
	mockhttp("GET", fmt.Sprintf("/last_checkpoint/%s/%s", p.Slug, p.TrackingNumber), 200, LastCheckpointEnvelope{
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
	res, err := endpoint.GetLastCheckpoint(p, nil)
	assert.Equal(t, exp, res)
	assert.Nil(t, err)
}

func TestGetLastCheckpointWithOptionalParams(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	p := common.SingleTrackingParam{
		Slug:           "xq-express",
		TrackingNumber: "LS404494276CN",
	}

	op := &GetCheckpointParams{
		Fields: "slug",
		Lang:   "en",
	}

	exp := LastCheckpoint{
		ID:             "5b74f4958776db0e00b6f5ed",
		TrackingNumber: "111111111111",
		Checkpoint: Checkpoint{
			Slug: "slug",
		},
	}
	mockhttp("GET", fmt.Sprintf("/last_checkpoint/%s/%s?fields=%s&lang=%s", p.Slug, p.TrackingNumber, op.Fields, op.Lang), 200, LastCheckpointEnvelope{
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
	res, err := endpoint.GetLastCheckpoint(p, op)
	assert.Equal(t, exp, res)
	assert.Nil(t, err)
}

func TestError(t *testing.T) {
	req := request.NewRequest(&common.AfterShipConf{}, nil)
	endpoint := NewEndpoint(req)

	// empty id, slug and tracking_number
	p := common.SingleTrackingParam{
		ID:             "",
		Slug:           "",
		TrackingNumber: "",
		OptionalParams: nil,
	}

	_, err := endpoint.GetLastCheckpoint(p, nil)
	assert.NotNil(t, err)
	assert.Equal(t, "HandlerError", err.Type)

	// Response with error
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	p = common.SingleTrackingParam{
		Slug:           "xq-express",
		TrackingNumber: "LS404494276CN",
	}

	mockhttp("GET", fmt.Sprintf("/last_checkpoint/%s/%s", p.Slug, p.TrackingNumber), 401, LastCheckpointEnvelope{
		response.Meta{
			Code:    401,
			Message: "Invalid API key.",
			Type:    "Unauthorized",
		},
		LastCheckpoint{},
	}, nil)

	_, err = endpoint.GetLastCheckpoint(p, nil)
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
