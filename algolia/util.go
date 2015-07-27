package algolia

import "fmt"

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
