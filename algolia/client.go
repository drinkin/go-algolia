package algolia

import (
	"net/http"

	"github.com/drinkin/di/env"
)

type Client struct {
	service *Service
}

func FromEnv() *Client {
	return NewClient(env.MustGet("ALGOLIA_APP_ID"), env.MustGet("ALGOLIA_API_KEY"))
}

func NewClient(appId, apiKey string) *Client {
	service := &Service{
		appId:      appId,
		apiKey:     apiKey,
		httpClient: http.DefaultClient,
	}

	return &Client{
		service: service,
	}
}

func (c *Client) MockIndex(name string) Index {
	return NewIndexMock(name)
}

func (c *Client) Index(name string) Index {
	return &IndexService{
		name:    name,
		service: c.service,
	}
}
