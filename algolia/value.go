package algolia

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Value struct {
	*http.Response
	err error
}

func NewErrValue(err error) *Value {
	return &Value{err: err}
}

func NewValue(r *http.Response) *Value {
	return &Value{Response: r}
}

func (v *Value) asTaskResp() (*TaskResp, error) {
	tr := new(TaskResp)
	return tr, v.Scan(tr)
}

func (v *Value) getBody() (io.ReadCloser, error) {
	if v.err != nil {
		return nil, v.err
	}
	if v.Response == nil {
		return nil, fmt.Errorf("No Response")
	}

	return v.Body, nil
}

func (v *Value) String() (string, error) {
	b, err := v.getBody()
	if err != nil {
		return "", err
	}
	defer b.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(b)

	return buf.String(), nil
}

func (v *Value) Scan(obj interface{}) error {
	b, err := v.getBody()
	if err != nil {
		return err
	}
	defer b.Close()

	decoder := json.NewDecoder(b)

	if v.StatusCode >= 300 {
		var apiErr Err
		if err := decoder.Decode(&apiErr); err != nil {
			return err
		}
		return &apiErr
	}

	return decoder.Decode(obj)

}
