package algolia

import (
	"fmt"
	"net/http"
)

type ClientService struct {
	service *Service
	prefix  string
}

func NewClientService(appId, apiKey string) *ClientService {
	service := &Service{
		appId:      appId,
		apiKey:     apiKey,
		httpClient: http.DefaultClient,
	}

	return &ClientService{
		service: service,
	}
}

func (c *ClientService) IsMock() bool {
	return false
}

func (c *ClientService) SetIndexPrefix(p string) {
	c.prefix = p
}

func (c *ClientService) Index(n string) Index {
	return &IndexService{
		name:    fmt.Sprintf("%s%s", c.prefix, n),
		service: c.service,
	}
}
