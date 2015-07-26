package algolia

import (
	"encoding/json"
	"net/http"
)

func ErrValue(err error) *Value {
	return &Value{
		err: err,
	}
}

type Value struct {
	err      error
	Response *http.Response
}

func (v *Value) Scan(obj interface{}) error {
	if v.err != nil {
		return v.err
	}

	defer v.Response.Body.Close()

	return json.NewDecoder(v.Response.Body).Decode(obj)

}
