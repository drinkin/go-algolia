package algolia

type BatchItem struct {
	Action string    `json:"action"`
	Body   Indexable `json:"body"`
}
