package algolia

import "github.com/drinkin/di/env"

// Indexable represents objects that can be saved to the search index.
type Indexable interface {
	// AlgoliaId returns a value that should be used for the `objectID`
	AlgoliaId() string

	// AlgoliaBeforeIndex is called for each item before indexing.
	// You should set the model's objectID here if you want to batchUpdate.
	AlgoliaBeforeIndex()
}

// A Client connects to the algolia service.
type Client interface {
	Index(string) Index

	// SetIndexPrefix allows you to set a prefix for all following Indexes.
	// This is useful for
	SetIndexPrefix(string)
}

// Index represents a backend.
type Index interface {
	// Name returns the index name.
	// If the client had `SetIndexPrefix`, it will be included.
	Name() string

	// GetTaskStatus checks on the status of a task.
	GetTaskStatus(id int64) (*TaskStatus, error)
	UpdateObject(Indexable) (*Task, error)
	BatchUpdate([]Indexable) (*Task, error)
	GetObject(id string, attrs ...string) Value
	Settings() *SettingsBuilder
	SetSettings(*Settings) (*Task, error)
	Clear() (*Task, error)
	DeleteObject(id string) Value
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

// FromEnv creates a new Client
// The environment variables `ALGOLIA_APP_ID` and `ALGOLIA_API_KEY` are used.
// If useMock is true the client is a fake algolia implementation.
func FromEnv(useMock ...bool) Client {
	return New(env.MustGet("ALGOLIA_APP_ID"), env.MustGet("ALGOLIA_API_KEY"), useMock...)
}
