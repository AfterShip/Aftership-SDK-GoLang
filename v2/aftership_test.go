package aftership

import (
	"testing"

	"github.com/aftership/aftership-sdk-go/v2/common"
	"github.com/stretchr/testify/assert"
)

func TestNilConfig(t *testing.T) {
	// Config is nil
	_, err := NewAfterShip(nil)
	assert.NotNil(t, err)

	expectedMsg := "ConstructorError: config is nil"
	if err.Message != expectedMsg {
		t.Errorf("Expected = %s, but actual = %s", err.Message, expectedMsg)
	}
}

func TestInvalidAPIKey(t *testing.T) {
	// API Key is not specified
	_, err := NewAfterShip(&common.AfterShipConf{})
	assert.NotNil(t, err)

	expectedMsg := "ConstructorError: Invalid API key"
	if err.Message != expectedMsg {
		t.Errorf("Expected = %s, but actual = %s", err.Message, expectedMsg)
	}
}

func TestDefaultConfig(t *testing.T) {
	// API Key is specified
	client, err := NewAfterShip(&common.AfterShipConf{
		APIKey: "YOUR_API_KEY",
	})

	assert.Nil(t, err)
	assert.Equal(t, "YOUR_API_KEY", client.Config.APIKey)
	assert.Equal(t, "https://api.aftership.com/v4", client.Config.Endpoint)
	assert.Equal(t, "aftership-sdk-go", client.Config.UserAagentPrefix)
}

func TestSpecificedConfig(t *testing.T) {
	apiKey := "YOUR_API_KEY"
	endpoint := "YOUR_ENDPOINT"
	agent := "YOUR_AGENT"

	client, err := NewAfterShip(&common.AfterShipConf{
		APIKey:           apiKey,
		Endpoint:         endpoint,
		UserAagentPrefix: agent,
	})

	assert.Nil(t, err)
	assert.Equal(t, apiKey, client.Config.APIKey)
	assert.Equal(t, endpoint, client.Config.Endpoint)
	assert.Equal(t, agent, client.Config.UserAagentPrefix)
}
