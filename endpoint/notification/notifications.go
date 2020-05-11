package notification

import "github.com/aftership/aftership-sdk-go/v2/response"

// Notification is the model describing an AfterShip notification
type Notification struct {
	Emails []string `json:"emails"`
	Smses  []string `json:"smses"`
}

// Data is the hash describes the notification information.
type Data struct {
	Notification Notification `json:"notification"`
}

// Envelope is the message envelope for the notification API responses
type Envelope struct {
	Meta response.Meta `json:"meta"`
	Data Data          `json:"data"`
}

// GetMeta returns the response meta
func (e *Envelope) GetMeta() response.Meta {
	return e.Meta
}
