package algolia_test

import (
	"strconv"

	"github.com/drinkin/di/random"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

const TestIndexName = "go_test"

type Example struct {
	Id       int64  `json:"id"`
	ObjectId string `json:"objectID"`
	Name     string `json:"name"`
}

func (c *Example) AlgoliaBeforeIndex() {
	c.ObjectId = c.AlgoliaId()
}

func (e *Example) AlgoliaId() string {
	return strconv.FormatInt(e.Id, 10)
}

func RandomExample() *Example {
	return &Example{
		Id:   random.Int64(1, 1000000),
		Name: random.Base64(10),
	}
}

func TestAlgolia(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Algolia Suite")
}
