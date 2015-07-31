package algolia

import (
	"fmt"
	"strconv"
)

// Int64ToString is a utility to convert int64 to objectID string
func Int64ToString(v int64) string {
	return strconv.FormatInt(v, 10)
}

func HostsForAppId(appId string) []string {
	hosts := make([]string, 3)

	for i := 0; i < 3; i++ {
		hosts[i] = fmt.Sprintf("%s-%v.algolianet.com", appId, i+1)
	}
	return hosts
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
