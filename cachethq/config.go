package cachethq

import (
	"github.com/andygrunwald/cachet"
)

// Config is per-provider, specifies where to connect to cachethq
type Config struct {
	Token   string
	BaseURL string
}

// Client returns a *cachet.Client to interact with the configured cachethq instance
func (c *Config) Client() (interface{}, error) {
	// Configure TLS/SSL

	client, err := cachet.NewClient(c.BaseURL, nil)
	if err != nil {
		return nil, err
	}

	client.Authentication.SetTokenAuth(c.Token)
	_, _, err = client.General.Ping()
	if err != nil {
		return nil, err
	}

	return client, err
}
