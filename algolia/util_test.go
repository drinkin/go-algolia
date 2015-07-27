package algolia_test

import (
	"testing"

	"github.com/drinkin/go-algolia/algolia"
	"github.com/stretchr/testify/require"
)

func TestHostsForAppId(t *testing.T) {
	assert := require.New(t)
	hosts := algolia.HostsForAppId("a")

	assert.Equal(hosts[0], "a-1.algolianet.com")
}
