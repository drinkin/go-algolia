package algolia

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dghubble/sling"
	"github.com/k0kubun/pp"
)

type Index struct {
	Name    string
	service *Service
}

func (idx *Index) read() *sling.Sling {
	return idx.service.read.Path(fmt.Sprintf("indexes/%s/", idx.Name))
}

func (idx *Index) GetObject(id string, attrs ...string) *Value {
	req, err := idx.read().Get(id).Request()
	pp.Print(req.URL)
	if err != nil {
		return ErrValue(err)
	}

	if len(attrs) > 0 {
		req.Form.Set("attributes", strings.Join(attrs, ","))
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return ErrValue(err)
	}

	return &Value{
		Response: resp,
	}
}
