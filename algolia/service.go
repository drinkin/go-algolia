package algolia

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
)

type Service struct {
	appId      string
	apiKey     string
	httpClient *http.Client
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

func (s *Service) makeRequest(m string, u *url.URL, obj interface{}) *httpValue {
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

func (s *Service) Get(pth string) *httpValue {
	u := &url.URL{
		Scheme: "https",
		Host:   s.appId + "-dsn.algolia.net",
		Path:   pth,
	}
	return s.makeRequest("GET", u, nil)

}
func (s *Service) Post(pth string, obj interface{}) *httpValue {
	return s.writeRequest("POST", pth, obj)
}
func (s *Service) Put(pth string, obj interface{}) *httpValue {
	return s.writeRequest("PUT", pth, obj)
}
func (s *Service) Delete(pth string) *httpValue {
	return s.writeRequest("DELETE", pth, nil)
}

func (s *Service) writeRequest(m, pth string, obj interface{}) *httpValue {
	u := &url.URL{
		Scheme: "https",
		Host:   s.appId + ".algolia.net",
		Path:   pth,
	}

	return s.makeRequest(m, u, obj)
}
