package algolia

// Indexable represents objects that can be saved to the search index.
type Indexable interface {
	AlgoliaId() string

	AlgoliaBeforeIndex()
}
