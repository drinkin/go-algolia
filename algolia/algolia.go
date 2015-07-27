package algolia

import "github.com/drinkin/di/env"

// Indexable represents objects that can be saved to the search index.
type Indexable interface {
	AlgoliaId() string

	AlgoliaBeforeIndex()
}

type Client interface {
	Index(string) Index
}

// Index represents a backend.
type Index interface {
	Name() string
	Must() *MustIndex
	GetTaskStatus(id int64) (*TaskStatus, error)
	UpdateObject(Indexable) (*Task, error)
	BatchUpdate([]Indexable) (*BatchTask, error)
	GetObject(id string, attrs ...string) Value
}

type Value interface {
	Scan(obj interface{}) error
}

func New(appId, apiKey string, useMock ...bool) Client {
	if len(useMock) > 0 && useMock[0] {
		return NewClientMock()
	}

	return NewClientService(appId, apiKey)
}

func FromEnv(useMock ...bool) Client {
	return New(env.MustGet("ALGOLIA_APP_ID"), env.MustGet("ALGOLIA_API_KEY"), useMock...)
}
