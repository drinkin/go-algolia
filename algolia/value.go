package algolia

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type httpValue struct {
	r   *http.Response
	err error
}

func NewErrValue(err error) *httpValue {
	return &httpValue{err: err}
}

func NewValue(r *http.Response) *httpValue {
	return &httpValue{r: r}
}

func (v *httpValue) getBody() (io.ReadCloser, error) {
	if v.err != nil {
		return nil, v.err
	}
	if v.r == nil {
		return nil, fmt.Errorf("No Response")
	}

	return v.r.Body, nil
}

func (v *httpValue) Scan(obj interface{}) error {
	b, err := v.getBody()
	if err != nil {
		return err
	}
	defer b.Close()

	decoder := json.NewDecoder(b)

	if v.r.StatusCode >= 300 {
		var apiErr Err
		if err := decoder.Decode(&apiErr); err != nil {
			return err
		}
		return &apiErr
	}

	return decoder.Decode(obj)

}
