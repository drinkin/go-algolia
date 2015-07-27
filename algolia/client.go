package algolia

import "net/http"

type ClientService struct {
	service *Service
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

func (c *ClientService) Index(name string) Index {
	return &IndexService{
		name:    name,
		service: c.service,
	}
}
