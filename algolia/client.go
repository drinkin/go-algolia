package algolia

import (
	"fmt"
	"time"

	"github.com/dghubble/sling"
	"github.com/drinkin/di/env"
)

const (
	defaultConnectTimeout   = 1 * time.Second
	defaultBuildReadTimeout = 30 * time.Second
)

type Service struct {
	base  *sling.Sling
	read  *sling.Sling
	write *sling.Sling
}

type Client struct {
	appId  string
	apiKey string

	service *Service
}

func FromEnv() *Client {
	return NewClient(env.MustGet("ALGOLIA_APP_ID"), env.MustGet("ALGOLIA_API_KEY"))
}

func NewClient(appId, apiKey string) *Client {

	base := sling.New().
		Add("X-Algolia-Application-Id", appId).
		Add("X-Algolia-API-Key", apiKey)

	service := &Service{
		base:  base,
		read:  base.Base(fmt.Sprintf("https://%s-dsn.algolia.net/1/", appId)),
		write: base.Base(fmt.Sprintf("https://%s.algolia.net/1/", appId)),
	}

	return &Client{
		appId:   appId,
		apiKey:  apiKey,
		service: service,
	}
}

func (c *Client) Index(name string) *Index {
	return &Index{
		Name:    name,
		service: c.service,
	}
}
