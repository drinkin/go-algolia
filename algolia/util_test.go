package algolia_test

import (
	"strconv"
	"testing"

	"github.com/drinkin/go-algolia/algolia"
	"github.com/stretchr/testify/require"
)

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

func TestHostsForAppId(t *testing.T) {
	assert := require.New(t)
	hosts := algolia.HostsForAppId("a")

	assert.Equal(hosts[0], "a-1.algolianet.com")
}

func TestClient(t *testing.T) {
	assert := require.New(t)
	example := &Example{Id: 1, Name: "george"}
	client := algolia.FromEnv()

	idx := client.Index(TestIndexName)

	tr, err := idx.UpdateObject(example)
	assert.NoError(err)
	assert.Equal(tr.ObjectId, "1")

	savedObj := new(Example)
	err = idx.GetObject("1").Scan(savedObj)
	assert.NoError(err)
	assert.Equal(savedObj, example)

	example.Name = "hi"
	btr, err := idx.BatchUpdate([]algolia.Indexable{example})
	assert.NoError(err)
	assert.Equal(btr.ObjectIds[0], "1")
}
