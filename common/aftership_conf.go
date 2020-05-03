package common

// AfterShipConf is the config of AfterShip SDK client
type AfterShipConf struct {
	APIKey string
	// Endpoint is the URL of AfterShip API, default 'https://api.aftership.com/v4'
	Endpoint string
	// UserAagentPrefix is the prefix of User-Agent in headers, default 'aftership-sdk-go'
	UserAagentPrefix string
}
