package algolia_test

import (
	"testing"

	"github.com/drinkin/go-algolia/algolia"
	"github.com/k0kubun/pp"
	"github.com/stretchr/testify/require"
)

func TestHostsForAppId(t *testing.T) {
	assert := require.New(t)
	hosts := algolia.HostsForAppId("a")

	assert.Equal(hosts[0], "a-1.algolianet.com")
}

func TestClient(t *testing.T) {
	assert := require.New(t)
	client := algolia.FromEnv()

	var obj map[string]interface{}
	err := client.Index("test_users").GetObject("zxmrjyrp").Scan(&obj)
	assert.NoError(err)
	pp.Print(obj)
}
