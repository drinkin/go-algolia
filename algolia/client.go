package algolia

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"time"

	"github.com/drinkin/di/env"
)

const (
	defaultConnectTimeout   = 1 * time.Second
	defaultBuildReadTimeout = 30 * time.Second
)

type Service struct {
	appId      string
	apiKey     string
	httpClient *http.Client
}

func (s *Service) writeURL() *url.URL {
	return &url.URL{Scheme: "https", Host: s.appId + ".algolia.net"}
}

func (s *Service) newRequest(m, u string, obj interface{}) (*http.Request, error) {

	b := new(bytes.Buffer)
	if obj != nil {
		if err := json.NewEncoder(b).Encode(obj); err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(m, u, b)
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-Algolia-Application-Id", s.appId)
	req.Header.Set("X-Algolia-API-Key", s.apiKey)
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

func (s *Service) makeRequest(m string, u *url.URL, obj interface{}) *Value {
	req, err := s.newRequest(m, u.String(), obj)
	if err != nil {
		return NewErrValue(err)
	}

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return NewErrValue(err)
	}

	return NewValue(resp)
}

func (s *Service) Get(pth string) *Value {
	u := &url.URL{
		Scheme: "https",
		Host:   s.appId + "-dsn.algolia.net",
		Path:   pth,
	}
	return s.makeRequest("GET", u, nil)

}
func (s *Service) Post(pth string, obj interface{}) *Value {
	return s.writeRequest("POST", pth, obj)
}
func (s *Service) Put(pth string, obj interface{}) *Value {
	return s.writeRequest("PUT", pth, obj)
}

func (s *Service) writeRequest(m, pth string, obj interface{}) *Value {
	u := &url.URL{
		Scheme: "https",
		Host:   s.appId + ".algolia.net",
		Path:   pth,
	}

	return s.makeRequest(m, u, obj)
}

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

func (c *Client) Index(name string) *Index {
	return &Index{
		Name:    name,
		service: c.service,
	}
}
