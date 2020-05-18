package tracking

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildTrackingUrl(t *testing.T) {
	// slug and tracking number
	p := SingleTrackingParam{
		"",
		"xq-express",
		"LS404494276CN",
		nil,
	}

	url, err := p.BuildTrackingURL("", "")
	assert.Nil(t, err)
	assert.Equal(t, "/trackings/xq-express/LS404494276CN", url)

	// slug and tracking number, with subpath
	url, err = p.BuildTrackingURL("trackings", "retrack")
	assert.Nil(t, err)
	assert.Equal(t, "/trackings/xq-express/LS404494276CN/retrack", url)

	// slug and tracking number, with optional parameters
	p = SingleTrackingParam{
		"",
		"xq-express",
		"LS404494276CN",
		&SingleTrackingOptionalParams{
			TrackingPostalCode: "1234",
		},
	}
	url, err = p.BuildTrackingURL("", "")
	assert.Nil(t, err)
	assert.Equal(t, "/trackings/xq-express/LS404494276CN?tracking_postal_code=1234", url)

	// slug and tracking number, with optional parameters and subpath
	url, err = p.BuildTrackingURL("trackings", "retrack")
	assert.Nil(t, err)
	assert.Equal(t, "/trackings/xq-express/LS404494276CN/retrack?tracking_postal_code=1234", url)

	// id
	p = SingleTrackingParam{
		ID: "1234567890",
	}

	url, err = p.BuildTrackingURL("", "")
	assert.Nil(t, err)
	assert.Equal(t, "/trackings/1234567890", url)

	// id, with subpath
	url, err = p.BuildTrackingURL("trackings", "retrack")
	assert.Nil(t, err)
	assert.Equal(t, "/trackings/1234567890/retrack", url)

	// id, with optional parameters
	p = SingleTrackingParam{
		ID: "1234567890",
		OptionalParams: &SingleTrackingOptionalParams{
			TrackingPostalCode: "1234",
		},
	}
	url, err = p.BuildTrackingURL("", "")
	assert.Nil(t, err)
	assert.Equal(t, "/trackings/1234567890?tracking_postal_code=1234", url)

	// id, with optional parameters and subpath
	url, err = p.BuildTrackingURL("trackings", "retrack")
	assert.Nil(t, err)
	assert.Equal(t, "/trackings/1234567890/retrack?tracking_postal_code=1234", url)

	// Encode slug and tracking number
	p = SingleTrackingParam{
		"",
		"usps",
		"ABCD/1234",
		nil,
	}

	url, err = p.BuildTrackingURL("", "")
	assert.Nil(t, err)
	assert.Equal(t, "/trackings/usps/ABCD%2F1234", url)
}

func TestBuildTrackingUrlError(t *testing.T) {
	// should get error when no id, slug and tracking number
	p := SingleTrackingParam{
		"",
		"",
		"",
		nil,
	}
	_, err := p.BuildTrackingURL("", "")
	assert.NotNil(t, err)

	// should get error when only slug
	p = SingleTrackingParam{
		"",
		"xq-express",
		"",
		nil,
	}
	_, err = p.BuildTrackingURL("", "")
	assert.NotNil(t, err)

	// should get error when only tracking_number
	p = SingleTrackingParam{
		"",
		"",
		"LS404494276CN",
		nil,
	}
	_, err = p.BuildTrackingURL("", "")
	assert.NotNil(t, err)
}

type mockQueryString struct {
	A string `url:"a"`
	B string `url:"b"`
}

func TestBuildURLWithQueryString(t *testing.T) {
	uri := "/trackings"

	// Nil querystring
	actual, err := BuildURLWithQueryString(uri, nil)
	assert.Equal(t, uri, actual)
	assert.Nil(t, err)

	// One querystring
	params := mockQueryString{
		A: "1",
	}
	actual, err = BuildURLWithQueryString(uri, params)
	exp := fmt.Sprintf("%s?a=1&b=", uri)
	assert.Equal(t, exp, actual)
	assert.Nil(t, err)

	// two querystrings
	params = mockQueryString{
		A: "1",
		B: "2",
	}
	actual, err = BuildURLWithQueryString(uri, params)
	exp = fmt.Sprintf("%s?a=1&b=2", uri)
	assert.Equal(t, exp, actual)
	assert.Nil(t, err)

	// uri has querystring, add other querysrings
	uri = "/trackings?test=123"
	params = mockQueryString{
		A: "1",
		B: "2",
	}
	actual, err = BuildURLWithQueryString(uri, params)
	exp = fmt.Sprintf("%s&a=1&b=2", uri)
	assert.Equal(t, exp, actual)
	assert.Nil(t, err)
}

func TestBuildURLWithQueryStringError(t *testing.T) {
	uri := "/trackings"

	// Incorrect querystring obj
	_, err := BuildURLWithQueryString(uri, "test")
	assert.NotNil(t, err)
	assert.Equal(t, "HandlerError", err.Type)
}
